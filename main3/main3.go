package main

import "fmt"

/*运算符*/
func main() {
	//简单的加减乘除
	var a int = 21
	var b int = 32
	var c int
	c = a + b
	fmt.Println(c)

	//逻辑运算符
	var d = false
	var e = true
	if d && e {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	//&取地址
	fmt.Println(&d)
}
func init() {

}
