package golang

import (
	"fmt"
	"power-ci/utils"
	"strings"

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
		filepath := utils.WriteScript("install-golang.sh", script)
		utils.ExecuteScript(filepath)
		fmt.Println("Install success. Please run 'source ~/.bashrc' to set the env")
	},
}
