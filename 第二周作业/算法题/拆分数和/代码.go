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
package main

import "fmt"

var a = [10001]int{1}
var n int

func search(s int, t int) {
	for i := a[t-1]; i <= s; i++ {
		if i < n {
			a[t] = i
			s -= i
			if s == 0 {
				print(t)
			} else {
				search(s, t+1)
			}
			s += i
		}
	}
}
func print(t int) {
	for i := 1; i <= t-1; i++ {
		fmt.Printf("%d+", a[i])
	}
	fmt.Printf("%d\n", a[t])
}
func main() {
	fmt.Scanln(&n)
	search(n, 1)
	return
}
