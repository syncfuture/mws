package tests

import (
	"testing"

	"github.com/syncfuture/mws/reports"

	"github.com/stretchr/testify/assert"
)

func TestGetReportList(t *testing.T) {
	xml, err := _apiSet.Reports.GetReportList(&reports.GetReportListQuery{
		MaxCount: "10",
		// ReportTypeListTypes: []string{"_GET_FBA_MYI_UNSUPPRESSED_INVENTORY_DATA_"},
		ReportTypeListTypes: []string{"_GET_FLAT_FILE_ALL_ORDERS_DATA_BY_ORDER_DATE_", "_GET_FBA_FULFILLMENT_CUSTOMER_RETURNS_DATA_", "_GET_MERCHANT_LISTINGS_ALL_DATA_"},
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, xml)
	t.Log(xml)
}

func TestGetReport(t *testing.T) {
	xml, err := _apiSet.Reports.GetReport(&reports.GetReportQuery{
		ReportID: "20512864646018387",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, xml)
	t.Log(xml)
}
