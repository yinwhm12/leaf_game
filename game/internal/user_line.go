package internal

import (
	"leaf/gate"
	"github.com/name5566/leaf/go"
)

type UserLine	struct {
	gate.Agent //申请代理
	*g.LinearContext
	Cards	[]*CardData //牌
	User	*UserData
}


