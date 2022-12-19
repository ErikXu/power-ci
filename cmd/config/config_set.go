package config

import (
	"fmt"
	"power-ci/utils"

	"github.com/spf13/cobra"
)

func init() {
	configSetCmd.Flags().StringVarP(&SetKey, "key", "k", "", "key")
	configSetCmd.MarkFlagRequired("key")
	configSetCmd.Flags().StringVarP(&SetValue, "value", "v", "", "value")
	configSetCmd.MarkFlagRequired("value")
}

var SetKey string
var SetValue string

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "set config",
	Run: func(cmd *cobra.Command, args []string) {
		configs := utils.GetConfigs()
		configs[SetKey] = SetValue
		path := utils.SaveConfigs(configs)

		fmt.Printf("Config is set and saved to [%s]\n", path)
	},
}
