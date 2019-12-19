package mws

import (
	log "github.com/kataras/golog"
	"github.com/syncfuture/go/config"
)

var (
	_apiSet *APISet
)

func init() {
	configProvider := config.NewJsonConfigProvider()
	_apiSet = NewAPISet("LF", configProvider)
	log.Info("Starting test.")
}
