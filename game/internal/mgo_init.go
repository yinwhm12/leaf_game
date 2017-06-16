package internal

import (
	"github.com/name5566/leaf/db/mongodb"
	"leaf/log"
)

const DB_INFO  = "mongodb://yin_test:123456@localhost:27017/runmongo"

func init()  {

	db, err := mongodb.Dial(DB_INFO,10)
	if err != nil{
		log.Fatal("db %v",err)
		return
	}
	err =
}
