package jenkins

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	JenkinsCmd.AddCommand(JenkinsInstallCmd)
	JenkinsCmd.AddCommand(jenkinsPasswordCmd)
}

var JenkinsCmd = &cobra.Command{
	Use:   "jenkins",
	Short: "Jenkins tools",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Jenkins tools")
	},
}
