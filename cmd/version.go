package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get version of power-ci",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("0.1.1")
	},
}
