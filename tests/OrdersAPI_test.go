package tests

import (
	"testing"
	"time"

	"github.com/subchen/go-xmldom"
	"github.com/syncfuture/go/u"
	"github.com/syncfuture/mws/orders"

	"github.com/stretchr/testify/assert"
)

func TestListOrders(t *testing.T) {
	xml, err := _apiSet.Orders.ListOrders(&orders.ListOrdersQuery{
		CreatedAfter: time.Now().UTC().Format(time.RFC3339),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, xml)
}

func TestGetOrder(t *testing.T) {
	xml, err := _apiSet.Orders.GetOrder(&orders.GetOrderQuery{
		AmazonOrderIDList: []string{
			// "114-4159787-5585006",
			"114-4129208-8882629",
			"114-4120921-6108219",
			"114-4111432-0434665",
			"114-4085915-7674634",
			"114-4083763-9206640",
			"114-4029893-1472205",
			"114-3985589-4618667",
			"114-3922333-4654606",
			"702-9037655-7457068",
			"702-8522828-2001837",
			"702-6646996-8193056",
			"702-4949684-9505058",
			"702-4822303-5971455",
			"702-4221199-2497009",
		},
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, xml)
	t.Log(xml)

	doc, err := xmldom.ParseXML(xml)
	if u.LogError(err) {
		return
	}

	path := "GetOrderResult/Orders/Order"
	nodes := doc.Root.Query(path)
	for _, node := range nodes {
		date := node.QueryOne("PurchaseDate").Text
		t.Log(date)
	}
}

// func TestGetOrder2(t *testing.T) {
// 	url := ""

// 	resp, err := http.Get(url)
// }
