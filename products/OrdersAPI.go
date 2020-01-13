package products

import (
	"strconv"

	"github.com/syncfuture/go/config"
	"github.com/syncfuture/mws/core"
)

type ProductsAPI struct {
	core.APIBase
}

func NewProductsAPI(seller string, configProvider config.IConfigProvider) *ProductsAPI {
	if seller == "" {
		panic("seller cannot be empty")
	}
	r := new(ProductsAPI)
	r.Seller = seller
	r.ConfigProvider = configProvider
	r.Module = r.ConfigProvider.GetStringDefault("Products.Module", "Products")
	r.Version = r.ConfigProvider.GetStringDefault("Products.Version", "2011-10-01")
	return r
}

type GetMatchingProductQuery struct {
	core.QueryBase
	MarketplaceId string
	ASINList      []string
}

func (x *ProductsAPI) GetMatchingProduct(query *GetMatchingProductQuery) (string, error) {
	client := x.NewClient("GetMatchingProduct")

	client.SetParameter("MarketplaceId", client.MarketplaceID)
	for i, v := range query.ASINList {
		client.SetParameter("ASINList.ASIN."+strconv.Itoa(i+1), v)
	}

	return client.Get()
}
