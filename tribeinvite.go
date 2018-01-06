package topclient

import (
	"encoding/json"
	"errors"
	"strconv"
)

// 邀请用户加入群接口
type OpenimTribeInviteRequest struct {
	User    *OpenImUser   `json:"user"`
	TribeId int           `json:"tribe_id"`
	Members []*OpenImUser `json:"members"`
}

func (tribe *OpenimTribeInviteRequest) getApiMethodName() string {
	return `taobao.openim.tribe.invite`
}

func (tribe *OpenimTribeInviteRequest) check() (bool, error) {
	if len(tribe.User.Uid) == 0 {
		return false, errors.New(`用户ID不能为空`)
	}
	if len(tribe.User.AppKey) == 0 {
		return false, errors.New(`APPKEY不能为空`)
	}
	if tribe.TribeId == 0 {
		return false, errors.New(`群ID不能为空`)
	}
	for _, item := range tribe.Members {
		if len(item.Uid) == 0 {
			return false, errors.New(`入群用户ID不能为空`)
		}
		if len(item.AppKey) == 0 {
			return false, errors.New(`入群用户APPKEY不能为空`)
		}
	}
	return true, nil
}

func (tribe *OpenimTribeInviteRequest) getApiParas() (map[string]string, error) {
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
	result[`tribe_id`] = strconv.Itoa(tribe.TribeId)
	return result, nil
}
