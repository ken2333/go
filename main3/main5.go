package main

import "fmt"

/*函数的定义
func function_name(parameter lis)  return_type  {

	return return_type
}
*/
func main() {
	var c int = add(12, 3)
	fmt.Println(c)

}
func add(a int, b int) int {
	return a + b
}
