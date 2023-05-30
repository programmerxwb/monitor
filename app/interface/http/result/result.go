package result

import (
	"encoding/json"
)

var (
	OK = response(200, "success")
)

type Response struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (res *Response) WithMsg(message string) Response {
	return Response{
		Msg:  message,
		Data: res.Data,
	}
}

func (res *Response) WithData(data any) Response {
	return Response{
		Msg:  res.Msg,
		Data: data,
	}
}

func (res *Response) ToString() string {
	err := &struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{
		Msg:  res.Msg,
		Data: res.Data,
	}
	raw, _ := json.Marshal(err)
	return string(raw)
}

func response(code int, msg string) *Response {
	return &Response{
		Msg:  msg,
		Data: nil,
	}
}
