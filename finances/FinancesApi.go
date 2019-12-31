package finances

import (
	"github.com/syncfuture/go/config"
	"github.com/syncfuture/mws/core"
)

type FinancesAPI struct {
	core.APIBase
}

func NewFinancesAPI(seller string, configProvider config.IConfigProvider) *FinancesAPI {
	r := new(FinancesAPI)
	r.Seller = seller
	r.ConfigProvider = configProvider
	r.Module = r.ConfigProvider.GetStringDefault("Orders.Module", "Finances")
	r.Version = r.ConfigProvider.GetStringDefault("Orders.Version", "2015-05-01")
	return r
}

type ListFanancialEventsQuery struct {
	core.QueryBase
	MaxResultsPerPage     string
	AmazonOrderId         string
	FinancialEventGroupId string
	PostedAfter           string
	PostedBefore          string
}

func (x *FinancesAPI) ListFinancialEvents(query *ListFanancialEventsQuery) (string, error) {
	client := x.NewClient("ListFinancialEvents")

	client.SetParameter("MaxResultsPerPage", query.MaxResultsPerPage)
	client.SetParameter("PostedAfter", query.PostedAfter)
	client.SetParameter("PostedBefore", query.PostedBefore)

	return client.Get()
}

type ListFinancialEventsByNextTokenQuery struct {
	core.QueryBase
	NextToken string
}

func (x *FinancesAPI) ListFinancialEventsByNextToken(query *ListFinancialEventsByNextTokenQuery) (string, error) {
	client := x.NewClient("ListFinancialEventsByNextToken")

	client.SetParameter("NextToken", query.NextToken)

	return client.Get()
}
