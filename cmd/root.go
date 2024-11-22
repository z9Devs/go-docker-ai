package cmd

import (

	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "tool",
		Short: "Tool is a CLI application for lint.",
	}
	//rootCmd.AddCommand()
	return rootCmd
}

func Execute() error {
	return NewRootCommand().Execute()
}