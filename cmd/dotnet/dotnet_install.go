package dotnet

import (
	"fmt"
	"power-ci/utils"

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
		filepath := utils.WriteScript("install-dotnet.sh", script)
		utils.ExecuteScript(filepath)
		fmt.Println("Install success")
	},
}
