package main

import (
	"github.com/lifestyledesign/tmp/hoge"
	"log"
	"reflect"
)

func main() {
	ms := &hoge.MyStruct{}
	receiver := reflect.ValueOf(ms)
	log.Println("before")
	receiver.MethodByName("Method1").Call([]reflect.Value{})
	log.Println("after")
}
