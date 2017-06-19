package main

import (
	"github.com/name5566/leaf/go"
	"time"
	"fmt"
)

func ExampleLinearContext()  {
	d := g.New(10)//chan 的申请 并初始化 10容量

	// 随机 执行
	d.Go(func() {
		//time.Sleep(time.Second * 1)
		fmt.Println("1")
	},nil)

	d.Go(func() {
		//time.Sleep(time.Second / 2)
		fmt.Println("2")
	},nil)

	//d.Go(func() {
	//	//time.Sleep(time.Second / 2)
	//	fmt.Println("3")
	//},nil)

	//似乎清除 之前 的协程 避免堵塞 或者没必要的 顺序
	d.Cb(<-d.ChanCb)
	d.Cb(<-d.ChanCb)
	//d.Cb(<-d.ChanCb)


	//线性 执行 容器 申请 不管 申请顺序 或者 有无挂起 都要顺序进行(挂起者等待后顺序执行)
	c := d.NewLinearContext()
	c.Go(func() {
		//time.Sleep(time.Second /2)
		fmt.Println("11")
	},nil)

	c.Go(func() {
		time.Sleep(time.Second / 2)
		fmt.Println("21")
	},nil)

	c.Go(func() {
		//time.Sleep(time.Second )
		fmt.Println("31")
	},nil)

	d.Close()
}

func main()  {
	ExampleLinearContext()

}