package topclient

import (
	"errors"
	"strings"
)

type OpenimUsersDeleteRequest struct {
	UserId []string `json:"userids"`
}

func (user *OpenimUsersDeleteRequest) getApiMethodName() string {
	return `taobao.openim.users.delete`
}

func (user *OpenimUsersDeleteRequest) check() (bool, error) {
	if len(user.UserId) == 0 {
		return false, errors.New(`待删除用户ID不能为空`)
	}
	if len(user.UserId) > 100 {
		return false, errors.New(`一次性删除用户最多不能超过100个`)
	}
	return true, nil
}

func (user *OpenimUsersDeleteRequest) getApiParas() (map[string]string, error) {
	var result = make(map[string]string)
	result["userids"] = strings.Join(user.UserId, `,`)
	return result, nil
}
