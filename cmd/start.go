package cmd

import (
	"github.com/Kubectl-UI/kubectl-ui/config"
	"github.com/Kubectl-UI/kubectl-ui/server"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Server (Default port: 3553)",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := getConfig()
		port, _ := cmd.Flags().GetString("port")
		cfg.ReplaceDefault(config.ApplicationPort, port)
		startServer(cfg)

	},
}

func init() {
	// get the value of the port flag
	startCmd.Flags().StringP("port", "p", "3553", "Port to listen on")

	rootCmd.AddCommand(startCmd)
}

func getConfig() *config.KubectlUI {
	cfg := config.Load()
	return cfg
}

func startServer(cfg *config.KubectlUI) {
	server.Start(cfg)
}
