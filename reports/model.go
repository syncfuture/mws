package reports

import (
	"time"

	"github.com/syncfuture/mws/core"
)

// #region Interface

type IReportListResponse interface {
	GetReportListResult() *ReportListResult
	GetResponseMetadata() *core.ResponseMetadata
	GetError() *core.ResponseError
}

type IReportResponse interface {
	SetReportResult(interface{})
	GetReportResult() interface{}
	GetError() *core.ResponseError
}

// #endregion

// #region ReportList Model

type ReportListResult struct {
	NextToken   string
	ReportInfos []*ReportInfo `xml:"ReportInfo"`
}
type ReportInfo struct {
	ReportType      string
	ReportId        string
	ReportRequestId string
	AvailableDate   time.Time
}

// #endregion

// #region Report Model

type AllOrderReportResult struct {
	Entries []*AllOrdersReport
}
type ReturnReportResult struct {
	Entries []*ReturnReport
}
type AllListingsReportResult struct {
	Entries []*AllListingsReport
}

type AllOrdersReport struct {
	SKU          string    `csv:"sku"`
	OrderStatus  string    `csv:"order-status"`
	PurchaseDate time.Time `csv:"purchase-date"`
}
type ReturnReport struct {
	SKU                 string `csv:"sku"`
	OrderID             string `csv:"order-id"`
	Quantity            int32  `csv:"quantity"`
	DetailedDisposition string `csv:"detailed-disposition"`
}
type AllListingsReport struct {
	ItemName string `csv:"item-name"`
	ASIN     string `csv:"asin1"`
}

// #endregion

// #region Report

type AllOrdersReportResponse struct {
	Result *AllOrderReportResult
	Error  *core.ResponseError
}

func (x *AllOrdersReportResponse) SetReportResult(v interface{}) {
	x.Result = v.(*AllOrderReportResult)
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

// #endregion

// #region Report List

type ReportListResponse struct {
	ReportListResult *ReportListResult `xml:"GetReportListResult"`
	ResponseMetadata *core.ResponseMetadata
	Error            *core.ResponseError
}

func (x *ReportListResponse) GetReportListResult() *ReportListResult {
	return x.ReportListResult
}
func (x *ReportListResponse) GetResponseMetadata() *core.ResponseMetadata {
	return x.ResponseMetadata
}
func (x *ReportListResponse) GetError() *core.ResponseError {
	return x.Error
}

type ReportListByNextTokenResponse struct {
	ReportListResult *ReportListResult `xml:"GetReportListByNextTokenResult"`
	ResponseMetadata *core.ResponseMetadata
	Error            *core.ResponseError
}

func (x *ReportListByNextTokenResponse) GetReportListResult() *ReportListResult {
	return x.ReportListResult
}
func (x *ReportListByNextTokenResponse) GetResponseMetadata() *core.ResponseMetadata {
	return x.ResponseMetadata
}
func (x *ReportListByNextTokenResponse) GetError() *core.ResponseError {
	return x.Error
}

// #endregion
