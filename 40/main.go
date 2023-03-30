package main

import "fmt"

func main() {
	//a、b、c 的长度和容量分别是多少
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]
	fmt.Println(len(a), cap(a), len(b), cap(b), len(c), cap(c))
}