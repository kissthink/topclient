package topclient

import (
	"encoding/json"
	"errors"
	"strconv"
)

// 设置群管理员接口
type OpenimTribeSetmanagerRequest struct {
	User    *OpenImUser `json:"user"`
	TribeId int         `json:"tribe_id"`
	Member  *OpenImUser `json:"member"`
}

func (tribe *OpenimTribeSetmanagerRequest) getApiMethodName() string {
	return `taobao.openim.tribe.setmanager`
}

func (tribe *OpenimTribeSetmanagerRequest) check() (bool, error) {
	if len(tribe.User.Uid) == 0 || len(tribe.Member.Uid) == 0 {
		return false, errors.New(`用户ID不能为空`)
	}
	if len(tribe.User.AppKey) == 0 || len(tribe.Member.AppKey) == 0 {
		return false, errors.New(`APPKEY不能为空`)
	}
	if tribe.TribeId == 0 {
		return false, errors.New(`群ID不能为空`)
	}
	return true, nil
}

func (tribe *OpenimTribeSetmanagerRequest) getApiParas() (map[string]string, error) {
	userInfo, err := json.Marshal(tribe.User)
	if err != nil {
		return map[string]string{}, err
	}
	member, err := json.Marshal(tribe.Member)
	if err != nil {
		return map[string]string{}, err
	}
	var result = make(map[string]string)
	result[`user`] = string(userInfo)
	result[`member`] = string(member)
	result[`tribe_id`] = strconv.Itoa(tribe.TribeId)
	return result, nil
}
