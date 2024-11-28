package dockerfile

import (
	wire "github.com/la-plas-growth/go-docker-ai/deps"
	"github.com/la-plas-growth/go-docker-ai/lib"
	"github.com/spf13/cobra"
)

func CreateDockerfileCommand() *cobra.Command {
	var lang, path string
	cmd := cobra.Command{
		Use:   "dockerfile",
		Short: "Create Dockerfile by language/type",
		Run: func(cmd *cobra.Command, args []string) {
			dockerFileService := wire.InitDockerfileService()
			r, err := dockerFileService.CreateDockerFile(lang)
			if err != nil {
				cmd.PrintErr("Error: ", err)
				return
			}
			cmd.Println(lib.PrettyPrint(r))
		},
	}

	cmd.Flags().StringVarP(&lang, "type", "t", "golang", "Create dockerfile language/type")
	cmd.Flags().StringVarP(&lang, "path", "p", "./", "Path for the dockerfile")

	_ = cmd.MarkFlagRequired("type")

	return &cmd
}
