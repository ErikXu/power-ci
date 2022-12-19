package docker

import (
	"bufio"
	"fmt"
	"os"
	"power-ci/utils"
	"strings"

	"github.com/spf13/cobra"
)

var scriptCentos = `#!/bin/bash
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

var scriptDebian = `#!/bin/bash
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

var scriptFedora = `#!/bin/bash
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

var scriptOpensuseLeap = `#!/bin/bash
zypper refresh

zypper update -y

zypper install -y docker

systemctl start docker

systemctl enable docker

docker info`

var scriptUbuntu = `#!/bin/bash
apt-get update -y

apt-get install -y \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --yes --dearmor -o /etc/apt/keyrings/docker.gpg

echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null

apt-get update -y

apt-get install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin

service docker start

service docker enable

docker info`

func getOsVersion() string {
	osVersion := "Unknown"
	file, err := os.Open("/etc/os-release")
	if err != nil {
		fmt.Println("Cannot get OS version")
		return osVersion
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "ID=") {
			temp := strings.Replace(line, "ID=", "", -1)
			osVersion = strings.Replace(temp, "\"", "", -1)
			break
		}
	}

	return osVersion
}

var dockerInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install docker",
	Run: func(cmd *cobra.Command, args []string) {
		osVersion := getOsVersion()

		filepath := ""
		filename := "install-docker.sh"
		switch osVersion {
		case "centos", "rocky":
			filepath = utils.WriteScript(filename, scriptCentos)
		case "debian":
			filepath = utils.WriteScript(filename, scriptDebian)
		case "fedora":
			filepath = utils.WriteScript(filename, scriptFedora)
		case "opensuse-leap":
			filepath = utils.WriteScript(filename, scriptOpensuseLeap)
		case "ubuntu":
			filepath = utils.WriteScript(filename, scriptUbuntu)
		default:
			fmt.Printf("Unsupported OS version: %s\n", osVersion)
			return
		}

		fmt.Printf("OS version: %s\n", osVersion)

		utils.ExecuteScript(filepath)

		fmt.Println("Install success. More info please refer https://docs.docker.com/engine/install/")
	},
}
