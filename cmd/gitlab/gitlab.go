package gitlab

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	GitlabCmd.AddCommand(gitlabInstallCmd)
	GitlabCmd.AddCommand(gitlabStartCmd)
	GitlabCmd.AddCommand(gitlabPasswordCmd)
	GitlabCmd.AddCommand(gitlabInitCmd)
	GitlabCmd.AddCommand(gitlabDockerCmd)
}

var GitlabCmd = &cobra.Command{
	Use:   "gitlab",
	Short: "Gitlab tools",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gitlab tools")
	},
}
