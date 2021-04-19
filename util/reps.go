package util

import (
	"github.com/kataras/iris/v12"
	"os/exec"
	"strings"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func Ok() *Response {
	cmd := exec.Command("node", "cxh.js")
	var chouxiang, _ = cmd.Output()
	c := strings.Split(Bytes2String(chouxiang), "\n")
	return &Response{
		Code: 0,
		Data: iris.Map{
			"msg":     c[1],
			"msg_raw": c[0],
		},
	}
}

func OkData(data interface{}) *Response {
	cmd := exec.Command("node", "cxh.js")
	var chouxiang, _ = cmd.Output()
	c := strings.Split(Bytes2String(chouxiang), "\n")
	return &Response{
		Code: 0,
		Data: iris.Map{
			"ass":    data,
			"msg":     c[1],
			"msg_raw": c[0],
		},
	}
}

func Error(data interface{}) *Response {
	return &Response{
		Code: 1,
		Data: data,
	}
}
