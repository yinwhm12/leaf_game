package internal

import (
	"github.com/name5566/leaf/gate"
	"fmt"
)

func init() {//与gate 进行"交流"
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC("LoginAgent", rpcLoginAgent)
}

func rpcNewAgent(args []interface{}) {
	fmt.Println("--rpcNew--",args)
	a := args[0].(gate.Agent)
	fmt.Println("args[0]:",a)
	fmt.Println("len():",len(args))
	for i := 0; i < len(args); i++{
		//fmt.Fprintln("i=%d,arg[%d]=%v",i,i,args[i])
		fmt.Printf("i=%d,arg[%d]=%v \n",i,i,args[i] )
	}

	_ = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

func rpcLoginAgent(args []interface{})  {
	fmt.Println("-rpclon-:",args)
	m := args[0]
	fmt.Println("get m--:",m)
	fmt.Println("len--:",len(args))
}
