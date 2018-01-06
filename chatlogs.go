package topclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// 获取聊天记录接口
type OpenimChatlogsGetRequest struct {
	Begin int64       `json:"begin"`
	End   int64       `json:"end"`
	Count int         `json:"count"`
	User1 *OpenImUser `json:"user1"`
	User2 *OpenImUser `json:"user2"`
}

func (chat *OpenimChatlogsGetRequest) getApiMethodName() string {
	return `taobao.openim.chatlogs.get`
}

func (chat *OpenimChatlogsGetRequest) check() (bool, error) {
	if len(chat.User1.Uid) == 0 || len(chat.User2.Uid) == 0 {
		return false, errors.New(`用户ID不能为空`)
	}
	if chat.Begin == 0 || chat.End == 0 {
		return false, errors.New(`查询时间范围不能为空`)
	}
	if chat.Count == 0 {
		return false, errors.New(`请确定需要查询的总数[count]`)
	}
	return true, nil
}

func (chat *OpenimChatlogsGetRequest) getApiParas() (map[string]string, error) {
	userInfo1, err := json.Marshal(chat.User1)
	if err != nil {
		return map[string]string{}, err
	}
	userInfo2, err := json.Marshal(chat.User2)
	if err != nil {
		return map[string]string{}, err
	}
	var result = make(map[string]string)
	result[`begin`] = strconv.FormatInt(chat.Begin, 10)
	result[`end`] = strconv.FormatInt(chat.End, 10)
	result[`count`] = strconv.Itoa(chat.Count)
	result[`user1`] = string(userInfo1)
	result[`user2`] = string(userInfo2)
	fmt.Println(result)
	return result, nil
}
