package main

import (
	"io"
	"io/ioutil"
	"os/exec"

	"github.com/kr/pty"
)

func main() {
	cmd := exec.Command("vim")
	thepty, err := pty.Start(cmd)
	if err != nil {
		panic(err)
	}
	go func() {
		io.Copy(ioutil.Discard, thepty)
	}()
	cmd.Process.Kill()
}
