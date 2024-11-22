package cmd

import (
	"github.com/la-plas-growth/GO-DockerLint-AI/cmd/command/lint"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "tool",
		Short: "Tool is a CLI application for lint.",
	}
	rootCmd.AddCommand(lint.NewBaseCommand())
	return rootCmd
}

func Execute() error {
	return NewRootCommand().Execute()
}
