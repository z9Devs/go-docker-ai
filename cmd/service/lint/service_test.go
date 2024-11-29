package lint_test

import (
	"testing"

	"github.com/la-plas-growth/go-docker-ai/lib"
)

func TestLintDockerfile(t *testing.T) {
	f := "Dockerfile"
	t.Log(lib.PrettyPrint(f))
}
