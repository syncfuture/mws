package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syncfuture/go/config"
	mwsconfig "github.com/syncfuture/mws/config"
)

func TestDefaultConfigGetter(t *testing.T) {
	configProvider := config.NewJsonConfigProvider()
	mwsConfigGetter := mwsconfig.NewDefaultConfigGetter(configProvider)

	a, err := mwsConfigGetter.Get("ams_report", "LF")
	assert.NoError(t, err)
	t.Log(a)
	a, err = mwsConfigGetter.Get("ams_report", "FF")
	assert.NoError(t, err)
	t.Log(a)
}
