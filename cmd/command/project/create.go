package project

import (
	"github.com/la-plas-growth/go-docker-ai/lib"
	"github.com/spf13/cobra"
)

func CreateProjectCommand() *cobra.Command {
    var lang string 
    cmd := cobra.Command{
		Use:   "create",
		Short: "Create project by language",
		Run: func(cmd *cobra.Command, args []string) {
			//
			cmd.Println(lib.PrettyPrint(args))
		},
	}

    cmd.Flags().StringVarP(&lang, "type", "t", "go", "Create projecet by language/type")

	_ = cmd.MarkFlagRequired("type")

	return &cmd
}