package internal

import (
	"github.com/name5566/leaf/db/mongodb"
	"leaf/log"
)

const DB_INFO  = "mongodb://yin_test:123456@localhost:27017/runmongo"
var mongoDB *mongodb.DialContext

const DB  = "runmongo"

func init()  {

	db, err := mongodb.Dial(DB_INFO,10)
	if err != nil{
		//fmt.Println("----connecting----")
		log.Fatal("db %v",err)
		//fmt.Println("------connected----")
		return
	}
	mongoDB = db
	//fmt.Println("------connected----")
}

func mongoDBDestroy()  {
	mongoDB.Close()
	mongoDB = nil

}
