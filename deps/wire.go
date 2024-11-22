//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/la-plas-growth/GO-DockerLint-AI/env"
	zap_log "github.com/la-plas-growth/GO-DockerLint-AI/lib/log"
)

var CommonSet = wire.NewSet(
	env.NewConfiguration,
	zap_log.NewLogger,
)

// add all commands 