package golang

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"power-ci/consts"
	"strings"

	"github.com/creack/pty"
	"github.com/spf13/cobra"
)

func init() {
	golangInstallCmd.Flags().StringVarP(&Version, "version", "v", "1.19.1", "Golang Version")
}

var Version string

var template = `#!/bin/bash
wget https://dl.google.com/go/go{VERSION}.linux-amd64.tar.gz

tar -xvf go{VERSION}.linux-amd64.tar.gz

cp -r ./go /usr/local/go

cat>>~/.bashrc<<EOF
export PATH=$PATH:/usr/local/go/bin
EOF

source ~/.bashrc`

var golangInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install golang",
	Run: func(cmd *cobra.Command, args []string) {
		script := strings.Replace(template, "{VERSION}", Version, -1)

		homeDir, _ := os.UserHomeDir()
		os.MkdirAll(path.Join(homeDir, consts.Workspace), os.ModePerm)

		filepath := path.Join(homeDir, consts.Workspace, "install-golang.sh")
		f, _ := os.Create(filepath)

		f.WriteString(script)

		command := exec.Command("bash", filepath)
		f, err := pty.Start(command)
		if err != nil {
			panic(err)
		}
		io.Copy(os.Stdout, f)

		fmt.Println("Please run 'source ~/.bashrc' to set the env")
	},
}
