package cmd

import (
	"fmt"
	"os"

	"github.com/atomisadev/portkey/internal/tui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "portkey",
	Short: "PortKey: Visual SSH Connection Manager",
	Long: `PortKey is a TUI-based manager for SSH connections and tunnels.
	It allows you to visualize and manage your port forwardings with ease.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := tui.Start(); err != nil {
			fmt.Printf("Error running PortKey: %v\n", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Init() {

}
