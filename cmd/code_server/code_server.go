package code_server

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	CodeServerCmd.AddCommand(codeServerInstallCmd)
	CodeServerCmd.AddCommand(codeServerStartCmd)
}

var CodeServerCmd = &cobra.Command{
	Use:   "code-server",
	Short: "Code server tools",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Code server tools")
	},
}
