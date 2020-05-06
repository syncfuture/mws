package reports

import (
	"time"

	"github.com/syncfuture/mws/core"
)

type IReportListResponse interface {
	GetReportListResult() *ReportListResult
	GetReportResult() *ReportResult
	GetResponseMetadata() *core.ResponseMetadata
	GetError() *core.ResponseError
}

// #region Report List

type ReportListResponse struct {
	ReportListResult *ReportListResult `xml:"GetReportListResult"`
	ReportResult     *ReportResult
	ResponseMetadata *core.ResponseMetadata
	Error            *core.ResponseError
}

func (x *ReportListResponse) GetReportListResult() *ReportListResult {
	return x.ReportListResult
}
func (x *ReportListResponse) GetReportResult() *ReportResult {
	return x.ReportResult
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
	ReportResult     *ReportResult
	ResponseMetadata *core.ResponseMetadata
	Error            *core.ResponseError
}

func (x *ReportListByNextTokenResponse) GetReportListResult() *ReportListResult {
	return x.ReportListResult
}
func (x *ReportListByNextTokenResponse) GetReportResult() *ReportResult {
	return x.ReportResult
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

type ReportResult struct {
	AllOrdersReport   []*AllOrdersReport
	ReturnReport      []*ReturnReport
	AllListingsReport []*AllListingsReport
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
