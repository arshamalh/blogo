package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var globalConfig struct {
	SomeSetting    string
	AnotherSetting int
}

var rootCmd = &cobra.Command{
	Use:   "blogo",
	Short: "Blogo is a simple blog for educational purposes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("SomeSetting: %s\n", globalConfig.SomeSetting)
		fmt.Printf("AnotherSetting: %d\n", globalConfig.AnotherSetting)
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	globalConfig = struct {
		SomeSetting    string
		AnotherSetting int
	}{
		SomeSetting:    "default_value",
		AnotherSetting: 42,
	}
}
