package topclient

import "encoding/json"

type OpenimUsersAddRequest struct {
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

func (user *OpenimUsersAddRequest) check() (bool, error) {
	return true, nil
}

func (user *OpenimUsersAddRequest) getApiMethodName() string {
	return `taobao.openim.users.add`
}

func (user *OpenimUsersAddRequest) getApiParas() (map[string]string, error) {
	jsonData, err := json.Marshal(user)
	if err != nil {
		return map[string]string{}, err
	}
	var result = make(map[string]string)
	result[`userinfos`] = string(jsonData)
	return result, nil
}
