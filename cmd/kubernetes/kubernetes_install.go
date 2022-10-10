package kubernetes

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
	kubernetesInstallCmd.Flags().StringVarP(&Masters, "masters", "m", "", "Master IP or IP Range")
	kubernetesInstallCmd.MarkFlagRequired("masters")
	kubernetesInstallCmd.Flags().StringVarP(&Nodes, "nodes", "n", "", "Node IP or IP Range")
	kubernetesInstallCmd.MarkFlagRequired("nodes")
	kubernetesInstallCmd.Flags().StringVarP(&Password, "password", "p", "", "SSH Password")
	kubernetesInstallCmd.MarkFlagRequired("password")
}

var Masters string
var Nodes string
var Password string

var template = `#!/bin/bash
wget https://github.com/labring/sealos/releases/download/v4.1.3/sealos_4.1.3_linux_amd64.tar.gz
tar -zxvf sealos_4.1.3_linux_amd64.tar.gz sealos
chmod +x sealos
mv sealos /usr/bin

sealos run labring/kubernetes:v1.24.0 labring/calico:v3.22.1 \
    --masters {MASTERS} \
    --nodes {NODES} \
	-p {PASSWORD}`

var kubernetesInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install kubernetes",
	Run: func(cmd *cobra.Command, args []string) {
		script := strings.Replace(template, "{MASTERS}", Masters, -1)
		script = strings.Replace(script, "{NODES}", Nodes, -1)
		script = strings.Replace(script, "{PASSWORD}", Password, -1)

		homeDir, _ := os.UserHomeDir()
		os.MkdirAll(path.Join(homeDir, consts.Workspace), os.ModePerm)

		filepath := path.Join(homeDir, consts.Workspace, "install-kubernetes.sh")
		f, _ := os.Create(filepath)

		f.WriteString(script)

		command := exec.Command("bash", filepath)
		f, err := pty.Start(command)
		if err != nil {
			fmt.Println("Install failed")
			return
		}
		io.Copy(os.Stdout, f)

		fmt.Println("Install success, more info please refer https://www.sealos.io/docs/getting-started/installation")
	},
}
