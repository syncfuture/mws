package mws

import (
	"strconv"

	"github.com/syncfuture/go/config"
)

type ReportsAPI struct {
	apiBase
}

func newReportsAPI(seller string, configProvider config.IConfigProvider) *ReportsAPI {
	r := new(ReportsAPI)
	r.Seller = seller
	r.ConfigProvider = configProvider
	r.Module = r.ConfigProvider.GetStringDefault("Reports.Module", "Reports")
	r.Version = r.ConfigProvider.GetStringDefault("Reports.Version", "2009-01-01")
	return r
}

type GetReportListQuery struct {
	queryBase
	MaxCount            string
	ReportTypeListTypes []string
}

func (x *ReportsAPI) GetReportList(query *GetReportListQuery) (string, error) {
	client := x.newClient("GetReportList")

	client.setParameter("MaxCount", query.MaxCount)
	for i, v := range query.ReportTypeListTypes {
		client.setParameter("ReportTypeList.Type."+strconv.Itoa(i+1), v)
	}

	return client.Get()
}

type GetReportQuery struct {
	queryBase
	ReportID string
}

func (x *ReportsAPI) GetReport(query *GetReportQuery) (string, error) {
	client := x.newClient("GetReport")

	client.setParameter("ReportId", query.ReportID)

	return client.Get()
}
