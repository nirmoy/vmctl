package cmd

import (
	"fmt"
	"os"

	"github.com/nirmoy/vmctl/api/app"
	"github.com/spf13/cobra"
)

var port = "3000"

var rootCmd = &cobra.Command{
	Use:   "vmctl",
	Short: "vmctl is a tool to control VM resource in private/public Cloud",
	Run: func(cmd *cobra.Command, args []string) {
		app := &app.App{}
		app.Initialize()
		fmt.Printf("Listening on %s\n", port)
		app.Run(":" + port)
	},
}

func init() {
	rootCmd.Flags().StringVar(&port, "port", "3000", "Port on which vmctl will accept request")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
