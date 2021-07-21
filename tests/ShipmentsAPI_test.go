package tests

import (
	"testing"
	"time"

	"github.com/syncfuture/mws/shipments"

	"github.com/stretchr/testify/assert"
)

func TestListInboundShipments(t *testing.T) {
	a, b, c := time.Now().UTC().Date()
	date := time.Date(a, b, c, 0, 0, 0, 0, &time.Location{})

	xml, err := _apiSet.Shipments.ListInboundShipments(&shipments.ListInboundShipmentsQuery{
		//ShipmentStatusList: []string{"WORKING", "SHIPPED", "IN_TRANSIT", "DELIVERED", "CHECKED_IN", "RECEIVING", "CLOSED"},
		LastUpdatedAfter:  date.Add(-365 * 24 * 90 * time.Hour).Format(time.RFC3339),
		LastUpdatedBefore: date.Add(-24 * 1 * time.Hour).Format(time.RFC3339),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, xml)
	t.Log(xml)
}

func TestListInboundShipmentItems(t *testing.T) {
	xmlStr, err := _apiSet.Shipments.ListInboundShipmentItems(&shipments.ListInboundShipmentItemsQuery{
		ShipmentId: "FBA4GMLMFS",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, xmlStr)
	t.Log(xmlStr)

	// resp := new(shipments.InboundShipmentItemsResponse)
	// data := u.StrToBytes(xmlStr)
	// err = xml.Unmarshal(data, resp)
	// if err != nil {
	// 	t.Log(err.Error())
	// }
	// t.Log(resp)
}

func TestGetTransportContent(t *testing.T) {
	xml, err := _apiSet.Shipments.GetTransportContent(&shipments.GetTransportContentQuery{
		ShipmentId: "FBA4GMLMFS",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, xml)
	t.Log(xml)
}
