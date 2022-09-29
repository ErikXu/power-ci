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

var installScript = `#!/bin/bash
curl -fsSL https://code-server.dev/install.sh | sh`

var codeServerInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install code server",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, _ := os.UserHomeDir()
		os.MkdirAll(path.Join(homeDir, consts.Workspace), os.ModePerm)

		filepath := path.Join(homeDir, consts.Workspace, "install-code-server.sh")
		f, _ := os.Create(filepath)

		f.WriteString(installScript)

		command := exec.Command("bash", filepath)
		f, err := pty.Start(command)
		if err != nil {
			panic(err)
		}
		io.Copy(os.Stdout, f)
	},
}
