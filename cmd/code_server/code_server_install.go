package code_server

import (
	"fmt"
	"power-ci/utils"

	"github.com/spf13/cobra"
)

var installScript = `#!/bin/bash
curl -fsSL https://code-server.dev/install.sh | sh`

var codeServerInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install code server",
	Run: func(cmd *cobra.Command, args []string) {
		filepath := utils.WriteScript("install-code-server.sh", installScript)
		utils.ExecuteScript(filepath)
		fmt.Println("Install success")
	},
}
