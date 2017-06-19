package internal

import (
	"server/msg"
	"gopkg.in/mgo.v2/bson"

	"fmt"
	"github.com/name5566/leaf/log"
)

type User struct {
	Name string
	Pwd string
	Age int
	Address string
}

const USERDB  = "users"

//func init()  {
//	skeletonRegister(&msg.UserLoginInfo{},login)
//	//skeletonRegister(&msg.UserLoginInfo{},register)
//}
//
//func skeletonRegister(m interface{}, h interface{})  {
//	skeleton.RegisterChanRPC(reflect.TypeOf(m),h)
//}

func register(userInfo *msg.RegisterUserInfo)  (err error) {//注册
	//var user User
	//userInfo := args[0].(*msg.RegisterUserInfo)
	skeleton.Go(func() {
		db := mongoDB.Ref()
		defer mongoDB.UnRef(db)
		err := db.DB(DB).C(USERDB).Insert(userInfo)
		if err != nil{
			//log.Fatal("err register --%v",err)
			log.Fatal("err register - %v, err ",err )
			return
 		}
	}, func() {

	})
	return
}

func login(user  *msg.UserLoginInfo)(err error) {
	//var user User
	//fmt.Println("---lognin------",args)
	//user := args[0].(*msg.UserLoginInfo)
	fmt.Println("---userinfo---",user)
	var result User
	skeleton.Go(func() {
		db := mongoDB.Ref()
		defer mongoDB.UnRef(db)
		// check user
		err := db.DB(DB).C(USERDB).Find(bson.M{"name":user.Name,"pwd":user.Pwd}).One(&result)
		if err != nil{
			//log.Fatal("login err - %v",err)
			//a := args[1].(gate.Agent)
			//ChanRPC.Go("LoginAgent",&msg.LoginError{1,"no user"})
			//fmt.Println("---over----?")
			//time.Sleep(15*time.Second)
			//a.WriteMsg(&msg.LoginError{State:-1,Message:"no user"})
			log.Fatal("login err - %v",err)

			return

		}
	}, func() {

	})
	return
}

//检查用户是否已注册过
func checkExitedUser(userName string) (err error){
	//skeleton.Go(func() {
	//	db := mongoDB.Ref()
	//	defer mongoDB.UnRef(db)
	//	err := db.DB(DB).C(USERDB).Find(bson.M{"name":bson.M{"$exists":userName}})
	//	if err != nil {
	//
	//	}
	//},nil)

	db := mongoDB.Ref()
	defer mongoDB.UnRef(db)
	var userInfo msg.RegisterUserInfo
	err = db.DB(DB).C(USERDB).Find(bson.M{"name":userName}).One(&userInfo)
	if err != nil{
		return err
	}
	return nil
}

