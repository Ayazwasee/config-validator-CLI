package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewValidateCommand() *cobra.Command {
	var (
		configPath string
		jsonOutput bool
	)

	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate a .env or .yaml config file",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("✅ Running config validation...")
			fmt.Printf("📁 Config file: %s\n", configPath)
			fmt.Printf("📤 JSON output: %v\n", jsonOutput)
		},
	}

	cmd.Flags().StringVar(&configPath, "config", ".env", "Path to config file")
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "Enable JSON output")

	return cmd
}
