package config

import (
	"fmt"
	"power-ci/utils"

	"github.com/spf13/cobra"
)

func init() {
	configGetCmd.Flags().StringVarP(&GetKey, "key", "k", "", "key")
	configGetCmd.MarkFlagRequired("key")
}

var GetKey string

var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get config",
	Run: func(cmd *cobra.Command, args []string) {
		configs := utils.GetConfigs()
		if configs == nil {
			configs = make(map[string]string)
		}

		value, ok := configs[GetKey]
		if ok {
			fmt.Printf("The value of key [%s] is %s\n", GetKey, value)
		} else {
			fmt.Printf("Can not find value of key [%s]\n", GetKey)
		}
	},
}
