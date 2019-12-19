package mws

import "github.com/syncfuture/go/config"

var (
	_apiSet *APISet
)

func init() {
	configProvider := config.NewJsonConfigProvider()
	_apiSet = NewAPISet("LF", configProvider)
}
