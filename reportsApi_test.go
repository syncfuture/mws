package mws

import (
	"testing"

	"github.com/syncfuture/mws/reports"

	"github.com/stretchr/testify/assert"
)

func TestGetReportList(t *testing.T) {
	xml, err := _apiSet.Reports.GetReportList(&reports.GetReportListQuery{
		MaxCount:            "1",
		ReportTypeListTypes: []string{"_GET_FBA_MYI_UNSUPPRESSED_INVENTORY_DATA_"},
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, xml)
	t.Log(xml)
}

func TestGetReport(t *testing.T) {
	xml, err := _apiSet.Reports.GetReport(&reports.GetReportQuery{
		ReportID: "18150051229018249",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, xml)
	t.Log(xml)
}
