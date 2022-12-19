package utils

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"power-ci/consts"

	"github.com/creack/pty"
)

func WriteScript(filename string, script string) string {
	homeDir, _ := os.UserHomeDir()
	os.MkdirAll(path.Join(homeDir, consts.Workspace), os.ModePerm)
	filepath := path.Join(homeDir, consts.Workspace, filename)
	f, _ := os.Create(filepath)

	f.WriteString(script)
	return filepath
}

func ExecuteScript(filepath string) {
	fmt.Printf("[%s] is executing...\n", filepath)
	command := exec.Command("bash", filepath)
	f, err := pty.Start(command)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, f)
}
