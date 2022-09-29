package golang

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	GolangCmd.AddCommand(golangInstallCmd)
}

var GolangCmd = &cobra.Command{
	Use:   "go",
	Short: "Golang tools",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Golang tools")
	},
}
