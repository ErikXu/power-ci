package docker

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	DockerCmd.AddCommand(dockerInstallCmd)
}

var DockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Docker tools",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Docker tools")
	},
}
