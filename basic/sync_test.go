package basic

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSync(t *testing.T) {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go incrCountSync()
	}
	wg.Wait()
	fmt.Println(counter)
}
func TestSyncGroup(t *testing.T) {
	SyncGroup()
}
func TestSyncOnce(t *testing.T) {
	SyncOnce()
}

var counter = 0
var mutex sync.Mutex
var wg sync.WaitGroup

func incrCount() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		counter++
	}
}

func incrCountSync() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}
func SyncGroup() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println("add", i)
			wg.Add(1)
			wg.Done()
		}()
	}
	for i := 0; i < 100; i++ {
		go func() {
			for {
				wg.Wait()
			}
		}()
	}
	fmt.Println("done!")
	select {}
}
func SyncOnce() {
	var once sync.Once
	var count = 0
	go func() {
		defer func() {
			count++
			recover()
		}()
		once.Do(func() {
			fmt.Println("exec do")
			count = 1 / count
		})
	}()
	time.Sleep(time.Second)
	once.Do(func() {
		fmt.Println("exec here")
		count = 1 / count
		fmt.Println("exec here", count)

	})
	fmt.Println("end")
}
