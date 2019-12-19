package mws

import (
	"strconv"
)

type reportsAPI struct {
	apiBase
}

func NewReportsAPI(seller string) *reportsAPI {
	r := new(reportsAPI)
	r.Seller = seller
	r.Module = ConfigProvider.GetStringDefault("Reports.Module", "Reports")
	r.Version = ConfigProvider.GetStringDefault("Reports.Version", "2009-01-01")
	return r
}

type GetReportListQuery struct {
	queryBase
	MaxCount            string
	ReportTypeListTypes []string
}

func (x *reportsAPI) GetReportList(query *GetReportListQuery) (string, error) {
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

func (x *reportsAPI) GetReport(query *GetReportQuery) (string, error) {
	client := x.newClient("GetReport")

	client.setParameter("ReportId", query.ReportID)

	return client.Get()
}
