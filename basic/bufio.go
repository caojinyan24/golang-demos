package basic

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// read from file
// go run main.go /home/vv/Documents/app/path/go/README.md
func BufioMain() {
	counts := make(map[string](int))
	for _, fileName := range os.Args[1:] {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println("read err:", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		fmt.Println(n, line)
	}
}

func NetMain() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintln(os.Stderr, "fetch error,", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, "fetch:read err:", url, err)
			os.Exit(1)
		}
		fmt.Println(string(b))
	}
}
