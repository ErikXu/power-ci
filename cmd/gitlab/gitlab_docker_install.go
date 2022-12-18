package gitlab

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

var gitlabDockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Install gitlab using docker",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use --help to get the command detils.")
	},
}

func init() {
	gitlabDockerCmd.AddCommand(gitlabDockerInstallCmd)

	gitlabDockerInstallCmd.Flags().StringVarP(&Hostname, "hostname", "H", "", "Gitlab access host or IP, eg: gitlab.example.com or 127.0.0.1")
	gitlabDockerInstallCmd.MarkFlagRequired("hostname")
}

var Hostname string

var docker_install_template = `#!/bin/bash
mkdir -p /etc/gitlab
mkdir -p /var/log/gitlab
mkdir -p /var/opt/gitlab

docker run --detach \
  --hostname {HOSTNAME} \
  --publish 443:443 --publish 80:80 --publish 10022:22 \
  -e "GITLAB_SHELL_SSH_PORT=10022" \
  --name gitlab \
  --restart always \
  --volume /etc/gitlab:/etc/gitlab \
  --volume /var/log/gitlab:/var/log/gitlab \
  --volume /var/opt/gitlab:/var/opt/gitlab \
  --shm-size 256m \
  gitlab/gitlab-ce:latest
`

var gitlabDockerInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install gitlab using docker",
	Run: func(cmd *cobra.Command, args []string) {
		script := strings.Replace(docker_install_template, "{HOSTNAME}", Hostname, -1)

		homeDir, _ := os.UserHomeDir()
		os.MkdirAll(path.Join(homeDir, consts.Workspace), os.ModePerm)

		filepath := path.Join(homeDir, consts.Workspace, "install-gitlab-docker.sh")
		f, _ := os.Create(filepath)

		f.WriteString(script)

		command := exec.Command("bash", filepath)
		f, err := pty.Start(command)
		if err != nil {
			fmt.Println("Install failed")
			return
		}
		io.Copy(os.Stdout, f)

		fmt.Println("Install success, more info please refer https://docs.gitlab.com/ee/install/docker.html")
	},
}
