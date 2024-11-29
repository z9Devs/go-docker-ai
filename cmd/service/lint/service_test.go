package lint_test

import (
	"testing"

	wire "github.com/la-plas-growth/go-docker-ai/deps"
	"github.com/la-plas-growth/go-docker-ai/lib"
)

func TestLintDockerfile *testing.T) {
	f := "Dockerfile"
	t.Log(lib.PrettyPrint(f))
	return 
	dockerFileService := wire.InitLint()
	r, err := dockerFileService.AnalyzeDockerFile(f)
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}
	t.Log(lib.PrettyPrint(r))
}
