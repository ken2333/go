package main

import "fmt"

func main() {

	test5("1", "@", 1, 2, 3, 4, 5, 6)
}

/*变参只能放到参数的最后一位，同时只能是同一个类型*/

func test5(x, y string, arr ...int) {

	fmt.Printf("%T\n", arr)
	fmt.Println(arr)
}

/*retsult:
[]int
[1 2 3 4 5 6]
*/

func test6() (int, s string, e error) {

	return "", "s", nil
}
