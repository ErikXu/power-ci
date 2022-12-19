package cmd

import (
	"fmt"
	"os"
	"path"
	"power-ci/cmd/code_server"
	"power-ci/cmd/config"
	"power-ci/cmd/docker"
	"power-ci/cmd/dotnet"
	"power-ci/cmd/gitlab"
	"power-ci/cmd/golang"
	"power-ci/cmd/istio"
	"power-ci/cmd/jenkins"
	"power-ci/cmd/kubernetes"
	"power-ci/consts"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(docker.DockerCmd)
	rootCmd.AddCommand(gitlab.GitlabCmd)
	rootCmd.AddCommand(golang.GolangCmd)
	rootCmd.AddCommand(code_server.CodeServerCmd)
	rootCmd.AddCommand(jenkins.JenkinsCmd)
	rootCmd.AddCommand(kubernetes.KubernetesCmd)
	rootCmd.AddCommand(istio.IstioCmd)
	rootCmd.AddCommand(dotnet.DotnetCmd)
	rootCmd.AddCommand(config.ConfigCmd)
}

var rootCmd = &cobra.Command{
	Use:   "power-ci",
	Short: "power-ci is a powerful tools for devops",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("power-ci is a powerful tools for devops")

		homeDir, _ := os.UserHomeDir()
		workspace := path.Join(homeDir, consts.Workspace)
		os.MkdirAll(workspace, os.ModePerm)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
