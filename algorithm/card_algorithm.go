package algorithm

//统计 牌类型
func CountSimilarCards(a []int,similarMap map[int]int)  {
	length := len(a)
	for index := 0; index < length; index++{
		//if _, ok := similarMap[a[index]]; ok{
		//	similarMap[a[index]]++
		//}else {
		//
		//}
		//fmt.Println("similarCards=",a[index])
		similarMap[a[index]]++
		//if similarMap[a[index]] == 2 {//两个数量
		//	similarMap[2] ++
		//}else if similarMap[a[index]] == 3{//三个 数量
		//	similarMap[3]++
		//	similarMap[2]--
		//}
	}
	//return 0
}
//结果 有多少对 三个的 单的
func EndSum(similarMap map[int]int) (a [4]int) {
	for _, v := range similarMap{
		if v == 0{//没牌 即该类型都配对了
			a[0]++
		}else if v == 1{//剩下 单
			a[1]++
		}else if v == 2{
			a[2]++
		}else if v == 3{
			a[3]++
		}
	}
	return
}
//胡 判断
func CardsHu(a []int,similarMap map[int]int) int {
	count := 0
	for _,v := range a{
		if similarMap[v] == 3{
			if v % 100 == 4{//东南西北
				//count++
				//similarMap[3] --
			}else {
				if similarMap[v+1] > 0 && similarMap[v+2] > 0{
					if similarMap[v+1] == 1&& similarMap[v+2] >= 1{
						if similarMap[v+2] == 1 &&similarMap[v+3] > 0{ //555 6 7 8 d d  后结合
							similarMap[v+1]--
							similarMap[v+2]--
							similarMap[v+3]--
							count++
						}else {//解决 111 2 333 44 55 / 555 6 7 888 前结合
							if similarMap[v+2] == 3{//原来是三个的
								//similarMap[3]--
								//similarMap[2]++
							}
							similarMap[v]--
							similarMap[v+1]--
							similarMap[v+2]--
							count++
							//similarMap[3]--
							//similarMap[2]++
						}
					}else if similarMap[v+1] == 2 && similarMap[v+2] == 2&& similarMap[v+3] == 1{ //555 66 77 8
						similarMap[v]--
						similarMap[v+1]--
						similarMap[v+2]--
						count++
						//similarMap[3]--
						//3个变成2个，所以这里+1 2对的都没了，所以这里-2 所以下面直接-1
						//similarMap[2]--
					}else{//其他情况 不能与其他牌进行有效结合
						//count++
						//similarMap[3] -= 3
					}
				}
			}
		}else if similarMap[v] == 1{
			if similarMap[v+1]  > 0 && similarMap [v+2] > 0{//6 7 8

				//if similarMap[v+1] == 2{//2变1
				//	similarMap[2]--
				//}else if similarMap[v+1] == 3{//3变2
				//	similarMap[3]--
				//	similarMap[2]++
				//}else if similarMap[v+1] == 4{//4变3
				//	similarMap[4]--
				//	similarMap[3]++
				//}
				//
				//if similarMap[v+2] == 2{
				//	similarMap[2]--
				//}else if similarMap[v+2] == 3{
				//	similarMap[3]--
				//	similarMap[2]++
				//}else if similarMap[v+2] == 4{
				//	similarMap[4]--
				//	similarMap[3]++
				//}
				similarMap[v] --
				similarMap[v+1] --
				similarMap[v+2] --
				count++

			}
		}else if similarMap[v] == 2 {
			if similarMap[v+1] >= 2 && similarMap[v+2] >= 2{
				//if similarMap[v+1] == 3{
				//	similarMap[3]--
				//}else if similarMap[v+1] == 4{
				//	similarMap[4]--
				//	similarMap[2]++
				//}
				//if similarMap[v+2] == 3{
				//	similarMap[3]--
				//}else if similarMap[v+2] == 4{
				//	similarMap[4]--
				//	similarMap[2]++
				//}
				similarMap[v] = similarMap[v] - 2
				similarMap[v+1] = similarMap[v+1] - 2
				similarMap[v+2] = similarMap[v+2] - 2
				//similarMap[2] -= 3
				count += 2
			}
		}else if similarMap[v] == 4{
			if similarMap[v+1]>=1 && similarMap[v+2]>=1{
				similarMap[v]--
				similarMap[v+1]--
				similarMap[v+2]--
				count++
			}
		}else {//异常

		}
	}
	return count
}
