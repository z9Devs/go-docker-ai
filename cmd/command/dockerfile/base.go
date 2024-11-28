package dockerfile

import "github.com/spf13/cobra"

func NewBaseCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dockerfile",
		Short: "Dockerfile related commands",
	}

	cmd.AddCommand(CreateDockerfileCommand())
	cmd.AddCommand(LintCommand())

	return cmd
}
