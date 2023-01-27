package code_server

import (
	"fmt"
	"power-ci/utils"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	codeServerStartCmd.Flags().StringVarP(&PORT, "port", "p", "80", "Bind port")
}

var PORT string

var template = `#!/bin/bash
PASSWORD=$(uuidgen)

mkdir -p ~/.config/code-server
cat>~/.config/code-server/config.yaml<<EOF
bind-addr: 0.0.0.0:{PORT}
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
		script := strings.Replace(template, "{PORT}", PORT, -1)
		filepath := utils.WriteScript("start-code-server.sh", script)
		utils.ExecuteScript(filepath)
		fmt.Println("Start success")
	},
}
