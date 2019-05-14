package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	s := "123321续页"
	var s1 string
	fmt.Println(len(s))
	fmt.Println(s)
	fmt.Println(s1 == "")
	fmt.Println(len(s1))
	s3 := `asdasdsad
	第二段`
	fmt.Println(s3)
	s4 := `"123321"`
	fmt.Println(s4)

	for i := 0; i < len(s4); i++ {
		fmt.Printf("%d:[%c]\n", s4[i])
	}

	s5 := []byte(s4)
	fmt.Println(s5)
	s6 := string(s5)
	fmt.Println(s6)

	t1 := time.Now()

	var str1 string
	for i := 0; i < 1000; i++ {
		str1 += "a"
	}
	fmt.Println(str1)

	t2 := time.Now()
	fmt.Println(t2.Unix() - t1.Unix())

	s2 := make([]string, 1000)

	for i := 0; i < 1000; i++ {
		s2[i] = "a"
	}
	fmt.Println(strings.Join(s2, ""))

}
