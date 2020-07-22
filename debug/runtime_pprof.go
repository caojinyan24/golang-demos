package debug

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func Main() {
	fmt.Println("start pprof...")
	startCPUProfile()
	startMemProfile()
	defer stopCPUProfile()
	for {
		time.Sleep(1 * time.Second)
	}

}

var cpuProfile = "cpu.prof"
var memProfile = "mem.prof"

func startCPUProfile() {
	if cpuProfile != "" {
		f, err := os.Create(cpuProfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can not create cpu profile output file: %s",
				err)
			return
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "Can not start cpu profile: %s", err)
			f.Close()
			return
		}
	}
}
func stopCPUProfile() {
	if cpuProfile != "" {
		pprof.StopCPUProfile() // 把记录的概要信息写到已指定的文件
	}
}

func startMemProfile() {
	if memProfile != "" {
		f, err := os.Create(memProfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can not create mem profile output file: %s", err)
			return
		}
		if err = pprof.WriteHeapProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "Can not write %s: %s", memProfile, err)
		}
		f.Close()
	}
}
