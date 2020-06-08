package mws

import (
	"github.com/syncfuture/mws/finances"
	"github.com/syncfuture/mws/orders"
	"github.com/syncfuture/mws/products"
	"github.com/syncfuture/mws/protoc/mwsconfig"
	"github.com/syncfuture/mws/reports"
)

type APISet struct {
	Finances *finances.FinancesAPI
	Orders   *orders.OrdersAPI
	Reports  *reports.ReportsAPI
	Products *products.ProductsAPI
}

func NewAPISet(config *mwsconfig.MWSConfig) (r *APISet) {
	r = new(APISet)
	r.Finances = finances.NewFinancesAPI(config)
	r.Orders = orders.NewOrdersAPI(config)
	r.Reports = reports.NewReportsAPI(config)
	r.Products = products.NewProductsAPI(config)
	return r
}
