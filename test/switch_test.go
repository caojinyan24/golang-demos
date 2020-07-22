package test

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	i := 1
	switch i {
	case 0:
		fmt.Printf("0")
	case 1:
		fmt.Printf("1")
	case 2:
		fallthrough
	case 3:
		fmt.Printf("3")
	case 4, 5, 6:
		fmt.Printf("4, 5, 6")
	default:
		fmt.Printf("Default")
	}
	fmt.Println("test return", getf(i))
}
func getf(i int) int {
	switch i {
	//case 0:
		//return 0
	//case 1:
		//return 1
	case 2,0,1:
		return 2
	case 3:
		//return 3
	case 4, 5, 6:
		return 4
	default:
		return 999
	}
	return 8
}
