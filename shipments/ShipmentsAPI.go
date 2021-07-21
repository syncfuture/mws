package shipments

import (
	"github.com/syncfuture/go/sconv"
	"github.com/syncfuture/go/u"
	"github.com/syncfuture/mws/core"
	"github.com/syncfuture/mws/protoc/mwsconfig"
)

type ShipmentsAPI struct {
	core.APIBase
}

func NewShipmentsAPI(config *mwsconfig.MWSConfig, args ...string) *ShipmentsAPI {
	r := new(ShipmentsAPI)
	r.Config = config
	r.Module = "FulfillmentInboundShipment"
	if len(args) > 0 {
		r.Version = args[0]
	} else {
		r.Version = "2010-10-01"
	}
	return r
}

// #region ListInboundShipments

type ListInboundShipmentsQuery struct {
	core.QueryBase
	ShipmentStatusList []string
	LastUpdatedAfter   string
	LastUpdatedBefore  string
}

type ListInboundShipmentsByNextTokenQuery struct {
	core.QueryBase
	NextToken string
}

func (x *ShipmentsAPI) ListInboundShipments(query *ListInboundShipmentsQuery) (r string, err error) {
	var client *core.MWSClient
	client, err = x.NewClient("ListInboundShipments")

	if u.IsMissing(query.ShipmentStatusList) {
		// default
		query.ShipmentStatusList = []string{"WORKING", "SHIPPED", "IN_TRANSIT", "DELIVERED", "CHECKED_IN", "RECEIVING", "CLOSED"}
	}

	client.SetParameter("LastUpdatedAfter", query.LastUpdatedAfter)
	client.SetParameter("LastUpdatedBefore", query.LastUpdatedBefore)
	for i, item := range query.ShipmentStatusList {
		client.SetParameter("ShipmentStatusList.member."+sconv.ToString(i+1), item)
	}

	return client.Get()
}

func (x *ShipmentsAPI) ListInboundShipmentsByNextToken(query *ListInboundShipmentsByNextTokenQuery) (r string, err error) {
	var client *core.MWSClient
	client, err = x.NewClient("ListInboundShipmentsByNextToken")
	client.SetParameter("NextToken", query.NextToken)

	return client.Get()
}

// #endregion

// #region ListInboundShipmentItems

type ListInboundShipmentItemsQuery struct {
	core.QueryBase
	ShipmentId string
}

type ListInboundShipmentItemsByNextTokenQuery struct {
	core.QueryBase
	NextToken string
}

func (x *ShipmentsAPI) ListInboundShipmentItems(query *ListInboundShipmentItemsQuery) (r string, err error) {
	var client *core.MWSClient
	client, err = x.NewClient("ListInboundShipmentItems")
	client.SetParameter("ShipmentId", query.ShipmentId)

	return client.Get()
}

func (x *ShipmentsAPI) ListInboundShipmentItemsByNextToken(query *ListInboundShipmentItemsByNextTokenQuery) (r string, err error) {
	var client *core.MWSClient
	client, err = x.NewClient("ListInboundShipmentItemsByNextToken")
	client.SetParameter("NextToken", query.NextToken)

	return client.Get()
}

// #endregion

// #region GetTransportContent

type GetTransportContentQuery struct {
	core.QueryBase
	ShipmentId string
}

func (x *ShipmentsAPI) GetTransportContent(query *GetTransportContentQuery) (r string, err error) {
	var client *core.MWSClient
	client, err = x.NewClient("GetTransportContent")
	client.SetParameter("ShipmentId", query.ShipmentId)

	return client.Get()
}

// #endregion
