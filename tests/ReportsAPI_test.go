package tests

import (
	"encoding/csv"
	"strings"
	"testing"
	"time"

	"github.com/gocarina/gocsv"

	"github.com/syncfuture/mws/reports"

	"github.com/stretchr/testify/assert"
)

func TestRequestReport(t *testing.T) {
	a, b, c := time.Now().UTC().Date()
	date := time.Date(a, b, c, 0, 0, 0, 0, &time.Location{})

	// ReportTypeList: []string{"_GET_FLAT_FILE_ALL_ORDERS_DATA_BY_ORDER_DATE_", "_GET_FBA_FULFILLMENT_CUSTOMER_RETURNS_DATA_", "_GET_MERCHANT_LISTINGS_ALL_DATA_", "_GET_FBA_ESTIMATED_FBA_FEES_TXT_DATA_"},
	xml, err := _apiSet.Reports.RequestReport(&reports.RequestReportQuery{
		ReportType: "_GET_FLAT_FILE_ALL_ORDERS_DATA_BY_ORDER_DATE_",
		StartDate:  date.Add(-24 * 2 * time.Hour).Format(time.RFC3339),
		EndDate:    date.Add(-24 * 1 * time.Hour).Format(time.RFC3339),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, xml)
	t.Log(xml)
}

func TestGetReportList(t *testing.T) {
	xml, err := _apiSet.Reports.GetReportRequestList(&reports.GetReportRequestListQuery{
		MaxCount:        "100",
		ReportRequestId: "281727018432",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, xml)
	t.Log(xml)
}

func TestGetReport(t *testing.T) {
	xml, err := _apiSet.Reports.GetReport(&reports.GetReportQuery{
		ReportID: "21346085966018432",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, xml)
	t.Log(xml)

	var resp []*reports.AllOrdersReport
	reader := csv.NewReader(strings.NewReader(xml))
	reader.Comma = '\t'
	reader.LazyQuotes = true

	err = gocsv.UnmarshalCSV(reader, &resp)
	if err != nil {
		t.Log(err.Error())
	}

	t.Log(len(resp))
}
