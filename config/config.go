package config

import (
	"errors"

	config "github.com/syncfuture/go/sconfig"
	"github.com/syncfuture/mws/protoc/mwsconfig"
)

var ConfigNotFoundError = errors.New("Config not found")

// type MWSConfig struct {
// 	ID               string
// 	UserID           string
// 	SellerID         string
// 	MarketplaceID    string
// 	BaseURL          string
// 	SignatureVersion string
// 	SignatureMethod  string
// 	AuthToken        string
// 	AccessKey        string
// 	AccessSecret     string
// }

type IConfigGetter interface {
	Get(userID, configID string) (*mwsconfig.MWSConfig, error)
}

type IConfigManager interface {
	Create(config *mwsconfig.MWSConfig) error
	Update(config *mwsconfig.MWSConfig) error
	Delete(userid, id string) error
}

func NewDefaultConfigGetter(configProvider config.IConfigProvider) IConfigGetter {
	return &defaultConfigGetter{
		ConfigProvider: configProvider,
	}
}

type defaultConfigGetter struct {
	ConfigProvider config.IConfigProvider
}

func (x *defaultConfigGetter) Get(userID, configID string) (r *mwsconfig.MWSConfig, err error) {
	userJsonPath := "MWS." + userID + "." + configID

	r = new(mwsconfig.MWSConfig)
	r.ID = configID
	r.UserID = userID
	r.SellerID = x.ConfigProvider.GetString(userJsonPath + ".SellerID")
	if r.SellerID == "" {
		err = errors.New("'SellerID' cannot be empty")
		return
	}
	r.MarketplaceID = x.ConfigProvider.GetString(userJsonPath + ".MarketplaceID")
	if r.MarketplaceID == "" {
		err = errors.New("'MarketplaceID' cannot be empty")
		return
	}
	r.BaseURL = x.ConfigProvider.GetString(userJsonPath + ".BaseURL")
	if r.BaseURL == "" {
		err = errors.New("'BaseURL' cannot be empty")
		return
	}
	r.SignatureMethod = x.ConfigProvider.GetString(userJsonPath + ".SignatureMethod")
	if r.SignatureMethod == "" {
		err = errors.New("'SignatureMethod' cannot be empty")
		return
	}
	r.SignatureVersion = x.ConfigProvider.GetString(userJsonPath + ".SignatureVersion")
	if r.SignatureVersion == "" {
		err = errors.New("'SignatureVersion' cannot be empty")
		return
	}
	r.AccessKey = x.ConfigProvider.GetString(userJsonPath + ".AccessKey")
	if r.AccessKey == "" {
		err = errors.New("'AccessKey' cannot be empty")
		return
	}
	r.AccessSecret = x.ConfigProvider.GetString(userJsonPath + ".AccessSecret")
	if r.AccessSecret == "" {
		err = errors.New("'AccessSecret' cannot be empty")
		return
	}
	r.AuthToken = x.ConfigProvider.GetString(userJsonPath + ".AuthToken")

	return
}
