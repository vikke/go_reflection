package main

import "log"

type I interface {
	Do(s string) string
}

type S struct {
	S string
}

func (sp *S) Do(s string) string {
	return sp.S + s
}

func main() {
	s := S{S: "abc"}
	tmp := s.Do("foobar")
	log.Println(tmp)
}
