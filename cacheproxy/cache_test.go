package cacheproxy

import (
	"fmt"
	"testing"
	"time"
)

func queryDbData(param string) (string, error) {
	return withithCache("queryDbData", 4, func() (s string, err error) {
		fmt.Println("get from db")
		return param, nil
	})

}

func TestCacheProxy(t *testing.T) {
	for i := 0; i < 10; i++ {
		data, err := queryDbData("abc")
		fmt.Println(data, err)
		time.Sleep(1 * time.Second)
	}
}
