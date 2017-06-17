package msg

import (
	"github.com/name5566/leaf/network/json"
)
//注册消息内容 即 类型(结构体)
//var Processor network.Processor
var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Hello{})
	Processor.Register(&UserLoginInfo{})
	Processor.Register(&LoginError{})
}

type Hello struct {
	Name string
}

type UserLoginInfo struct {
	Name string
	Pwd	string
}

type LoginError struct {
	State int
	Message string
}