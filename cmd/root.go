package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "power-ci",
	Short: "power-ci is a helpful tool to deal with devops",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("power-ci is a helpful tool to deal with devops")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
