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

func init() {
	gitlabInstallCmd.Flags().StringVarP(&Url, "url", "u", "", "Gitlab Url")
	gitlabInstallCmd.MarkFlagRequired("url")
}

var Url string

var template = `#!/bin/bash
yum install curl policycoreutils-python openssh-server perl -y

yum install -y postfix
systemctl enable postfix
systemctl start postfix

curl https://packages.gitlab.com/install/repositories/gitlab/gitlab-ce/script.rpm.sh | bash

EXTERNAL_URL="{EXTERNAL_URL}" yum install gitlab-ce -y`

var gitlabInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install gitlab",
	Run: func(cmd *cobra.Command, args []string) {
		script := strings.Replace(template, "{EXTERNAL_URL}", Url, -1)

		homeDir, _ := os.UserHomeDir()
		os.MkdirAll(path.Join(homeDir, consts.Workspace), os.ModePerm)

		filepath := path.Join(homeDir, consts.Workspace, "install-gitlab.sh")
		f, _ := os.Create(filepath)

		f.WriteString(script)

		command := exec.Command("bash", filepath)
		f, err := pty.Start(command)
		if err != nil {
			fmt.Print("Install failed")
			return
		}
		io.Copy(os.Stdout, f)

		fmt.Print("Install success, more info please refer https://about.gitlab.com/install/#centos-7")
	},
}
