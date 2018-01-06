package topclient

import "errors"

// 设置需要过滤的铭感词
type OpenimSnfilterwordSetfilterRequest struct {
	Creator string `json:"creator"`
	Word    string `json:"word"`
	Desc    string `json:"desc"`
}

func (filter *OpenimSnfilterwordSetfilterRequest) getApiMethodName() string {
	return `taobao.openim.snfilterword.setfilter`
}

func (filter *OpenimSnfilterwordSetfilterRequest) check() (bool, error) {
	if len(filter.Creator) == 0 {
		return false, errors.New(`敏感词添加者不能为空[Creator]`)
	}
	if len(filter.Word) == 0 {
		return false, errors.New(`敏感词不能为空`)
	}
	return true, nil
}

func (filter *OpenimSnfilterwordSetfilterRequest) getApiParas() (map[string]string, error) {
	var result = make(map[string]string)
	result[`creator`] = filter.Creator
	result[`filterword`] = filter.Word
	result[`desc`] = filter.Desc
	return result, nil
}
