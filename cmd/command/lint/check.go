package lint

import (
	wire "github.com/la-plas-growth/GO-DockerLint-AI/deps"
	"github.com/la-plas-growth/GO-DockerLint-AI/lib"
	"github.com/spf13/cobra"
)

func NewCheckDockerfile() *cobra.Command {
	var dockerFile string
	cmd := &cobra.Command{
		Use:   "lint",
		Short: "Check dockerfile for lint",
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

	cmd.Flags().StringVarP(&dockerFile, "lint", "l", "", "lint dockerfile")

	_ = cmd.MarkFlagRequired("lint")

	return cmd

}
