package jenkins

import (
	"fmt"
	"power-ci/utils"

	"github.com/spf13/cobra"
)

var installScript = `#!/bin/bash
wget -O /etc/yum.repos.d/jenkins.repo \
    https://pkg.jenkins.io/redhat-stable/jenkins.repo --no-check-certificate

rpm --import https://pkg.jenkins.io/redhat-stable/jenkins.io.key

yum upgrade -y

yum install java-11-openjdk -y

yum install jenkins -y

# sed -i 's|\JENKINS_USER="jenkins"|JENKINS_USER="root"|g' /etc/sysconfig/jenkins

# chown -R root:root /var/lib/jenkins
# chown -R root:root /var/cache/jenkins
# chown -R root:root /var/log/jenkins

systemctl daemon-reload

systemctl start jenkins

systemctl enable jenkins`

var JenkinsInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install jenkins",
	Run: func(cmd *cobra.Command, args []string) {
		filepath := utils.WriteScript("install-jenkins.sh", installScript)
		utils.ExecuteScript(filepath)
		fmt.Println("Install success. More info please refer https://www.jenkins.io/doc/book/installing/linux/#red-hat-centos")
	},
}
