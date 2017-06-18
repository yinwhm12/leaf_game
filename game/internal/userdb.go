package internal

import (
	"server/msg"
	"reflect"
	"gopkg.in/mgo.v2/bson"

	"fmt"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/gate"
)

type User struct {
	Name string
	Pwd string
	Age int
	Address string
}

const USERDB  = "users"

func init()  {
	skeletonRegister(&msg.UserLoginInfo{},login)
	//skeletonRegister(&msg.UserLoginInfo{},register)
}

func skeletonRegister(m interface{}, h interface{})  {
	skeleton.RegisterChanRPC(reflect.TypeOf(m),h)
}

func register(args []interface{})   {//注册
	//var user User
	userInfo := args[0].(*msg.UserLoginInfo)
	skeleton.Go(func() {
		db := mongoDB.Ref()
		defer mongoDB.UnRef(db)
		err := db.DB(DB).C(USERDB).Insert(userInfo)
		if err != nil{
			//log.Fatal("err register --%v",err)
			log.Fatal("err register - %v, err ",err )

		}
	}, func() {

	})
	return
}

func login(args []interface{}) {
	//var user User
	fmt.Println("---lognin------",args)
	user := args[0].(*msg.UserLoginInfo)
	fmt.Println("---userinfo---",user)
	var result User
	skeleton.Go(func() {
		db := mongoDB.Ref()
		defer mongoDB.UnRef(db)
		// check user
		err := db.DB(DB).C(USERDB).Find(bson.M{"name":user.Name,"pwd":user.Pwd}).One(&result)
		if err != nil{
			//log.Fatal("login err - %v",err)
			a := args[1].(gate.Agent)
			//ChanRPC.Go("LoginAgent",&msg.LoginError{1,"no user"})
			//fmt.Println("---over----?")
			//time.Sleep(15*time.Second)
			a.WriteMsg(&msg.LoginError{State:-1,Message:"no user"})
			log.Fatal("login err - %v",err)

			return

		}
	}, func() {

	})
	return
}
