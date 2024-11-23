//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/la-plas-growth/GO-DockerLint-AI/cmd/service/lint"
	"github.com/la-plas-growth/GO-DockerLint-AI/env"
	zap_log "github.com/la-plas-growth/GO-DockerLint-AI/lib/log"
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
