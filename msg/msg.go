package msg

import (
	"github.com/name5566/leaf/network/json"
)
//注册消息内容 即 类型(结构体)
//var Processor network.Processor
var Processor = json.NewProcessor()

//状态 常量标记
const (
	MSG_Register_Existed	= 0 //注册用户已存在
	MSG_Register_OK		= 1 //注册成功
	MSG_Login_Error		= 2 //登录失败 信息错误
	MSG_Login_OK	= 3 //登录成功

	MSG_DB_Error	= 111 //数据库出错
)

func init() {
	Processor.Register(&Hello{})
	Processor.Register(&UserLoginInfo{})
	Processor.Register(&LoginError{})

	Processor.Register(&RegisterUserInfo{})

	Processor.Register(&CodeState{})
}

type CodeState struct {
	MSG_STATE int // const
}


type Hello struct {
	Name string
}

type UserLoginInfo struct {//登录
	Name string
	Pwd	string
}

type LoginError struct {
	State int
	Message string
}

type RegisterUserInfo struct {//注册
	Name string
	Pwd string
	Age int
	Email string
}