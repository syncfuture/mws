package tests

import (
	"encoding/csv"
	"strings"
	"testing"

	"github.com/gocarina/gocsv"

	"github.com/syncfuture/mws/reports"

	"github.com/stretchr/testify/assert"
)

func TestGetReportList(t *testing.T) {
	xml, err := _apiSet.Reports.GetReportList(&reports.GetReportListQuery{
		MaxCount:            "100",
		AvailableFromDate:   "2020-05-11T00:29:53+00:00",
		ReportTypeListTypes: []string{"_GET_FBA_ESTIMATED_FBA_FEES_TXT_DATA_"},
		// ReportTypeListTypes: []string{"_GET_FLAT_FILE_ALL_ORDERS_DATA_BY_ORDER_DATE_", "_GET_FBA_FULFILLMENT_CUSTOMER_RETURNS_DATA_", "_GET_MERCHANT_LISTINGS_ALL_DATA_"},
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, xml)
	t.Log(xml)
}

func TestGetReport(t *testing.T) {
	xml, err := _apiSet.Reports.GetReport(&reports.GetReportQuery{
		ReportID: "20609622731018393",
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

	t.Log(resp)
}
