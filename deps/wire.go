//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/la-plas-growth/go-docker-ai/cmd/service/dockerfile"
	"github.com/la-plas-growth/go-docker-ai/cmd/service/lint"
	"github.com/la-plas-growth/go-docker-ai/env"
	zap_log "github.com/la-plas-growth/go-docker-ai/lib/log"
)

var CommonSet = wire.NewSet(
	env.NewConfiguration,
	zap_log.NewLogger,
)

// add all commands
func InitLint() lint.IService {
	wire.Build(
		CommonSet,
		lint.NewService,
	)
	return nil
}

// add all commands
func InitDockerfileService() dockerfile.IService {
	wire.Build(
		CommonSet,
		dockerfile.NewService,
	)
	return nil
}
