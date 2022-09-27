package main

import (
	"io"
	"os"
	"os/exec"
	"github.com/creack/pty"
)

func main() {
	cmd := exec.Command("yum", "update", "-y")
	f, err := pty.Start(cmd)
    if err != nil {
        panic(err)
    }
    io.Copy(os.Stdout, f)
}

