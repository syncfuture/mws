package mws

import (
	"github.com/syncfuture/go/config"
)

type APISet struct {
	Finances *FinancesAPI
	Orders   *OrdersAPI
	Reports  *ReportsAPI
}

func NewAPISet(seller string, configProvider config.IConfigProvider) (r *APISet) {
	r = new(APISet)
	r.Finances = newFinancesAPI(seller, configProvider)
	r.Orders = newOrdersAPI(seller, configProvider)
	r.Reports = newReportsAPI(seller, configProvider)
	return r
}
