package main

import "fmt"

/*
	目前问题:结构体如何传参,是复制还是引用
			如何定义全局数组
			如何将效率提高
*/

type P struct{
	X int
	Y int
	T int//toward
}

func main() {
	S := `*...*.....
		......*...
		...*...*..
		..........
		...*.F....
		*.....*...
		...*......
		..C......*
		...*.*....
		.*.*......`
	/////////////////////////构建地图//////////////////
	// 首先make第一维的大小
	var arr = make([][]int, 10)
	// 然后分别对其中的进行make
	for i := range arr {
		arr[i] = make([]int, 10)
	}
	Index := 0
	for _ /*index*/, value := range S {
		if value == 9 {
			continue
		}
		if value == 10 {
			continue
		}
		arr[Index/10][Index%10] = int(value)
		// fmt.Println(Index)
		Index += 1
		// fmt.Println(value)
	}
	fmt.Println(arr)
	// fmt.Printf("%d","\n")
	////////////找位置//////////////////////////////
	var KEY1 bool = true
	var KEY2 bool = true
	X, Y := 0, 0 //约翰的位置
	x, y := 0, 0 //牛的位置
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if arr[i][j] == 70 && KEY1 {
				X = j
				Y = i
				KEY1 = false
			}

			if arr[i][j] == 67 && KEY2 {
				x = j
				y = i
				KEY2 = false
			}
			if !KEY1 && !KEY2  {
				break
			}
		}
	}
	pp:=P{X,Y,1}
	cp:=P{x,y,1}
	T,t:=0,0//朝向
	// fmt.Print(X, " ", Y)
	// fmt.Print(x, " ", y)
	//每分钟进行回合制游戏
	count:=-1
	for i:=1;;{
		pp=movement(pp)
		cp=movement(cp)
		if pp.X==cp.x && pp.Y==cp.y{
			count=i
			break
		}
		i+=1
		if i>1000{
			count=-1
			break
		}
	}

	fmt.Print(count)
}

//////////////////无痛人牛//////////////////
/*
	How cow movement is the same that john movement
	简单理解就是一个回合制游戏
	且运动方式一样,每回合都往前一步,如果没有前路就右转花费一回合
	GameOver的条件就是回合结束时人和牛在一个格上
*/
func movement(position P)(nposition P){





}
/*
	T 1 north
	T 2 east
	T 3 south
	T 4 west
*/
func checkout(p P)i bool{
switch p.T{
case 1:p.Y-=1
case 2:p.X-=1
case 3:p.Y+=1
case 4:p.X+=1
}



return true
}
