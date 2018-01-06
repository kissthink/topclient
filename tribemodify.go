package topclient

import (
	"encoding/json"
	"errors"
	"strconv"
)

// 创建群接口
type OpenimTribeModifytribeinfoRequest struct {
	User      *OpenImUser `json:"user"`
	TribeName string      `json:"tribe_name"`
	Notice    string      `json:"notice"`
	TribeId   int         `json:"tribe_id"`
}

func (tribe *OpenimTribeModifytribeinfoRequest) getApiMethodName() string {
	return `taobao.openim.tribe.modifytribeinfo`
}

func (tribe *OpenimTribeModifytribeinfoRequest) check() (bool, error) {
	if tribe.TribeId == 0 {
		return false, errors.New(`群ID不能为空`)
	}
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

func (tribe *OpenimTribeModifytribeinfoRequest) getApiParas() (map[string]string, error) {
	userInfo, err := json.Marshal(tribe.User)
	if err != nil {
		return map[string]string{}, err
	}
	var result = make(map[string]string)
	result[`user`] = string(userInfo)
	result[`tribe_name`] = tribe.TribeName
	result[`tribe_id`] = strconv.Itoa(tribe.TribeId)
	result[`notice`] = tribe.Notice
	return result, nil
}
