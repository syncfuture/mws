package mws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReportList(t *testing.T) {
	client := NewReportsAPI("LF")
	xml, err := client.GetReportList(&GetReportListQuery{
		MaxCount:            "1",
		ReportTypeListTypes: []string{"_GET_FBA_MYI_UNSUPPRESSED_INVENTORY_DATA_"},
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, xml)
	t.Log(xml)
}

func TestGetReport(t *testing.T) {
	client := NewReportsAPI("LF")
	xml, err := client.GetReport(&GetReportQuery{
		ReportID: "18150051229018249",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, xml)
	t.Log(xml)
}
