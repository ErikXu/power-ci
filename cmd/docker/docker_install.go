package docker

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
yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-engine

yum install yum-utils -y

yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo

yum install docker-ce docker-ce-cli containerd.io docker-compose-plugin -y

systemctl start docker

systemctl enable docker

docker info`

var dockerInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install docker",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, _ := os.UserHomeDir()
		os.MkdirAll(path.Join(homeDir, consts.Workspace), os.ModePerm)

		filepath := path.Join(homeDir, consts.Workspace, "install-docker.sh")
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
