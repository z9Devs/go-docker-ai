package cmd

import (
	"github.com/la-plas-growth/go-docker-ai/cmd/command/dockerfile"
	"github.com/la-plas-growth/go-docker-ai/cmd/command/project"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "tool",
		Short: "Tool is a CLI application for lint.",
	}
	//
	rootCmd.AddCommand(dockerfile.NewBaseCommand())
	rootCmd.AddCommand(project.NewBaseCommand())
	//
	return rootCmd
}

func Execute() error {
	return NewRootCommand().Execute()
}
