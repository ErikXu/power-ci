package docker

import (
	"bufio"
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

var script_centos = `#!/bin/bash
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

var script_debian = `#!/bin/bash
apt-get remove -y docker docker-engine docker.io containerd runc

apt-get update -y

apt-get install -y \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --yes --dearmor -o /etc/apt/keyrings/docker.gpg

echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/debian \
  $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null

apt-get update -y

apt-get install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin

service docker start

service docker enable

docker info`

var script_fedora = `#!/bin/bash
dnf remove -y docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-selinux \
                  docker-engine-selinux \
                  docker-engine

dnf install -y dnf-plugins-core

dnf config-manager \
    --add-repo \
    https://download.docker.com/linux/fedora/docker-ce.repo

dnf install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin

systemctl start docker

systemctl enable docker

docker info`

var script_opensuse_leap = `#!/bin/bash
zypper refresh

zypper update -y

zypper install -y docker

systemctl start docker

systemctl enable docker

docker info`

var dockerInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install docker",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open("/etc/os-release")
		if err != nil {
			fmt.Println("Cannot get OS version")
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		osVersion := "Unknown"

		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "ID=") {
				temp := strings.Replace(line, "ID=", "", -1)
				osVersion = strings.Replace(temp, "\"", "", -1)
				break
			}
		}

		homeDir, _ := os.UserHomeDir()
		os.MkdirAll(path.Join(homeDir, consts.Workspace), os.ModePerm)

		filepath := path.Join(homeDir, consts.Workspace, "install-docker.sh")
		f, _ := os.Create(filepath)

		switch osVersion {
		case "centos", "rocky":
			f.WriteString(script_centos)
		case "debian":
			f.WriteString(script_debian)
		case "fedora":
			f.WriteString(script_fedora)
		case "opensuse-leap":
			f.WriteString(script_opensuse_leap)
		default:
			fmt.Printf("Unsupported OS version: %s\n", osVersion)
			return
		}

		fmt.Printf("OS version: %s\n", osVersion)

		command := exec.Command("bash", filepath)
		f, err = pty.Start(command)
		if err != nil {
			fmt.Println("Install failed")
			return
		}
		io.Copy(os.Stdout, f)

		fmt.Println("Install success, more info please refer https://docs.docker.com/engine/install/centos/")
	},
}
