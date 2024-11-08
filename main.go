package main

import (
	"fmt"

	"github.com/kwdowicz/lgr/lgr"
)

type MyStruct struct{}

func (s *MyStruct) MyMethod1() {
	fmt.Println("From method 1")
}

func (s *MyStruct) MyMethod2() {
	fmt.Println("From method 2")
}

func main() {
	s := &MyStruct{}
	lgr.LogMethods(s)
	s.MyMethod1()
	s.MyMethod2()
}