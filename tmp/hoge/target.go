package hoge

import "log"

type MyStruct struct {
}

type ReturnValue struct {
	val string
}

func (ms *MyStruct) Method1(i int, s string) (out ReturnValue) {
	log.Printf("foobar")
	return ReturnValue{val: s + s}
}
