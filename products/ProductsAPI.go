package products

import (
	"strconv"

	mwsconfig "github.com/syncfuture/mws/config"
	"github.com/syncfuture/mws/core"
)

type ProductsAPI struct {
	core.APIBase
}

func NewProductsAPI(config *mwsconfig.MWSConfig, args ...string) *ProductsAPI {
	r := new(ProductsAPI)
	r.Config = config
	r.Module = "Products"
	if len(args) > 0 {
		r.Version = args[0]
	} else {
		r.Version = "2011-10-01"
	}
	return r
}

type GetMatchingProductQuery struct {
	core.QueryBase
	MarketplaceId string
	ASINList      []string
}

func (x *ProductsAPI) GetMatchingProduct(query *GetMatchingProductQuery) (r string, err error) {
	var client *core.MWSClient
	client, err = x.NewClient("GetMatchingProduct")

	client.SetParameter("MarketplaceId", client.MarketplaceID)
	for i, v := range query.ASINList {
		client.SetParameter("ASINList.ASIN."+strconv.Itoa(i+1), v)
	}

	return client.Get()
}
