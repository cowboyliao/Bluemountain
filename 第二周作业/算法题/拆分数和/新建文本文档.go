//package main
//
//import "fmt"
//
//func wa(i int) (S string) { //i是剩下的数
//
//	for j := 1; j <= i; j = 5 {
//		ch := string(j + 48)
//		S = ch + "+" + wa(i-j)
//		return
//	}
//	return S
//}
//
//func main() {
//	n := 7
//	for i := 1; i <= int(n/2); i++ {
//		fmt.Print(string(48+i) + "+")
//		fmt.Print(wa(n - i))
//		fmt.Print("\n") //每行完毕即换行
//	}
//}
