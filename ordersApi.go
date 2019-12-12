package mws

type ordersApi struct {
	api
}

func NewOrdersApi() *ordersApi {
	r := new(ordersApi)
	r.Module = ConfigProvider.GetStringDefault("Orders.Module", "Orders")
	r.Version = ConfigProvider.GetStringDefault("Orders.Version", "2013-09-01")
	return r
}

type ListOrdersQuery struct {
	CreatedAfter string
}

func (x *ordersApi) ListOrders(query *ListOrdersQuery) (string, error) {
	api := newAPI(x.Module, x.Version, "ListOrders")

	api.setParameter("CreatedAfter", query.CreatedAfter)
	api.setParameter("MarketplaceId.Id.1", api.MarketplaceID)

	return api.Get()
}
