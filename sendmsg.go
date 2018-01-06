package topclient

import (
	"encoding/json"
	"errors"
	"strconv"
)

// 发送群消息接口
type OpenimTribeSendmsgRequest struct {
	User    *OpenImUser `json:"user"`
	TribeId int         `json:"tribe_id"`
	Msg     *TribeMsg   `json:"msg"`
}

// 群消息结构体
type TribeMsg struct {
	AtFlag     string           `json:"at_flag"`   //是否是at消息， 0表示不是at消息，1表示at指定的用户，2表示at群里所有人
	AtMembers  []*OpenImUser    `json:"atmembers"` //当at_flag=1时，必须指定at的用户
	CustomPush *CustomPush      `json:"custom_push"`
	MediaAttrs *TribeMediaAttrs `json:"media_attrs"`
	MsgContent string           `json:"msg_content"`
	MsgType    string           `json:"msg_type"` //消息类型，0 表示普通文本消息；2 表示语音消息；16表示图片消息；17表示用户自定义消息
	Push       string           `json:"push"`     // 该消息是否需要push[true|false]
}

// 群消息多媒体数据结构
type TribeMediaAttrs struct {
	Height   int    `json:"height"`
	Width    int    `json:"width"`
	Type     string `json:"type"`     // 多媒体文件格式
	PlayTime int    `json:"playtime"` // 音频文件为时长，图片文件为空
}

//自定义push提醒参数，格式为json字符串，该参数为空时，采用系统默认的push
type CustomPush struct {
	D     string `json:"d"`
	Sound string `json:"sound"`
	Title string `json:"title"`
}

func (msg *OpenimTribeSendmsgRequest) getApiMethodName() string {
	return `taobao.openim.tribe.sendmsg`
}

func (msg *OpenimTribeSendmsgRequest) check() (bool, error) {
	if len(msg.User.Uid) == 0 {
		return false, errors.New(`用户ID不能为空`)
	}
	if len(msg.User.AppKey) == 0 {
		return false, errors.New(`APPKEY不能为空`)
	}
	if msg.TribeId == 0 {
		return false, errors.New(`群ID不能为空`)
	}
	if len(msg.Msg.AtFlag) == 0 {
		return false, errors.New(`请指定是否需要@群成员`)
	}
	if len(msg.Msg.MsgContent) == 0 {
		return false, errors.New(`发送的消息不能为空[MsgContent]`)
	}
	return true, nil
}

func (msg *OpenimTribeSendmsgRequest) getApiParas() (map[string]string, error) {
	msgInfo, err := json.Marshal(msg.Msg)
	if err != nil {
		return map[string]string{}, err
	}
	userInfo, err := json.Marshal(msg.User)
	if err != nil {
		return map[string]string{}, err
	}
	var result = make(map[string]string)
	result[`msg`] = string(msgInfo)
	result[`user`] = string(userInfo)
	result[`tribe_id`] = strconv.Itoa(msg.TribeId)
	return result, nil
}
