package main

import "fmt"

//通过range来便利数组
func main() {
	nums := []int{1, 2, 3, 4}

	for key, num := range nums {
		fmt.Println(key, num)
	}

	strs := map[string]string{"string": "sss", "name": "sun"}
	for key, value := range strs {
		fmt.Println(key, value)
	}

	delete(strs, "string")
	fmt.Println(strs)

}
