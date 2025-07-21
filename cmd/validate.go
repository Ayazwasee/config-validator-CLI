package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

			var (
				env map[string]string
				err error
			)

			ext := strings.ToLower(filepath.Ext(configFile))

			if ext == ".yaml" || ext == ".yml" {
				env, err = core.ParseYAMLFile(configFile)
			} else {
				env, err = core.ParseEnvFile(configFile)
			}

			if err != nil {
				fmt.Println("❌ Error reading config file:", err)
				return
			}

			missing := core.CheckRequiredKeys(env, core.RequiredKeys)
			typeErrors := core.ValidateTypes(env)

			if jsonOutput {
				output := map[string]interface{}{
					"missing_keys": missing,
					"type_errors":  typeErrors,
					"status":       len(missing) == 0 && len(typeErrors) == 0,
				}
				json.NewEncoder(os.Stdout).Encode(output)
			} else {
				if len(missing) > 0 {
					fmt.Println("❌ Missing required keys:")
					for _, key := range missing {
						fmt.Printf("  - %s\n", key)
					}
				}

				if len(typeErrors) > 0 {
					fmt.Println("⚠️  Type validation errors:")
					for _, err := range typeErrors {
						fmt.Printf("  - %s: %s\n", err.Key, err.Error)
					}
				}

				if len(missing) == 0 && len(typeErrors) == 0 {
					fmt.Println("✅ All required keys and types are valid.")
				}
			}
		},
	}

	cmd.Flags().StringVar(&configFile, "config", ".env", "Path to config file (.env or .yaml)")
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "Enable JSON output")

	return cmd
}
