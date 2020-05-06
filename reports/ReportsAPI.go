package reports

import (
	"strconv"

	"github.com/syncfuture/go/u"
	mwsconfig "github.com/syncfuture/mws/config"
	"github.com/syncfuture/mws/core"
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

type GetReportListQuery struct {
	core.QueryBase
	MaxCount            string
	ReportTypeListTypes []string
	AvailableFromDate   string
	AvailableToDate     string
}

func (x *ReportsAPI) GetReportList(query *GetReportListQuery) (r string, err error) {
	var client *core.MWSClient
	client, err = x.NewClient("GetReportList")

	client.SetParameter("MaxCount", query.MaxCount)
	client.SetParameter("AvailableFromDate", query.AvailableFromDate)
	client.SetParameter("AvailableToDate", query.AvailableToDate)
	for i, v := range query.ReportTypeListTypes {
		client.SetParameter("ReportTypeList.Type."+strconv.Itoa(i+1), v)
	}

	return client.Get()
}

type GetReportListByNextTokenQuery struct {
	core.QueryBase
	NextToken string
}

func (x *ReportsAPI) GetReportListByNextToken(query *GetReportListByNextTokenQuery) (r string, err error) {
	var client *core.MWSClient
	client, err = x.NewClient("GetReportListByNextToken")
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
