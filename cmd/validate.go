package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"config-validator-cli/core"

	"github.com/spf13/cobra"
)

var configFile string
var jsonOutput bool

func NewValidateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate a .env or .yaml config file",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("✅ Running config validation...")
			fmt.Printf("📁 Config file: %s\n", configFile)
			fmt.Printf("📤 JSON output: %v\n", jsonOutput)

			// Parse the .env file
			env, err := core.ParseEnvFile(configFile)
			if err != nil {
				fmt.Println("❌ Error reading config file:", err)
				return
			}

			// Check required keys
			missing := core.CheckRequiredKeys(env, core.RequiredKeys)

			if jsonOutput {
				// Output result in JSON
				json.NewEncoder(os.Stdout).Encode(map[string]interface{}{
					"missing_keys": missing,
					"status":       len(missing) == 0,
				})
			} else {
				// Output in plain text
				if len(missing) == 0 {
					fmt.Println("✅ All required keys are present.")
				} else {
					fmt.Println("❌ Missing required keys:")
					for _, key := range missing {
						fmt.Println("  -", key)
					}
				}
			}
		},
	}

	cmd.Flags().StringVar(&configFile, "config", ".env", "Path to config file")
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "Enable JSON output")

	return cmd
}
