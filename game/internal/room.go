package internal

import (
	"github.com/name5566/leaf/go"
	"server/msg"
	"time"
	"github.com/name5566/leaf/gate"
)

type position int // 位置 1-4

type Room struct {
	RoomData *RoomData
	MapUsers	map[int]*UserData //玩家信息以及座位
	//RoomState	int	//房间状态
	RoomOwner	*UserData //房管
	*g.LinearContext
	gate.Agent
	Players	int	//玩家数量
	//RoomTime	int //房间有效期限

}

//初始化房间信息 创建房间
func InitRoom(msgRoomInfo *msg.RoomInfo,roomOwner *UserData) (room *Room) {
	r := RoomData{
		RoomName:msgRoomInfo.RoomName,
		Volume:msgRoomInfo.Volume,
		GameType:msgRoomInfo.GameType,
		PayValue:msgRoomInfo.PayValue,
		BaseMoney:msgRoomInfo.BaseMoney,
		RoomPwd:msgRoomInfo.RoomPwd,
		CreatedAt:time.Now().UnixNano(),
		RoomState:1,
	}
	room.Players = 1
	room.RoomData = &r
	room.RoomOwner = roomOwner
	room.MapUsers[1]= roomOwner//房主做一号位 创建房间的
	room.LinearContext = skeleton.NewLinearContext()
	return room
}

//密码检查
func (r *Room)CheckPlayerAndPwd(pwd *msg.RoomPWDJoinCondition) bool {
	if r.RoomData.RoomPwd == pwd.Pwd && r.Players < 4{//密码正确 切 不满人
		return true
	}else {
		r.WriteMsg(&msg.CodeState{msg.MSG_ROOM_NOTAUTH,"密码错误"})
		return false
	}
}

//修改房间基本信息 仅有房主修改
func ChangeRoomInfo(userLine *UserLine,msgRoomInfo *msg.RoomInfo)  {
	if userLine.RoomLine.RoomOwner != userLine.UserData{//不是房主
		r := RoomData{
			RoomName:msgRoomInfo.RoomName,
			Volume:msgRoomInfo.Volume,
			GameType:msgRoomInfo.GameType,
			PayValue:msgRoomInfo.PayValue,
			BaseMoney:msgRoomInfo.BaseMoney,
			RoomPwd:msgRoomInfo.RoomPwd,
			CreatedAt:time.Now().UnixNano(),
			//RoomState:0,
		}
		userLine.RoomLine.RoomData = &r
	}else {//不是房主
		userLine.WriteMsg(&msg.CodeState{msg.MSG_ROOM_NOTAUTH,"你不是房主"})
	}
}

//加入房间
func JoinRoom(userLine *UserLine,room *Room)  {
	if userLine.RoomLine == nil{//确定是没有房间的人 可以进入
	//	分配位置 并进行初始赋值
		 room.Players = room.Players + 1
		//room.MapUsers[n] = userLine.UserData
		userLine.RoomLine.MapUsers[room.Players] = userLine.UserData
	}else {
		userLine.WriteMsg(&msg.CodeState{msg.MSG_ROOM_OVERVOLUME,"你已在其他房间进行游戏了，请退出房间，后操作！"})
		return
	}
}


//加入房间前 检查是否 有“资本” 进入
func CheckConditionForJoining(userLine *UserLine,room *Room)  {
	 if userLine.UserData.Money >= room.RoomData.PayValue{

	 }else {
		 userLine.WriteMsg(&msg.CodeState{msg.MSG_ROOM_NOMONEY,"你的资金不足，该房间的要求!"})
		 return
	 }
}



//解散(关闭)房间