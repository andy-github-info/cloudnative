package main

import (
	"fmt"
	"strings"
)

func main()  {
	//fmt.Println(os.Getenv("VERSION"))
	s := "127.0.0.1:30081"
	l := []string{"no-cache"}
	fmt.Println(l[0])
	fmt.Println(strings.Split(s, ":")[0])
}