package finances

import (
	"github.com/syncfuture/go/u"
	mwsconfig "github.com/syncfuture/mws/config"
	"github.com/syncfuture/mws/core"
)

type FinancesAPI struct {
	core.APIBase
}

func NewFinancesAPI(config *mwsconfig.MWSConfig, args ...string) *FinancesAPI {
	r := new(FinancesAPI)
	r.Config = config
	r.Module = "Finances"
	if len(args) > 0 {
		r.Version = args[0]
	} else {
		r.Version = "2015-05-01"
	}
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

func (x *FinancesAPI) ListFinancialEvents(query *ListFanancialEventsQuery) (r string, err error) {
	var client *core.MWSClient
	client, err = x.NewClient("ListFinancialEvents")
	if u.LogError(err) {
		return
	}

	client.SetParameter("MaxResultsPerPage", query.MaxResultsPerPage)
	client.SetParameter("PostedAfter", query.PostedAfter)
	client.SetParameter("PostedBefore", query.PostedBefore)

	return client.Get()
}

type ListFinancialEventsByNextTokenQuery struct {
	core.QueryBase
	NextToken string
}

func (x *FinancesAPI) ListFinancialEventsByNextToken(query *ListFinancialEventsByNextTokenQuery) (r string, err error) {
	var client *core.MWSClient
	client, err = x.NewClient("ListFinancialEventsByNextToken")
	if u.LogError(err) {
		return
	}

	client.SetParameter("NextToken", query.NextToken)

	return client.Get()
}
