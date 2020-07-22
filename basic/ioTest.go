package basic

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var inputStr string

func init() {
	flag.StringVar(&inputStr, "name", "value", "usage")
}
func BasicMain() {
	flag.Parse()
	fmt.Println(inputStr)
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("input")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("an error occured", err)
		os.Exit(1)
	} else {
		name := input[:len(input)-1]
		fmt.Println("hello,", name)
	}
	for {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("error", err)
			continue
		}
		input = input[:len(input)-1]
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("bye!")
			os.Exit(0)
		default:
			fmt.Println("sorry")

		}
	}
}
