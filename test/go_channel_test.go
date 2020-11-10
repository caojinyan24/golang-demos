package test

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {
	out := make(chan int)
	out <- 2
	go f1(out)
}

func f1(in chan int) {
	fmt.Println(<-in)
}
