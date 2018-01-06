package topclient

import (
	"encoding/json"
	"errors"
)

// 账号聊天关系接口
type OpenimRelationsGetRequest struct {
	BegDate string      `json:"beg_date"`
	EndDate string      `json:"end_date"`
	User    *OpenImUser `json:"user"`
}

type OpenImUser struct {
	Uid           string `json:"uid"`
	TaobaoAccount bool   `json:"taobao_account"`
	AppKey        string `json:"app_key"`
}

func (user *OpenimRelationsGetRequest) getApiMethodName() string {
	return `taobao.openim.relations.get`
}

func (user *OpenimRelationsGetRequest) check() (bool, error) {
	if len(user.User.Uid) == 0 {
		return false, errors.New(`用户ID不能为空`)
	}
	if len(user.BegDate) == 0 || len(user.EndDate) == 0 {
		return false, errors.New(`查询时间范围不能为空`)
	}
	return true, nil
}

func (user *OpenimRelationsGetRequest) getApiParas() (map[string]string, error) {
	jsonData, err := json.Marshal(user.User)
	if err != nil {
		return map[string]string{}, err
	}
	var result = make(map[string]string)
	result[`beg_date`] = user.BegDate
	result[`end_date`] = user.EndDate
	result[`user`] = string(jsonData)
	return result, nil
}
