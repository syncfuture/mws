package mws

import (
	"github.com/syncfuture/go/config"
)

type FinancesAPI struct {
	apiBase
}

func newFinancesAPI(seller string, configProvider config.IConfigProvider) *FinancesAPI {
	r := new(FinancesAPI)
	r.Seller = seller
	r.ConfigProvider = configProvider
	r.Module = r.ConfigProvider.GetStringDefault("Orders.Module", "Finances")
	r.Version = r.ConfigProvider.GetStringDefault("Orders.Version", "2015-05-01")
	return r
}

type ListFanancialEventsQuery struct {
	queryBase
	MaxResultsPerPage     string
	AmazonOrderId         string
	FinancialEventGroupId string
	PostedAfter           string
	PostedBefore          string
}

func (x *FinancesAPI) ListFinancialEvents(query *ListFanancialEventsQuery) (string, error) {
	client := x.newClient("ListFinancialEvents")

	client.setParameter("MaxResultsPerPage", query.MaxResultsPerPage)
	client.setParameter("PostedAfter", query.PostedAfter)

	return client.Get()
}

type ListFinancialEventsByNextTokenQuery struct {
	queryBase
	NextToken string
}

func (x *FinancesAPI) ListFinancialEventsByNextToken(query *ListFinancialEventsByNextTokenQuery) (string, error) {
	client := x.newClient("ListFinancialEventsByNextToken")

	client.setParameter("NextToken", query.NextToken)

	return client.Get()
}
