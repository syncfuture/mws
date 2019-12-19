package mws

type ordersAPI struct {
	apiBase
}

func NewOrdersAPI(seller string) *ordersAPI {
	if seller == "" {
		panic("seller cannot be empty")
	}
	r := new(ordersAPI)
	r.Seller = seller
	r.Module = ConfigProvider.GetStringDefault("Orders.Module", "Orders")
	r.Version = ConfigProvider.GetStringDefault("Orders.Version", "2013-09-01")
	return r
}

type ListOrdersQuery struct {
	queryBase
	CreatedAfter string
}

func (x *ordersAPI) ListOrders(query *ListOrdersQuery) (string, error) {
	client := x.newClient("ListOrders")

	client.setParameter("CreatedAfter", query.CreatedAfter)
	client.setParameter("MarketplaceId.Id.1", client.MarketplaceID)

	return client.Get()
}
