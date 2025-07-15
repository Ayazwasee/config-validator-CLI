package main

import (
	"fmt"
	"os"

	"config-validator-cli/cmd"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "config-validator",
		Short: "Validate .env or .yaml config files",
	}

	rootCmd.AddCommand(cmd.NewValidateCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
