package lint

import (
	wire "github.com/la-plas-growth/go-docker-ai/deps"
	"github.com/la-plas-growth/go-docker-ai/lib"
	"github.com/spf13/cobra"
)

func LintCommand() *cobra.Command {
	var dockerFile string
	cmd := &cobra.Command{
		Use:   "lint",
		Short: "Check Dockerfile for lint",
		Run: func(cmd *cobra.Command, args []string) {
			lintService := wire.InitLint()
			r, err := lintService.AnalyzeDockerFile(dockerFile)
			if err != nil {
				cmd.PrintErr("Error: ", err)
				return
			}
			cmd.Println(lib.PrettyPrint(r))
		},
	}

	cmd.Flags().StringVarP(&dockerFile, "file", "f", "Dockerfile", "Dockerfile to lint")

	_ = cmd.MarkFlagRequired("file")

	return cmd

}
