package topclient

import (
	"errors"
	"strings"
)

type OpenimUsersGetRequest struct {
	UserId []string `json:"userids"`
}

type OpenimUsersGetResp struct {
	OpenimUsersGetResponse struct {
		Userinfos struct {
			Userinfos []struct {
				Address     string `json:"address"`
				Age         int    `json:"age"`
				Career      string `json:"career"`
				Email       string `json:"email"`
				Extra       string `json:"extra"`
				Gender      string `json:"gender"`
				GmtModified string `json:"gmt_modified"`
				IconURL     string `json:"icon_url"`
				Mobile      string `json:"mobile"`
				Name        string `json:"name"`
				Nick        string `json:"nick"`
				Password    string `json:"password"`
				Qq          string `json:"qq"`
				Remark      string `json:"remark"`
				Status      int    `json:"status"`
				Taobaoid    string `json:"taobaoid"`
				Userid      string `json:"userid"`
				Vip         string `json:"vip"`
				Wechat      string `json:"wechat"`
				Weibo       string `json:"weibo"`
			} `json:"userinfos"`
		} `json:"userinfos"`
	} `json:"openim_users_get_response"`
}

func (user *OpenimUsersGetRequest) getApiMethodName() string {
	return `taobao.openim.users.get`
}

func (user *OpenimUsersGetRequest) check() (bool, error) {
	if len(user.UserId) == 0 {
		return false, errors.New(`待删除用户ID不能为空`)
	}
	if len(user.UserId) > 100 {
		return false, errors.New(`一次性删除用户最多不能超过100个`)
	}
	return true, nil
}

func (user *OpenimUsersGetRequest) getApiParas() (map[string]string, error) {
	var result = make(map[string]string)
	result["userids"] = strings.Join(user.UserId, `,`)
	return result, nil
}
