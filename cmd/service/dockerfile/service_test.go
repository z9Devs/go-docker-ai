package dockerfile_test

import (
	"testing"

	wire "github.com/la-plas-growth/go-docker-ai/deps"
	"github.com/la-plas-growth/go-docker-ai/lib"
)

func TestCreateDockerfile(t *testing.T) {
	lang := "golang"
	path := "./"
	dockerFileService := wire.InitDockerfileService()
	r, err := dockerFileService.CreateDockerFile(lang, path)
	if err != nil {
		t.Errorf("Error: ", err)
		return
	}
	t.Log(lib.PrettyPrint(r))
}
