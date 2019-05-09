package main

import "fmt"

func hello() {
	fmt.Println("hello ,woold!")
}
func exec(f func()) {
	f()
}
func main() {
	f := func() {
		fmt.Println("匿名函数")
	}
	exec(f)
	exec(func() {
		fmt.Println("匿名函数1")
	})

	x := 10
	fmt.Println("x执行前的地址：", &x)
	test(&x)
	fmt.Println(x)

	y := 10
	fmt.Println("y执行前的地址", &y)
	test2(y)

	var p *int
	test3(&p)
	fmt.Println(*p)

	fu := func() {
		fmt.Println("执行匿名参数")
	}
	fu()

	i := 20
	x = i
	fmt.Println(&x)
	fmt.Println(&i)

}

/*传入的是一个整数型的指针，在修改x会导致传入参数的修改*/
func test(x *int) {
	fmt.Println(x)
	*x = *x * *x
}

/*传入的是一个值拷贝传递，与指针不同的是，指针拷贝后依然指向
同一个地址，而基本类型的拷贝就会在函数内部新建一个，
主要函数内的参数地址和传入的参数地址是不同的。
*/
func test2(x int) {
	fmt.Println(&x)
	x = x * x
}

func test3(p **int) {
	x := 100
	*p = &x
}

func test4(a, b *int) {

	fmt.Println(a)
	fmt.Println(b)
}
