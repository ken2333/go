package main

import "fmt"

func main() {

	//与运算
	x := 7             //111
	y := 4             //100
	fmt.Println(x & y) //111&100
	//或运算
	fmt.Println(x | y) //111| 100=111
	//异或,只有同等位的数字不相同才为1，否则为
	fmt.Println(x ^ y) //111^100 =011

	//按位取反
	var z uint8 = 7
	fmt.Println(^x) // -8
	fmt.Println(^(8))
	fmt.Println(^z) //248
	x++
	fmt.Println(x)

	switch x {

	case 9, 8:
		{
			fmt.Println(9, "or ", 8)
		}
		fallthrough
	default:
		{
			fmt.Println("defaultS")
		}

	}

}
