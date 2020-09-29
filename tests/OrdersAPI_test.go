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
		AmazonOrderIDList: []string{"113-5364786-9457859", "111-9554494-4869823"},
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, xml)
	t.Log(xml)

	doc, err := xmldom.ParseXML(xml)
	if u.LogError(err) {
		return
	}

	path := "GetOrderResult/Orders"
	nodes := doc.Root.Query(path)
	for _, node := range nodes {
		date := node.QueryOne("Order/PurchaseDate").Text
		t.Log(date)
	}
}
