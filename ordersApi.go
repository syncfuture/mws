package mws

import "github.com/syncfuture/go/config"

type OrdersAPI struct {
	apiBase
}

func newOrdersAPI(seller string, configProvider config.IConfigProvider) *OrdersAPI {
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
	queryBase
	CreatedAfter string
}

func (x *OrdersAPI) ListOrders(query *ListOrdersQuery) (string, error) {
	client := x.newClient("ListOrders")

	client.setParameter("CreatedAfter", query.CreatedAfter)
	client.setParameter("MarketplaceId.Id.1", client.MarketplaceID)

	return client.Get()
}
