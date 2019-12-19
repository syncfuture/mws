package mws

import (
	"net/http"
	"net/url"

	"github.com/syncfuture/go/config"
)

var (
	ConfigProvider config.IConfigProvider
)

const (
	Seller_LF = "LF"
	Seller_FF = "FF"
)

type (
	apiBase struct {
		Seller  string
		Module  string
		Version string
	}
	queryBase struct {
	}
)

func init() {
	// Must put a configs.json file under same folder
	// Exmaple: configs.json
	// {
	// 	"BaseUrl": "https://mws.amazonservices.com/",
	// 	"AccessKey": "xxxx",
	// 	"AccessSecret": "xxxx",
	// 	"MarketplaceId": "xxxx",
	// 	"SellerId": "xxxx",
	// 	"SignatureVersion": "2",
	// 	"SignatureMethod": "HmacSHA256"
	// }
	ConfigProvider = config.NewJsonConfigProvider()

	proxy := ConfigProvider.GetString("Proxy")
	if proxy != "" {
		// 任意条件满足，则使用自定义传输层
		transport := new(http.Transport)
		// 使用代理
		transport.Proxy = func(r *http.Request) (*url.URL, error) {
			return url.Parse(proxy)
		}
		http.DefaultClient.Transport = transport
	}
}
