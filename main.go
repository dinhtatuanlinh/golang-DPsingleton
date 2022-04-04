package main

import (
	"DPsingleton/singleton"
	"fmt"
)

func main() {
	s1 := singleton.GetInstance()
	s2 := singleton.GetInstance()
	fmt.Println(s1.AddOne())
	fmt.Println(s1.AddOne())
	fmt.Println(s2.AddOne())
	fmt.Println(s2.AddOne())
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", s2)

}