package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd is the default command
var RootCmd = &cobra.Command{
	Use:   "go-dart",
	Short: "go-dart is cool",
	Long: `A better dart game than the chinese one. This binary is composed of a server and a cli client.
Complete documentation is available at http://github.com/Zenika/go-dart.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
	RootCmd.AddCommand(sampleCmd)
	RootCmd.AddCommand(versionCmd)
	RootCmd.PersistentFlags().StringP("server", "s", "http://localhost:8080/", "Server address")
	viper.BindPFlag("server", RootCmd.PersistentFlags().Lookup("server"))
}
