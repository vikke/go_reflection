package main

import (
	"github.com/reflection/tmp/hoge"
	"log"
	"reflect"
)

func main() {
	ms := &hoge.MyStruct{}
	receiver := reflect.ValueOf(ms)
	log.Println("before")
	receiver.MethodByName("Method1").Call([]reflect.Value{reflect.ValueOf(1), reflect.ValueOf("aa")})

}
