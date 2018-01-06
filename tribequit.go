package topclient

import (
	"encoding/json"
	"errors"
	"strconv"
)

// 退出群接口
type OpenimTribeQuitRequest struct {
	User    *OpenImUser `json:"user"`
	TribeId int         `json:"tribe_id"`
}

func (tribe *OpenimTribeQuitRequest) getApiMethodName() string {
	return `taobao.openim.tribe.quit`
}

func (tribe *OpenimTribeQuitRequest) check() (bool, error) {
	if len(tribe.User.Uid) == 0 {
		return false, errors.New(`用户ID不能为空`)
	}
	if len(tribe.User.AppKey) == 0 {
		return false, errors.New(`APPKEY不能为空`)
	}
	if tribe.TribeId == 0 {
		return false, errors.New(`群ID不能为空`)
	}
	return true, nil
}

func (tribe *OpenimTribeQuitRequest) getApiParas() (map[string]string, error) {
	userInfo, err := json.Marshal(tribe.User)
	if err != nil {
		return map[string]string{}, err
	}
	var result = make(map[string]string)
	result[`user`] = string(userInfo)
	result[`tribe_id`] = strconv.Itoa(tribe.TribeId)
	return result, nil
}
