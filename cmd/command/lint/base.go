package lint

import "github.com/spf13/cobra"

func NewBaseCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dockerlint",
		Short: "Dockerfile lint related commands",
	}

	cmd.AddCommand(NewCheckDockerfile())

	return cmd
}
