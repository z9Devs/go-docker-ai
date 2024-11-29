package project

import "github.com/spf13/cobra"

func NewBaseCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project",
		Short: "Project related commands",
	}

	return cmd
}
