package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "jbt",
		Short: "jbt is a toolset",
		Long:  "jbt is a toolset",
	}
	jetbrainsDataDir string
)

func command(name string) *cobra.Command {
	return &cobra.Command{
		Use:   name,
		Short: "Activate " + name + " to 2099",
		Long:  "Activate " + name + " to 2099",
		Run: func(cmd *cobra.Command, args []string) {
			copyJetBrainsFiles()
			copyAppKey(name)
			copyAppVmOptions(name)
			fmt.Println(getAppName(name) + " activated to 2099")
		},
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&jetbrainsDataDir, "dir", "d", "", "The JetBrains data dir, default user home dir")

	rootCmd.AddCommand(command("clion"))
	rootCmd.AddCommand(command("datagrip"))
	rootCmd.AddCommand(command("dataspell"))
	rootCmd.AddCommand(command("goland"))
	rootCmd.AddCommand(command("idea"))
	rootCmd.AddCommand(command("phpstorm"))
	rootCmd.AddCommand(command("pycharm"))
	rootCmd.AddCommand(command("rider"))
	rootCmd.AddCommand(command("webstorm"))
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
