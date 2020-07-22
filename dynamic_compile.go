package main

import (
	"fmt"
)

//编译文件： go build -buildmode=plugin dynamic_compile.go
//运行： go run main.go
func PrintFunc() {
	fmt.Println("hi, i am here!")
}
func Angry() {
	fmt.Println("now, i am angry! leave me alone!")
}
func SayAngry() string {
	return "now, i am angry! leave me alone!"
}
func DoAngryThings() interface{} {
	return struct {
		Type        string
		Consequence string
	}{
		Type:        "shout",
		Consequence: "make people angry too",
	}
}
