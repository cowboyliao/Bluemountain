package main

import (
	"fmt"
	//"strings"
)

func main() {
	var S string
	// S := "20x+x=-20x-2x"
	fmt.Scanln(&S)
	// S = string(S)
	var Box1 int32 = 0   //装未知数的值
	var Box2 int32 = 0   //装常数的值
	var Boxbit uint8 = 0 //每次的小份
	var U uint8 = 0
	Q := 1 //等号前后,前为1,后为-1
	F := 1 //符号,默认为正,-1为负
	// A := 0 //是否有未知数,没有为0,有为1
	// G := 0 //是否此数完毕,没有为0,有为1
	for i := 0; i < len(S); i++ {
		Boxbit = 0
		if S[i] == 61 {
			Q = -1
			continue
		}
		if S[i] == 45 {
			F = -1
			continue
		}
		if S[i] == 43 {
			F = 1
			continue
		}
		if S[i] <= uint8(57) && S[i] >= uint8(49) {
			Boxbit = S[i] - 48
			// fmt.Print(Boxbit)
			if i == len(S)-1 && S[i] <= 57 {
				Box2 += int32(F) * int32(Boxbit) * int32(Q)
			}
			for l := i + 1; l < len(S); l++ {
				if S[l] <= uint8(57) && S[l] >= uint8(48) {
					Boxbit *= 10
					Boxbit += S[l] - 48
					if l == len(S)-1 && S[l] <= 57 {
						Box2 += int32(F) * int32(Boxbit) * int32(Q)
						i = l
					}
				} else if S[l] > uint8(57) {
					if S[l] == 61 {
						// fmt.Printf("%d ", int32(F)*int32(Boxbit)*int32(Q))
						Box2 += int32(F) * int32(Boxbit) * int32(Q)
						Q = -1
						i = l - 1
						F = 1
						break
					}
					U = S[l]
					i = l
					// fmt.Printf("%da ", int32(F)*int32(Boxbit)*int32(Q))
					Box1 += int32(F) * int32(Boxbit) * int32(Q)
					break
				} else {
					Box2 += int32(F) * int32(Boxbit) * int32(Q)
					// fmt.Printf("%d ", int32(F)*int32(Boxbit)*int32(Q))
					i = l - 1
					break
				}
			}
		}
	}
	// fmt.Print("\n", Box1)
	// fmt.Print("\n", Box2, "\n")
	fmt.Printf("%c=%.3f", U, float32(Box2)/(-1.0*float32(Box1)))

}
