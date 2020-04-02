package main

import (
	"log"
	"reflect"
)

type MyStruct struct {
}

type ReturnValue struct {
	val string
}

func (ms *MyStruct) Method1(i int, s string) (out ReturnValue) {
	log.Printf("foobar")
	return ReturnValue{val: s + s}
}
func main() {
	ms := &MyStruct{}
	receiver := reflect.ValueOf(ms)
	log.Println("before")
	receiver.MethodByName("Method1").Call([]reflect.Value{reflect.ValueOf(1), reflect.ValueOf("aa")})

}
