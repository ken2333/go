package main

import "fmt"

/*

	接口的定义和使用
*/

type myinterface interface {
	get(str string) string
	post(str string) string
}

type dosometing struct {
}

func (this *dosometing) get(str string) string {
	fmt.Println(str)
	return str
}

func (this *dosometing) post(str string) string {
	fmt.Println(str)
	return str
}

func main() {
	var my myinterface = new(dosometing)
	var result1 = my.post("post")
	var result2 = my.get("get")
	fmt.Println("result1:" + result1)
	fmt.Println("result2:" + result2)

}
