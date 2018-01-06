package topclient

import (
	"encoding/json"
	"errors"
	"strings"
)

// 获取用户群列表
type OpenimTribeGetalltribesRequest struct {
	User *OpenImUser `json:"user"`
	Type []string    `json:"tribe_types"`
}

func (tribe *OpenimTribeGetalltribesRequest) getApiMethodName() string {
	return `taobao.openim.tribe.getalltribes`
}

func (tribe *OpenimTribeGetalltribesRequest) check() (bool, error) {
	if len(tribe.User.Uid) == 0 {
		return false, errors.New(`用户ID不能为空`)
	}
	if len(tribe.User.AppKey) == 0 {
		return false, errors.New(`APPKEY不能为空`)
	}
	return true, nil
}

func (tribe *OpenimTribeGetalltribesRequest) getApiParas() (map[string]string, error) {
	userInfo, err := json.Marshal(tribe.User)
	if err != nil {
		return map[string]string{}, err
	}
	var result = make(map[string]string)
	result[`user`] = string(userInfo)
	result[`tribe_types`] = strings.Join(tribe.Type, `,`)
	return result, nil
}
