package lint_test

import (
	"testing"

	wire "github.com/la-plas-growth/go-docker-ai/deps"
	"github.com/la-plas-growth/go-docker-ai/lib"
)

func TestCreateDockerfile(t *testing.T) {
	f := "Dockerfile"
	dockerFileService := wire.()
	r, err := dockerFileService.AnalyzeDockerFile(f)
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}
	t.Log(lib.PrettyPrint(r))
}
