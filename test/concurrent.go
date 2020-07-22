package test

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func main() {
	Cmd()
	PipeCmd()
}
func Cmd() {
	var output1 bytes.Buffer
	cmd1 := exec.Command("echo", "-n", "abc")
	stdout0, err := cmd1.StdoutPipe()
	if err != nil {
		fmt.Println("set out pipe line error,err=", err)
	}
	if err := cmd1.Start(); err != nil {
		fmt.Println("cmd start error,err=", err)
		return
	}
	for {
		tmpOut := make([]byte, 5)
		n, err := stdout0.Read(tmpOut)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read stdout error,err=", err)
			return
		}
		if n > 0 {
			output1.Write(tmpOut[0:n])
		}
	}
	fmt.Println("get stdout=", output1.String())
}
func PipeCmd() {
	psCmd := exec.Command("echo", "-n", "abc\ndef")
	var psOutput bytes.Buffer
	psCmd.Stdout = &psOutput
	grepCmd := exec.Command("grep", "a")
	var grepOutput bytes.Buffer
	grepCmd.Stdout = &grepOutput
	grepCmd.Stdin = &psOutput

	if err := psCmd.Start(); err != nil {
		fmt.Println("cmd start error,err=", err)
		return
	}
	if err := psCmd.Wait(); err != nil {
		fmt.Println("ps command wait error,err=", err)
		return
	}
	if err := grepCmd.Start(); err != nil {
		fmt.Println("grep start error,err=", err)
		return
	}
	if err := grepCmd.Wait(); err != nil {
		fmt.Println("grep command wait error,err=", err)
		return
	}
	fmt.Println("grep result", grepOutput.String())
}
