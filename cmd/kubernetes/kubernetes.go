package kubernetes

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	KubernetesCmd.AddCommand(kubernetesInstallCmd)
}

var KubernetesCmd = &cobra.Command{
	Use:   "k8s",
	Short: "Kubernetes tools",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Kubernetes tools")
	},
}
