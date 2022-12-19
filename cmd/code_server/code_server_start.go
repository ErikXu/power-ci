package code_server

import (
	"fmt"
	"power-ci/utils"

	"github.com/spf13/cobra"
)

var startScript = `#!/bin/bash
PASSWORD=$(uuidgen)

mkdir -p ~/.config/code-server
cat>~/.config/code-server/config.yaml<<EOF
bind-addr: 0.0.0.0:80
auth: password
password: ${PASSWORD}
cert: false
EOF

echo "Please use ${PASSWORD} to login"
code-server`

var codeServerStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start code server",
	Run: func(cmd *cobra.Command, args []string) {
		filepath := utils.WriteScript("start-code-server.sh", startScript)
		utils.ExecuteScript(filepath)
		fmt.Println("Start success")
	},
}
