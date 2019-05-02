package main

import "fmt"

/*

控制流语句
*/
func main() {
	//if语句
	if 1 > 0 {
		fmt.Println("1>0")
	} else if 2 > 0 {
		fmt.Print("2>0")
	} else {
		fmt.Println("xxx")
	}

	//switch
	x := 200
	switch {
	case x > 200:
		{
			fmt.Println("x>200")
		}
	default:
		{
			fmt.Println("defalut")
		}

	}

}
