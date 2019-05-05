package main

import "fmt"

/*指针
通过&就可以获取变量的内存地址
*/
func main() {
	var a string = "sunyehao"
	var b *string = &a
	fmt.Println(b)
	fmt.Println(&a)
	fmt.Println(*b)
	孙也好 := "123"
	fmt.Println(孙也好)
}
