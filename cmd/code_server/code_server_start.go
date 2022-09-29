package code_server

import (
	"io"
	"os"
	"os/exec"
	"path"
	"power-ci/consts"

	"github.com/creack/pty"
	"github.com/spf13/cobra"
)

var startScript = `#!/bin/bash
PASSWORD=$(uuidgen)

mkdir -p ~/.config/code-server
cat>~/.config/code-server/config.yaml<<EOF
bind-addr: 0.0.0.0:80
auth: password
password: ${PASSWORD}
cert: false
EOF

echo "Please use ${PASSWORD} to login"
code-server`

var codeServerStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start code server",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, _ := os.UserHomeDir()
		os.MkdirAll(path.Join(homeDir, consts.Workspace), os.ModePerm)

		filepath := path.Join(homeDir, consts.Workspace, "start-code-server.sh")
		f, _ := os.Create(filepath)

		f.WriteString(startScript)

		command := exec.Command("bash", filepath)
		f, err := pty.Start(command)
		if err != nil {
			panic(err)
		}
		io.Copy(os.Stdout, f)
	},
}
