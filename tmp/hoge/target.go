package hoge

import "log"

type MyStruct struct {
}

func (ms *MyStruct) method1() {
	log.Printf("foobar")
}
