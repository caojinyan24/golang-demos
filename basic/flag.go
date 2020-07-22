package basic

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", true, "omit trailing newline")
var sep = flag.String("s", " ", "seperator")

func FlagMain() {
	flag.Parse()
	//flag.Args()输出没有指定tag的变量，`go run main.go -n a bc def`中，flag.Args()=[a bc def]
	fmt.Println(flag.Args())
	fmt.Println(*sep, *n)
	fmt.Println(strings.Join(flag.Args(), *sep))
	fmt.Println("!n:", !*n)
	if !*n {
		fmt.Println("end")
	}
}
