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

type CardData struct {
	CardId	int	//牌的ID 顺序
	CardType	int	//牌的类型(1--万子 2--同子 3--条子 4--其他 0--初始)
	TotalCount	int //牌的数量 4
	AvailableCardCount	int	//可用的牌数量 0-4
	State	int	//状态(0--可用状态(没有被摸牌 初始) 1--有拥有者(user) 2--游离牌(已打出))
	UserID	int //拥有者 当 State==1时
	OldUserId	int	//曾经拥有者   （谁打出的牌记录)	State == 2
	CardValue	int	//牌的具体值 如 1万
	CardPosition int	//牌的位置 随机获得
}

func initCard(cs []CardData)  {
	//cs = new([]*CardData,136)
	//cs =  []*CardData{}
	//cs = make([]*CardData,136)

	//fmt.Println("in---",len(cs))
	for i := 0; i < 108; i++{
		//cs[i] = &CardData{CardType:( i / 36) + 1,CardValue:(i % 9)+1} //
		cs[i].CardType = ( i / 36) + 1
		cs[i].CardValue = (i % 9)+1 //

	}
	//	东、南、西、北、中、发、白 操作初始
	for i := 108; i< len(cs); i++{
		cs[i].CardType = (i / 36) + 1
		cs[i].CardValue = (i / 4) + 1
		//cs[i] = &CardData{CardType:(i / 36) + 1,CardValue:(i / 4) + 1}
	}
}

func main()  {
	//ExampleLinearContext()
	//fmt.Println(math.Floor(5/4.0))
	//fmt.Println(1/36)
	//fmt.Println(5%4)
	//cs := new([]*CardData,136)
	//cs := make([]*CardData,136)
	//var ss []*CardData = new([]CardData)
	//var cs []*CardData
	cs := make([]CardData,136)
	fmt.Println(len(cs))
	//cs = nil
	initCard(cs)
	for _,v := range cs {
		fmt.Printf("type=%d,value=%d\n",v.CardType,v.CardValue )
	}
}