package jenkins

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var jenkinsPasswordCmd = &cobra.Command{
	Use:   "password",
	Short: "Get initial root password of jenkins",
	Run: func(cmd *cobra.Command, args []string) {
		password, err := os.ReadFile("/var/lib/jenkins/secrets/initialAdminPassword")
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(string(password))
	},
}
