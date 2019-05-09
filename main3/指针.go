package main

import "fmt"

func swap(a, b *int) {
	//a,b=b,a
	fmt.Println("交换中")
	fmt.Println("&a:", &a)
	fmt.Println("&b:", &b)
	b, a = a, b
	fmt.Println("交换后")
	fmt.Println("&a:", &a)
	fmt.Println("&b:", &b)

}

func main() {
	x, y := 1, 2
	//交换前
	fmt.Println("&x:", &x)
	fmt.Println("&y:", &y)
	swap(&x, &y)
	//交换后
	fmt.Println("&x:", &x)
	fmt.Println("&y:", &y)
	fmt.Println("x:", x)
	fmt.Println("y:", y)

}
