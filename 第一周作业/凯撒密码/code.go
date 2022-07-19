package main

import "fmt"

func main() {
	n := 1
	var S string = "qwe"
	fmt.Scanln(&n)
	fmt.Scanln(&S)
	for _, value := range S {
		value = (value+int32(n)-97)%26 + 97
		fmt.Printf("%c", value)

	}

}