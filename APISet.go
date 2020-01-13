package mws

import (
	"github.com/syncfuture/go/config"
	"github.com/syncfuture/mws/finances"
	"github.com/syncfuture/mws/orders"
	"github.com/syncfuture/mws/products"
	"github.com/syncfuture/mws/reports"
)

type APISet struct {
	Finances *finances.FinancesAPI
	Orders   *orders.OrdersAPI
	Reports  *reports.ReportsAPI
	Products *products.ProductsAPI
}

func NewAPISet(seller string, configProvider config.IConfigProvider) (r *APISet) {
	r = new(APISet)
	r.Finances = finances.NewFinancesAPI(seller, configProvider)
	r.Orders = orders.NewOrdersAPI(seller, configProvider)
	r.Reports = reports.NewReportsAPI(seller, configProvider)
	r.Products = products.NewProductsAPI(seller, configProvider)
	return r
}
