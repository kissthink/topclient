package topclient

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
)

type IRequest interface {
	check() (bool, error)
	getApiMethodName() string
	getApiParas() (map[string]string, error)
}

type TopClient struct {
	AppKey       string `json:"app_key"`
	SecretKey    string `json:"secret_key"`
	GateWayUrl   string `json:"gate_way_url"`
	DataFormate  string `json:"data_formate"`
	ConnTimeout  int    `json:"conn_timeout"`
	ReadTimeout  int    `json:"read_timeout"`
	CheckRequest bool   `json:"check_request"`
	SignMethod   string `json:"sign_method"`
	ApiVersion   string `json:"api_version"`
	SdkVersion   string `json:"sdk_version"`
}

// 处理结果结构体
type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	instance *TopClient
	once     sync.Once
)

func NewClient(appKey, secretKey string) *TopClient {
	if instance != nil {
		return instance
	}
	once.Do(func() {
		instance = &TopClient{
			AppKey:       appKey,
			SecretKey:    secretKey,
			GateWayUrl:   `http://gw.api.taobao.com/router/rest`,
			DataFormate:  `json`,
			ConnTimeout:  5,
			ReadTimeout:  3,
			CheckRequest: true,
			SignMethod:   `md5`,
			ApiVersion:   `2.0`,
			SdkVersion:   `top-sdk-php-20151012`,
		}
	})
	return instance
}

func (client *TopClient) Execute(request IRequest, session, baseUrl string) *Result {
	result := &Result{Code: 0, Message: "OK"}
	// 是否开启请求入参检测
	if client.CheckRequest {
		checkStatus, err := request.check()
		if err != nil || checkStatus != true {
			result.Message = `请求检测错误:` + err.Error()
			return result
		}
	}
	//组装系统参数
	sysParams := make(map[string]string)
	sysParams[`app_key`] = client.AppKey
	sysParams[`v`] = client.ApiVersion
	sysParams[`format`] = client.DataFormate
	sysParams[`sign_method`] = client.SignMethod
	sysParams[`method`] = request.getApiMethodName()
	sysParams[`timestamp`] = time.Now().Format(`2006-01-02 15:04:05`)
	// 是否有会话信息
	if len(session) > 0 {
		sysParams[`session`] = session
	}
	var err error
	apiParams := make(map[string]string)
	apiParams, err = request.getApiParas()
	if err != nil {
		result.Message = `API参数错误:` + err.Error()
		return result
	}
	var requestUrl string
	if len(baseUrl) > 0 {
		requestUrl = baseUrl + `?`
		sysParams[`partner_id`] = client.getClusterTag()
	} else {
		requestUrl = client.GateWayUrl + `?`
		sysParams[`partner_id`] = client.SdkVersion
	}
	sysParams[`sign`] = client.createSign(merge(sysParams, apiParams))
	// 检测API参数中是否有上传文件的标识
	fileField := make(map[string]string)
	//for key, val := range apiParams {
	//
	//}
	query := &url.Values{}
	for field, value := range sysParams {
		query.Add(field, value)
	}
	// 拼接请求的URL
	requestUrl = requestUrl + query.Encode()
	var body string
	if len(fileField) > 0 {
		// 走文件上传的POST方法
	} else {
		_, body, err = client.httpPost(requestUrl, apiParams)
	}
	if err != nil {
		result.Message = `POST请求错误:` + err.Error()
		return result
	}
	result.Data = body
	return result
}

// 发起HTTP POST请求
func (client *TopClient) httpPost(gatewayURL string, data map[string]string) (map[string]string, string, error) {
	if !(strings.HasPrefix(gatewayURL, "https") || strings.HasPrefix(gatewayURL, "http")) {
		gatewayURL = "https://" + gatewayURL
	}
	host, err := url.ParseRequestURI(gatewayURL)
	if err != nil {
		return map[string]string{}, "", errors.New(`URL解析错误`)
	}
	ips, err := net.LookupIP(host.Host)
	if err != nil || len(ips) == 0 {
		return map[string]string{}, "", errors.New(`DNS解析错误`)
	}
	query := &url.Values{}
	// 整理参数
	if len(data) > 0 {
		for key, val := range data {
			query.Add(key, val)
		}
	}
	// 向主机请求数据
	httpClient := &http.Client{}
	request, err := http.NewRequest("POST", gatewayURL, strings.NewReader(query.Encode()))
	if err != nil {
		return map[string]string{}, "", errors.New(`发起请求失败`)
	}
	request.Header.Set("User-Agent", "top-sdk-php")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Set("Accept-Encoding", "gzip")
	response, err := httpClient.Do(request)
	if err != nil {
		return map[string]string{}, "", errors.New(`请求发生错误`)
	}
	defer response.Body.Close()
	var body string
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ := gzip.NewReader(response.Body)
		for {
			buffer := make([]byte, 1024)
			cnt, err := reader.Read(buffer)
			if err != nil && err != io.EOF {
				break
			}
			if cnt == 0 {
				break
			}
			body += string(buffer)
		}
		defer reader.Close()
	default:
		buffer, _ := ioutil.ReadAll(response.Body)
		body = string(buffer)
	}
	if err != nil {
		return map[string]string{}, "", errors.New(`数据读取错误`)
	}
	headerBuffer := new(bytes.Buffer)
	response.Header.Write(headerBuffer)
	// 去除最后一个换行符
	headerStr := strings.TrimRight(headerBuffer.String(), "\r\n")
	list := strings.Split(headerStr, "\r\n")
	var headerList = make(map[string]string)
	for _, value := range list {
		headerItem := strings.Split(value, ":")
		headerList[headerItem[0]] = headerItem[1]
	}
	return headerList, body, nil
}

// 创建签名
func (client *TopClient) createSign(params map[string]string) string {
	paramsList := make([]string, len(params))
	for key, _ := range params {
		paramsList = append(paramsList, key)
	}
	sort.Strings(paramsList)
	signStr := client.SecretKey
	for _, k := range paramsList {
		signStr = signStr + k + params[k]
	}
	signStr = signStr + client.SecretKey
	hash := md5.New()
	hash.Write([]byte(signStr))
	cipherStr := hash.Sum(nil)
	signStr = hex.EncodeToString(cipherStr)
	return strings.ToUpper(signStr)
}

func (client *TopClient) getClusterTag() string {
	return subStr(client.SdkVersion, 0, 11) + "-cluster" + subStr(client.SdkVersion, 11, 10)
}

// 合并数组
func merge(arr1, arr2 map[string]string) map[string]string {
	for k, v := range arr1 {
		arr2[k] = v
	}
	return arr2
}

// 截取字符串
func subStr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}
