package gitlab

import (
	"fmt"
	"power-ci/utils"
	"strings"

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
		filepath := utils.WriteScript("install-gitlab.sh", script)
		utils.ExecuteScript(filepath)
		fmt.Println("Install success. More info please refer https://about.gitlab.com/install/#centos-7")
	},
}
