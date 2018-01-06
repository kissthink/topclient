package topclient

import "encoding/json"

// IM标准消息发送
type OpenimImmsgPushRequest struct {
	FromUser   string       `json:"from_user"`
	ToUsers    []string     `json:"to_users"`
	MsgType    int          `json:"msg_type"` //消息类型。0:文本消息。1:图片消息，只支持jpg、gif。2:语音消息，只支持amr。8:地理位置信息。
	Context    string       `json:"context"`  //发送的消息内容。根据不同消息类型，传不同的值。0(文本消息):填消息内容字符串。1(图片):base64编码的jpg或gif文件。3(语音):base64编码的amr文件。8(地理位置):经纬度，格式如 111,222
	ToAppKey   string       `json:"to_appkey"`
	MediaAttr  *ImMediaAttr `json:"media_attr"`
	FromTaobao string       `json:"from_taobao"`
}

type ImMediaAttr struct {
	Type     string `json:"type"`     // 多媒体文件格式
	PlayTime int    `json:"playtime"` // 音频文件为时长，图片文件为空
}

func (msg *OpenimImmsgPushRequest) getApiMethodName() string {
	return `taobao.openim.immsg.push`
}

func (msg *OpenimImmsgPushRequest) check() (bool, error) {
	return true, nil
}

func (msg *OpenimImmsgPushRequest) getApiParas() (map[string]string, error) {
	msgInfo, err := json.Marshal(msg)
	if err != nil {
		return map[string]string{}, err
	}
	var result = make(map[string]string)
	result[`immsg`] = string(msgInfo)
	return result, nil
}
