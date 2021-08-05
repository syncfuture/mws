package reports

import (
	"time"

	"github.com/syncfuture/mws/core"
)

// #region Interface

type IRequestReportResponse interface {
	GetRequestReportResult() *RequestReportResult
	GetResponseMetadata() *core.ResponseMetadata
	GetError() *core.ResponseError
}

type IReportRequestListResponse interface {
	GetReportListResult() *ReportRequestListResult
	GetResponseMetadata() *core.ResponseMetadata
	GetError() *core.ResponseError
}

type IReportResponse interface {
	SetReportResult(interface{})
	GetReportResult() interface{}
	GetError() *core.ResponseError
}

// #endregion

// #region RequestReport

type RequestReportResponse struct {
	RequestReportResult *RequestReportResult
	ResponseMetadata    *core.ResponseMetadata
	Error               *core.ResponseError
}
type RequestReportResult struct {
	ReportRequestInfo *ReportRequestInfo
}

type ReportRequestInfo struct {
	ReportType             string
	ReportProcessingStatus string
	ReportRequestId        string
	GeneratedReportId      string
	StartDate              time.Time
	EndDate                time.Time
}

func (x *RequestReportResponse) GetRequestReportResult() *RequestReportResult {
	return x.RequestReportResult
}
func (x *RequestReportResponse) GetResponseMetadata() *core.ResponseMetadata {
	return x.ResponseMetadata
}
func (x *RequestReportResponse) GetError() *core.ResponseError {
	return x.Error
}

// #endregion

// #region ReportRequestList

type ReportRequestListResponse struct {
	ReportRequestListResult *ReportRequestListResult `xml:"GetReportRequestListResult"`
	ResponseMetadata        *core.ResponseMetadata
	Error                   *core.ResponseError
}
type ReportRequestListResult struct {
	NextToken          string
	ReportRequestInfos []*ReportRequestInfo `xml:"ReportRequestInfo"`
}

func (x *ReportRequestListResponse) GetReportListResult() *ReportRequestListResult {
	return x.ReportRequestListResult
}
func (x *ReportRequestListResponse) GetResponseMetadata() *core.ResponseMetadata {
	return x.ResponseMetadata
}
func (x *ReportRequestListResponse) GetError() *core.ResponseError {
	return x.Error
}

type ReportRequestListByNextTokenResponse struct {
	ReportRequestListResult *ReportRequestListResult `xml:"GetReportRequestLisByNextTokentResult"`
	ResponseMetadata        *core.ResponseMetadata
	Error                   *core.ResponseError
}

func (x *ReportRequestListByNextTokenResponse) GetReportListResult() *ReportRequestListResult {
	return x.ReportRequestListResult
}
func (x *ReportRequestListByNextTokenResponse) GetResponseMetadata() *core.ResponseMetadata {
	return x.ResponseMetadata
}
func (x *ReportRequestListByNextTokenResponse) GetError() *core.ResponseError {
	return x.Error
}

// #endregion

// #region Report

type AllOrdersReportResult struct {
	Entries []*AllOrdersReport
}
type ReturnReportResult struct {
	Entries []*ReturnReport
}
type AllListingsReportResult struct {
	Entries []*AllListingsReport
}
type FBAFeePreviewReportResult struct {
	Entries []*FBAFeePreviewReport
}
type InventoryHistoryReportResult struct {
	Entries []*InventoryHistoryReport
}
type ManageInventoryReportResult struct {
	Entries []*ManageInventoryReport
}

