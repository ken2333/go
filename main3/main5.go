package main

import "fmt"

/*函数的定义
func function_name(parameter lis)  return_type  {

	return return_type
}

返回多个变量
func  function_name(parameter list)  ( int ,sun)
{
 return var1,var2
}

还可以返回函数
func  function_name( list)  func()
{
    return func (){

}
}

*/
func main() {
	var c int = add(12, 3)
	fmt.Println(c)
	fmt.Println(add2(2, 4))

	//调用返回的函数
	var f func() = add3(2, 3)
	f()
}
func add(a int, b int) int {
	return a + b
}

func add2(a int, b int) (int, int) {

	return 2, a + b
}

//返回函数
func add3(x int, b int) func() {

	return func() {
		fmt.Println(x + b)
	}
}
