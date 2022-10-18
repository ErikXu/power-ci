package dotnet

import (
	"io"
	"os"
	"os/exec"
	"path"
	"power-ci/consts"

	"github.com/creack/pty"
	"github.com/spf13/cobra"
)

var script = `#!/bin/bash
rpm -Uvh https://packages.microsoft.com/config/centos/7/packages-microsoft-prod.rpm

yum install dotnet-sdk-6.0 -y

echo "Install success, dotnet version:"

dotnet --version`

var dotnetInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install .Net",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, _ := os.UserHomeDir()
		os.MkdirAll(path.Join(homeDir, consts.Workspace), os.ModePerm)

		filepath := path.Join(homeDir, consts.Workspace, "install-dotnet.sh")
		f, _ := os.Create(filepath)

		f.WriteString(script)

		command := exec.Command("bash", filepath)
		f, err := pty.Start(command)
		if err != nil {
			panic(err)
		}
		io.Copy(os.Stdout, f)
	},
}