type AllOrdersReport struct {
	OrderID               string    `csv:"amazon-order-id"`
	SKU                   string    `csv:"sku"`
	OrderStatus           string    `csv:"order-status"`
	Currency              string    `csv:"currency"`
	ItemPrice             float32   `csv:"item-price"`
	Quantity              int32     `csv:"quantity"`
	ItemPromotionDiscount float32   `csv:"item-promotion-discount"`
	PurchaseDate          time.Time `csv:"purchase-date"`
}
type ReturnReport struct {
	SKU                 string    `csv:"sku"`
	OrderID             string    `csv:"order-id"`
	Quantity            int32     `csv:"quantity"`
	DetailedDisposition string    `csv:"detailed-disposition"`
	ReturnDate          time.Time `csv:"return-date"`
}
type AllListingsReport struct {
	SKU      string `csv:"seller-sku"`
	ItemName string `csv:"item-name"`
	ASIN     string `csv:"asin1"`
}
type FBAFeePreviewReport struct {
	SKU               string  `csv:"sku"`
	LongestSide       float32 `csv:"longest-side"`
	MedianSide        float32 `csv:"median-side"`
	ShortestSide      float32 `csv:"shortest-side"`
	DimensionUnitType string  `csv:"unit-of-dimension"`
	ItemPackageWeight float32 `csv:"item-package-weight"`
	WeightUnitType    string  `csv:"unit-of-weight"`
	ProductSizeTier   string  `csv:"product-size-tier"`
	FulfillmentFee    float32 `csv:"expected-fulfillment-fee-per-unit"`
}
type InventoryHistoryReport struct {
	Month  string  `csv:"month"`
	SKU    string  `csv:"sku"`
	AvgQty float32 `csv:"average-quantity"`
}
type ManageInventoryReport struct {
	SKU               string `csv:"sku"`
	AfnFulfillableQty string `csv:"afn-fulfillable-quantity"`
}

// #endregion

// #region Report Response

type AllOrdersReportResponse struct {
	Result *AllOrdersReportResult
	Error  *core.ResponseError
}

func (x *AllOrdersReportResponse) SetReportResult(v interface{}) {
	x.Result = v.(*AllOrdersReportResult)
}
func (x *AllOrdersReportResponse) GetReportResult() interface{} {
	return x.Result
}
func (x *AllOrdersReportResponse) GetError() *core.ResponseError {
	return x.Error
}

type ReturnReportResponse struct {
	Result *ReturnReportResult
	Error  *core.ResponseError
}

func (x *ReturnReportResponse) SetReportResult(v interface{}) {
	x.Result = v.(*ReturnReportResult)
}
func (x *ReturnReportResponse) GetReportResult() interface{} {
	return x.Result
}
func (x *ReturnReportResponse) GetError() *core.ResponseError {
	return x.Error
}

type AllListingsReportResponse struct {
	Result *AllListingsReportResult
	Error  *core.ResponseError
}

func (x *AllListingsReportResponse) SetReportResult(v interface{}) {
	x.Result = v.(*AllListingsReportResult)
}
func (x *AllListingsReportResponse) GetReportResult() interface{} {
	return x.Result
}
func (x *AllListingsReportResponse) GetError() *core.ResponseError {
	return x.Error
}

type FBAFeePreviewReportResponse struct {
	Result *FBAFeePreviewReportResult
	Error  *core.ResponseError
}

func (x *FBAFeePreviewReportResponse) SetReportResult(v interface{}) {
	x.Result = v.(*FBAFeePreviewReportResult)
}
func (x *FBAFeePreviewReportResponse) GetReportResult() interface{} {
	return x.Result
}
func (x *FBAFeePreviewReportResponse) GetError() *core.ResponseError {
	return x.Error
}

type InventoryHistoryReportResponse struct {
	Result *InventoryHistoryReportResult
	Error  *core.ResponseError
}

func (x *InventoryHistoryReportResponse) SetReportResult(v interface{}) {
	x.Result = v.(*InventoryHistoryReportResult)
}
func (x *InventoryHistoryReportResponse) GetReportResult() interface{} {
	return x.Result
}
func (x *InventoryHistoryReportResponse) GetError() *core.ResponseError {
	return x.Error
}

type ManageInventoryReportResponse struct {
	Result *ManageInventoryReportResult
	Error  *core.ResponseError
}

func (x *ManageInventoryReportResponse) SetReportResult(v interface{}) {
	x.Result = v.(*ManageInventoryReportResult)
}
func (x *ManageInventoryReportResponse) GetReportResult() interface{} {
	return x.Result
}
func (x *ManageInventoryReportResponse) GetError() *core.ResponseError {
	return x.Error
}

// #endregion
