package basic

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSyncMap(t *testing.T) {
	var localCache = sync.Map{}
	i := 0
	for {
		if i > 50 {
			break
		}
		go func() {
			i = i + 1
			localCache.Store(fmt.Sprintf("%v", i), i)
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(localCache)
	localCache.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		localCache.Delete(key)
		return true
	})
	fmt.Println("@@@")
	//fmt.Println(localCache)
	localCache.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		//localCache.Delete(key)
		return true
	})

}
