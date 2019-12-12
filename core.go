package mws

import (
	"github.com/syncfuture/go/config"
)

var ConfigProvider config.IConfigProvider

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
}
