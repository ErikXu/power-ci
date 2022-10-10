package istio

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	IstioCmd.AddCommand(istioInstallCmd)
}

var IstioCmd = &cobra.Command{
	Use:   "istio",
	Short: "Istio tools",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Istio tools")
	},
}
