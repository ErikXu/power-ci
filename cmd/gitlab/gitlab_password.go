package gitlab

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var gitlabPasswordCmd = &cobra.Command{
	Use:   "password",
	Short: "Get initial root password of gitlab",
	Run: func(cmd *cobra.Command, args []string) {
		password, err := os.ReadFile("/etc/gitlab/initial_root_password")
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(string(password))
	},
}
