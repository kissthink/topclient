package topclient

import (
	"encoding/json"
	"errors"
)

// 自定发送消息接口
type OpenimCustmsgPushRequest struct {
	FromUserId string   `json:"from_user"`
	ToAppKey   string   `json:"to_appkey"`
	ToUsers    []string `json:"to_users"`
	Summary    string   `json:"summary"`
	Data       string   `json:"data"`
	Aps        string   `json:"aps"`
	Apns_param string   `json:"apns_param"`
	Invisible  int      `json:"invisible"`
	FromNick   string   `json:"from_nick"`
	FromTaobao int      `json:"from_taobao"`
}

func (msg *OpenimCustmsgPushRequest) getApiMethodName() string {
	return `taobao.openim.custmsg.push`
}

func (msg *OpenimCustmsgPushRequest) check() (bool, error) {
	if len(msg.FromUserId) == 0 {
		return false, errors.New(`发送方用户ID不能为空`)
	}
	if len(msg.ToUsers) == 0 {
		return false, errors.New(`请指定发送的目标用户ID`)
	}
	if len(msg.ToUsers) > 100 {
		return false, errors.New(`一次性发送用户最多不能超过100个`)
	}
	if len(msg.Summary) == 0 {
		return false, errors.New(`消息摘要[summary]不能为空`)
	}
	if len(msg.Data) == 0 {
		return false, errors.New(`发送的消息不能为空`)
	}
	return true, nil
}

func (msg *OpenimCustmsgPushRequest) getApiParas() (map[string]string, error) {
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return map[string]string{}, err
	}
	var result = make(map[string]string)
	result[`custmsg`] = string(jsonData)
	return result, nil
}
