package reports

import (
	"time"

	"github.com/syncfuture/mws/core"
)

type IReportListResponse interface {
	GetReportListResult() *ReportListResult
	GetResponseMetadata() *core.ResponseMetadata
	GetError() *core.ResponseError
}

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

// #endregion

// #region Report List By NextToken

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

// #region Result Model

type ReportListResult struct {
	NextToken  string
	ReportInfo []*ReportInfo `xml:"ReportInfo"`
}
type ReportInfo struct {
	ReportType      string
	ReportId        string
	ReportRequestId string
	AvailableDate   time.Time
}
type Reports struct {
	AllOrdersReport   *AllOrdersReport
	ReturnReport      *ReturnReport
	AllListingsReport *AllListingsReport
}
type AllOrdersReport struct {
	SKU          string    `xml:"sku"`
	OrderStatus  string    `xml:"order-status"`
	PurchaseDate time.Time `xml:"purchase-date"`
}
type ReturnReport struct {
	SKU                 string `xml:"sku"`
	OrderID             string `xml:"order-id"`
	Quantity            int32  `xml:"quantity"`
	DetailedDisposition string `xml:"detailed-disposition"`
}
type AllListingsReport struct {
	ItemName string `xml:"item-name"`
	ASIN     string `xml:"asin1"`
}

// #endregion
