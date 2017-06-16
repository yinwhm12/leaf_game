package msg

import (
	"github.com/name5566/leaf/network/json"
)
//注册消息内容 即 类型(结构体)
//var Processor network.Processor
var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Hello{})
	Processor.Register(&UserTest{})
}

type Hello struct {
	Name string
}

type UserTest struct {
	Name string
	Pwd	string
}