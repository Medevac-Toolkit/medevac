package cmd

import (
	"fmt"
	"os"

	"github.com/Medevac-Toolkit/medevac/internal/k8s"
	"github.com/spf13/cobra"

	"github.com/Medevac-Toolkit/medevac/internal/agent"
)

var configPath string

var rootCmd = &cobra.Command{
	Use:   "medevac",
	Short: "Medevac is a tool for diagnosing pod issues in a Kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running CLI to list pods...")
		clientset, err := k8s.GetKubernetesClient()
		if err != nil {
			fmt.Printf("Error getting Kubernetes client: %v\n", err)
			return
		}
		err = k8s.ListPods(clientset)
		if err != nil {
			fmt.Printf("Error listing pods: %v\n", err)
		}
	},
}

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Run the medevac agent (not yet implemented)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Agent functionality is not yet implemented.")
		// Placeholder for future implementation of the agent
		agent.RunAgent(configPath)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "config.yaml", "Path to the configuration file")
	rootCmd.AddCommand(agentCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
