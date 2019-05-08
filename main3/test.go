package main

import "fmt"

func main() {

	/*	f, err := os.Open("/dex")
		fmt.Println(err)
		fmt.Println(f)
		const sun = 10 - 3
		var s = "123"
		fmt.Println(sun)
		fmt.Println(&s)

		var result = add1(3)
		fmt.Println(result)

		fmt.Printf("type:%T\n",s)
		fmt.Printf("type:%T\n",result)
		fmt.Println(reflect.TypeOf(s))
		fmt.Println(reflect.TypeOf(reflect.TypeOf(s)))
		a:=255
		b:=byte(a)
		fmt.Println("b:",b)
		p:=20
		x:=(*int)(&p)
		fmt.Println(*x)*/
	/*
		var b int =8
		 x:=b>>1
		 c:=8
		y:=8<<2
		 fmt.Println(x)
		 fmt.Println(c<<1)
		fmt.Println(y)

		var d byte=100
		fmt.Println(d)

	*/

	data := "sunyehao"
	for i := range data {

		fmt.Println(i, string(data[i]))

	}
	fmt.Println(data)

	var data1 []int = []int{1, 2, 3, 4, 5, 6}
	for i, value := range data1 {
		fmt.Println(i, value)
		if i == 2 {
			goto exit
		}

	}
exit:
	fmt.Println("退出函数")

}

func add1(num uint64) uint64 {
	if num > 1 {
		return num * add1(num-1)
	} else {
		return 1
	}
}
