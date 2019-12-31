package reports

import (
	"strconv"

	"github.com/syncfuture/mws/core"

	"github.com/syncfuture/go/config"
)

type ReportsAPI struct {
	core.APIBase
}

func NewReportsAPI(seller string, configProvider config.IConfigProvider) *ReportsAPI {
	r := new(ReportsAPI)
	r.Seller = seller
	r.ConfigProvider = configProvider
	r.Module = r.ConfigProvider.GetStringDefault("Reports.Module", "Reports")
	r.Version = r.ConfigProvider.GetStringDefault("Reports.Version", "2009-01-01")
	return r
}

type GetReportListQuery struct {
	core.QueryBase
	MaxCount            string
	ReportTypeListTypes []string
}

func (x *ReportsAPI) GetReportList(query *GetReportListQuery) (string, error) {
	client := x.NewClient("GetReportList")

	client.SetParameter("MaxCount", query.MaxCount)
	for i, v := range query.ReportTypeListTypes {
		client.SetParameter("ReportTypeList.Type."+strconv.Itoa(i+1), v)
	}

	return client.Get()
}

type GetReportQuery struct {
	core.QueryBase
	ReportID string
}

func (x *ReportsAPI) GetReport(query *GetReportQuery) (string, error) {
	client := x.NewClient("GetReport")

	client.SetParameter("ReportId", query.ReportID)

	return client.Get()
}
