package shipments

import (
	"github.com/syncfuture/mws/core"
)

// #region Interface

type IInboundShipmentsResponse interface {
	GetInboundShipmentsResult() *InboundShipmentsResult
	GetResponseMetadata() *core.ResponseMetadata
	GetError() *core.ResponseError
}

type IInboundShipmentItemsResponse interface {
	GetInboundShipmentItemsResult() *InboundShipmentItemsResponse
	GetResponseMetadata() *core.ResponseMetadata
	GetError() *core.ResponseError
}

type IReportResponse interface {
	GetReportResult() interface{}
	GetError() *core.ResponseError
}

// #endregion

// #region ListInboundShipments

type InboundShipmentsResponse struct {
	InboundShipmentsResult *InboundShipmentsResult `xml:"ListInboundShipmentsResult"`
	ResponseMetadata       *core.ResponseMetadata
	Error                  *core.ResponseError
}
type InboundShipmentsResult struct {
	NextToken    string
	ShipmentData *ShipmentData `xml:"ShipmentData"`
}
type ShipmentData struct {
	Members []*ShipmentMember `xml:"member"`
}
type ShipmentMember struct {
	ShipmentId                     string   `xml:"ShipmentId"`
	ShipmentName                   string   `xml:"ShipmentName"`
	ShipmentStatus                 string   `xml:"ShipmentStatus"`
	LabelPrepType                  string   `xml:"LabelPrepType"`
	DestinationFulfillmentCenterId string   `xml:"DestinationFulfillmentCenterId"`
	ShipFromAddress                *Address `xml:"ShipFromAddress"`
}
type Address struct {
	Name                string `xml:"Name"`
	AddressLine1        string `xml:"AddressLine1"`
	AddressLine2        string `xml:"AddressLine2"`
	City                string `xml:"City"`
	DistrictOrCounty    string `xml:"DistrictOrCounty"`
	StateOrProvinceCode string `xml:"StateOrProvinceCode"`
	CountryCode         string `xml:"CountryCode"`
	PostalCod           string `xml:"PostalCod"`
}

func (x *InboundShipmentsResponse) GetInboundShipmentsResult() *InboundShipmentsResult {
	return x.InboundShipmentsResult
}
func (x *InboundShipmentsResponse) GetResponseMetadata() *core.ResponseMetadata {
	return x.ResponseMetadata
}
func (x *InboundShipmentsResponse) GetError() *core.ResponseError {
	return x.Error
}

type InboundShipmentsByNextTokenResponse struct {
	InboundShipmentsResult *InboundShipmentsResult `xml:"ListInboundShipmentsByNextTokenResult"`
	ResponseMetadata       *core.ResponseMetadata
	Error                  *core.ResponseError
}

func (x *InboundShipmentsByNextTokenResponse) GetInboundShipmentsResult() *InboundShipmentsResult {
	return x.InboundShipmentsResult
}
func (x *InboundShipmentsByNextTokenResponse) GetResponseMetadata() *core.ResponseMetadata {
	return x.ResponseMetadata
}
func (x *InboundShipmentsByNextTokenResponse) GetError() *core.ResponseError {
	return x.Error
}

// #endregion

// #region ListInboundShipmentItems

type InboundShipmentItemsResponse struct {
	InboundShipmentItemsResult *InboundShipmentItemsResult `xml:"ListInboundShipmentItemsResult"`
	ResponseMetadata           *core.ResponseMetadata
	Error                      *core.ResponseError
}
type InboundShipmentItemsResult struct {
	NextToken string
	ItemData  *ItemData `xml:"ItemData"`
}
type ItemData struct {
	Members []*ItemMember `xml:"member"`
}
type ItemMember struct {
	ShipmentId            string   `xml:"ShipmentId"`
	SellerSKU             string   `xml:"SellerSKU"`
	FulfillmentNetworkSKU string   `xml:"FulfillmentNetworkSKU"`
	QuantityShipped       int      `xml:"QuantityShipped"`
	QuantityReceived      int      `xml:"QuantityReceived"`
	QuantityInCase        int      `xml:"QuantityInCase"`
	ReleaseDate           string   `xml:"ReleaseDate"`
	ShipFromAddress       *Address `xml:"ShipFromAddrss"`
}

func (x *InboundShipmentItemsResponse) GetInboundShipmentItemsResult() *InboundShipmentItemsResult {
	return x.InboundShipmentItemsResult
}
func (x *InboundShipmentItemsResponse) GetResponseMetadata() *core.ResponseMetadata {
	return x.ResponseMetadata
}
func (x *InboundShipmentItemsResponse) GetError() *core.ResponseError {
	return x.Error
}

type InboundShipmentItemsByNextTokenResponse struct {
	InboundShipmentItemsResult *InboundShipmentItemsResult `xml:"ListInboundShipmentItemsByNextTokenResult"`
	ResponseMetadata           *core.ResponseMetadata
	Error                      *core.ResponseError
}

func (x *InboundShipmentItemsByNextTokenResponse) GetReportListResult() *InboundShipmentItemsResult {
	return x.InboundShipmentItemsResult
}
func (x *InboundShipmentItemsByNextTokenResponse) GetResponseMetadata() *core.ResponseMetadata {
	return x.ResponseMetadata
}
func (x *InboundShipmentItemsByNextTokenResponse) GetError() *core.ResponseError {
	return x.Error
}

// #endregion

// #region GetTransportContent

type TransportContentResponse struct {
	TransportContentResult *TransportContentResult `xml:"GetTransportContentResult"`
	ResponseMetadata       *core.ResponseMetadata
	Error                  *core.ResponseError
}
type TransportContentResult struct {
	TransportContent *TransportContent `xml:"TransportContent"`
}
type TransportContent struct {
	TransportDetails *TransportDetails `xml:"TransportDetails"`
}
type TransportDetails struct {
	PartneredLtlData         *PartneredLtlData         `xml:"PartneredLtlData"`
	PartneredSmallParcelData *PartneredSmallParcelData `xml:"PartneredSmallParcelData"`
}
type PartneredLtlData struct {
	PartneredEstimate *PartneredEstimate `xml:"PartneredEstimate"`
}
type PartneredSmallParcelData struct {
	PartneredEstimate *PartneredEstimate `xml:"PartneredEstimate"`
}
type PartneredEstimate struct {
	Amount *Amount `xml:"Amount"`
}
type Amount struct {
	CurrencyCode string `xml:"CurrencyCode"`
	Value        string `xml:"Value"`
}

func (x *TransportContentResponse) GetTransportContentResponseResult() *TransportContentResult {
	return x.TransportContentResult
}
func (x *TransportContentResponse) GetResponseMetadata() *core.ResponseMetadata {
	return x.ResponseMetadata
}
func (x *TransportContentResponse) GetError() *core.ResponseError {
	return x.Error
}

// #endregion
