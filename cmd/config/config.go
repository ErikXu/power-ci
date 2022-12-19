package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	ConfigCmd.AddCommand(configSetCmd)
	ConfigCmd.AddCommand(configGetCmd)
}

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Config tools",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Get or set config")
	},
}
