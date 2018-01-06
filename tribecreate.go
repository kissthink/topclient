package topclient

import (
	"encoding/json"
	"errors"
	"strconv"
)

// 创建群接口
type OpenimTribeCreateRequest struct {
	User      *OpenImUser `json:"user"`
	TribeName string      `json:"tribe_name"`
	Notice    string      `json:"notice"`
	TribeType int         `json:"tribe_type"`
	Members   *OpenImUser `json:"members"`
}

func (tribe *OpenimTribeCreateRequest) getApiMethodName() string {
	return `taobao.openim.tribe.create`
}

func (tribe *OpenimTribeCreateRequest) check() (bool, error) {
	if len(tribe.User.Uid) == 0 {
		return false, errors.New(`用户ID不能为空`)
	}
	if len(tribe.TribeName) == 0 {
		return false, errors.New(`群名称不能为空`)
	}
	if len(tribe.Notice) == 0 {
		return false, errors.New(`群公告不能为空[Notice]`)
	}
	return true, nil
}

func (tribe *OpenimTribeCreateRequest) getApiParas() (map[string]string, error) {
	userInfo, err := json.Marshal(tribe.User)
	if err != nil {
		return map[string]string{}, err
	}
	members, err := json.Marshal(tribe.Members)
	if err != nil {
		return map[string]string{}, err
	}
	var result = make(map[string]string)
	result[`user`] = string(userInfo)
	result[`members`] = string(members)
	result[`tribe_name`] = tribe.TribeName
	result[`tribe_type`] = strconv.Itoa(tribe.TribeType)
	result[`notice`] = tribe.Notice
	return result, nil
}
