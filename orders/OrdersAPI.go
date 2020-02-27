package orders

import (
	mwsconfig "github.com/syncfuture/mws/config"
	"github.com/syncfuture/mws/core"
)

type OrdersAPI struct {
	core.APIBase
}

func NewOrdersAPI(config *mwsconfig.MWSConfig, args ...string) *OrdersAPI {
	r := new(OrdersAPI)
	r.Config = config
	r.Module = "Orders"
	if len(args) > 0 {
		r.Version = args[0]
	} else {
		r.Version = "2013-09-01"
	}
	return r
}

type ListOrdersQuery struct {
	core.QueryBase
	CreatedAfter string
}

func (x *OrdersAPI) ListOrders(query *ListOrdersQuery) (r string, err error) {
	var client *core.MWSClient
	client, err = x.NewClient("ListOrders")

	client.SetParameter("CreatedAfter", query.CreatedAfter)
	client.SetParameter("MarketplaceId.Id.1", client.MarketplaceID)

	return client.Get()
}
