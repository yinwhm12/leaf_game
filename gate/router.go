package gate

import (
	"server/msg"
	"server/game"
)

//消息在此进行交割
func init() {
	msg.Processor.SetRouter(&msg.Hello{},game.ChanRPC)//参数消息内容 通信桥chanRPC
}
