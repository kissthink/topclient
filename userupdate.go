package topclient

import (
	"encoding/json"
	"errors"
)

type OpenimUsersUpdateRequest struct {
	Nick     string `json:"nick"`
	IconUrl  string `json:"icon_url"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Taobaoid string `json:"taobaoid"`
	Userid   string `json:"userid"`
	Password string `json:"password"`
	Remark   string `json:"remark"`
	Extra    string `json:"extra"`
	Vip      string `json:"vip"`
	Address  string `json:"address"`
	Name     string `json:"name"`
	Age      string `json:"age"`
	Gender   string `json:"gender"`
	Wechat   string `json:"wechat"`
	Qq       string `json:"qq"`
	Weibo    string `json:"webo"`
}

func (user *OpenimUsersUpdateRequest) check() (bool, error) {
	if len(user.Userid) == 0 {
		return false, errors.New(`USERID不能为空`)
	}
	return true, nil
}

func (user *OpenimUsersUpdateRequest) getApiMethodName() string {
	return `taobao.openim.users.update`
}

func (user *OpenimUsersUpdateRequest) getApiParas() (map[string]string, error) {
	jsonData, err := json.Marshal(user)
	if err != nil {
		return map[string]string{}, err
	}
	var result = make(map[string]string)
	result[`userinfos`] = string(jsonData)
	return result, nil
}
