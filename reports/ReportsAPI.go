package reports

import (
	"github.com/syncfuture/go/u"
	"github.com/syncfuture/mws/core"
	"github.com/syncfuture/mws/protoc/mwsconfig"
)

type ReportsAPI struct {
	core.APIBase
}

func NewReportsAPI(config *mwsconfig.MWSConfig, args ...string) *ReportsAPI {
	r := new(ReportsAPI)
	r.Config = config
	r.Module = "Reports"
	if len(args) > 0 {
		r.Version = args[0]
	} else {
		r.Version = "2009-01-01"
	}
	return r
}

type RequestReportQuery struct {
	core.QueryBase
	ReportType string
	StartDate  string
	EndDate    string
}

func (x *ReportsAPI) RequestReport(query *RequestReportQuery) (r string, err error) {
	var client *core.MWSClient
	client, err = x.NewClient("RequestReport")

	client.SetParameter("ReportType", query.ReportType)
	client.SetParameter("StartDate", query.StartDate)
	client.SetParameter("EndDate", query.EndDate)
	client.SetParameter("MarketplaceIdList.Id.1", client.MarketplaceID)

	return client.Get()
}

type GetReportRequestListQuery struct {
	core.QueryBase
	MaxCount        string
	ReportRequestId string
}

func (x *ReportsAPI) GetReportRequestList(query *GetReportRequestListQuery) (r string, err error) {
	var client *core.MWSClient
	client, err = x.NewClient("GetReportRequestList")

	client.SetParameter("MaxCount", query.MaxCount)
	client.SetParameter("ReportRequestIdList.Id.1", query.ReportRequestId)
	// for i, v := range query.ReportTypeList {
	// 	client.SetParameter("ReportRequestIdList.Type."+strconv.Itoa(i+1), v)
	// }

	return client.Get()
}

type GetReportRequestListByNextTokenQuery struct {
	core.QueryBase
	NextToken string
}

func (x *ReportsAPI) GetReportRequestListByNextToken(query *GetReportRequestListByNextTokenQuery) (r string, err error) {
	var client *core.MWSClient
	client, err = x.NewClient("GetReportRequestListByNextToken")
	if u.LogError(err) {
		return
	}

	client.SetParameter("NextToken", query.NextToken)

	return client.Get()
}

type GetReportQuery struct {
	core.QueryBase
	ReportID string
}

func (x *ReportsAPI) GetReport(query *GetReportQuery) (r string, err error) {
	var client *core.MWSClient
	client, err = x.NewClient("GetReport")

	client.SetParameter("ReportId", query.ReportID)

	return client.Get()
}
