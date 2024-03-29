package test

import (
	"errors"
	"fmt"
	"testing"
)

func Defer() (result string) {
	result = "init"
	defer func() {
		if result == "" {
			result = "defer"
		}
	}()
	return "add " + result
}
func DeferBlank() (result string) {
	result = "init"
	defer func() {
		fmt.Println(result)
		result = "defer"
		fmt.Println(result)
	}()
	return "add " + result
}
func DeferWithParam() string {
	var result = "init"
	defer func() {
		result = "defer"
	}()
	return result + " added"
}
func TestDefer(t *testing.T) {
	//println(Defer())
	println(DeferBlank())
	//println(DeferWithParam())

}

const PLATFORM_APP_IDENTITY_TYPE_UID = 10

func DeferMain() {
	//fmt.Println("###1")
	testDefer1()
	//fmt.Println("###2")
	//testDefer2()
	//fmt.Println("###panic-recover")
	// recoverFromPanic1()
}

// defer
func testDefer1() {
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func testDefer2() {
	for i := 0; i < 5; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func raisePanic() {
	fmt.Println("raising")
	panic(errors.New("raise a panic"))
}

func recoverFromPanic1() {
	fmt.Println("begin recover")
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("before raise panic")
			//raisePanic()
			fmt.Println("after raise panic")
		}
	}()
	raisePanic()
	fmt.Println("after recover")
}

func recoverFromPanic2() {
	fmt.Println("begin recover")
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("before raise panic")
			raisePanic()
			fmt.Println("after raise panic")
		}
	}()
	fmt.Println("after recover")
}
