package tests

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/syncfuture/mws"

	log "github.com/kataras/golog"
	"github.com/syncfuture/go/config"
	"github.com/syncfuture/go/u"
	mwsconfig "github.com/syncfuture/mws/config"
)

var (
	_apiSet *mws.APISet
)

func init() {
	configProvider := config.NewJsonConfigProvider()
	configHttpClient(configProvider)
	configGetter := mwsconfig.NewDefaultConfigGetter(configProvider)
	c, err := configGetter.Get("ams_report", "LF")
	u.LogFaltal(err)

	_apiSet = mws.NewAPISet(c)
	log.Info("Starting test.")
}

func configHttpClient(configProvider config.IConfigProvider) {
	// Http客户端配置
	skipCertVerification := configProvider.GetBool("Http.SkipCertVerification")
	proxy := configProvider.GetString("Http.Proxy")
	if skipCertVerification || proxy != "" {
		// 任意条件满足，则使用自定义传输层
		transport := new(http.Transport)
		if skipCertVerification {
			// 跳过证书验证
			transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: skipCertVerification}
		}
		if proxy != "" {
			// 使用代理
			transport.Proxy = func(r *http.Request) (*url.URL, error) {
				return url.Parse(proxy)
			}
		}
		http.DefaultClient.Transport = transport
	}
}
