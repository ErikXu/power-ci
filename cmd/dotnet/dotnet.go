package dotnet

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	DotnetCmd.AddCommand(dotnetInstallCmd)
}

var DotnetCmd = &cobra.Command{
	Use:   "dotnet",
	Short: ".Net tools",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(".Net tools")
	},
}
