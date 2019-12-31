package orders

import (
	"github.com/syncfuture/go/config"
	"github.com/syncfuture/mws/core"
)

type OrdersAPI struct {
	core.APIBase
}

func NewOrdersAPI(seller string, configProvider config.IConfigProvider) *OrdersAPI {
	if seller == "" {
		panic("seller cannot be empty")
	}
	r := new(OrdersAPI)
	r.Seller = seller
	r.ConfigProvider = configProvider
	r.Module = r.ConfigProvider.GetStringDefault("Orders.Module", "Orders")
	r.Version = r.ConfigProvider.GetStringDefault("Orders.Version", "2013-09-01")
	return r
}

type ListOrdersQuery struct {
	core.QueryBase
	CreatedAfter string
}

func (x *OrdersAPI) ListOrders(query *ListOrdersQuery) (string, error) {
	client := x.NewClient("ListOrders")

	client.SetParameter("CreatedAfter", query.CreatedAfter)
	client.SetParameter("MarketplaceId.Id.1", client.MarketplaceID)

	return client.Get()
}
