package finances

import (
	"time"

	"github.com/syncfuture/mws/core"
)

type Amount struct {
	CurrencyAmount float32
	CurrencyCode   string
}

type ChargeComponent struct {
	ChargeType   string
	ChargeAmount *Amount
}

type FeeComponent struct {
	FeeType   string
	FeeAmount *Amount
}

type FeeList struct {
	Fees *FeeComponent `xml:"FeeComponent"`
}

type IListFinancialEventsResponse interface {
	GetListFinancialEventsResult() *ListFinancialEventsResult
	GetResponseMetadata() *core.ResponseMetadata
	GetError() *core.ResponseError
}

type ListFinancialEventsResponse struct {
	ListFinancialEventsResult *ListFinancialEventsResult `xml:"ListFinancialEventsResult"`
	ResponseMetadata          *core.ResponseMetadata
	Error                     *core.ResponseError
}

func (x *ListFinancialEventsResponse) GetListFinancialEventsResult() *ListFinancialEventsResult {
	return x.ListFinancialEventsResult
}
func (x *ListFinancialEventsResponse) GetResponseMetadata() *core.ResponseMetadata {
	return x.ResponseMetadata
}
func (x *ListFinancialEventsResponse) GetError() *core.ResponseError {
	return x.Error
}

type ListFinancialEventsByNextTokenResponse struct {
	ListFinancialEventsResult *ListFinancialEventsResult `xml:"ListFinancialEventsByNextTokenResult"`
	ResponseMetadata          *core.ResponseMetadata
	Error                     *core.ResponseError
}

func (x *ListFinancialEventsByNextTokenResponse) GetListFinancialEventsResult() *ListFinancialEventsResult {
	return x.ListFinancialEventsResult
}
func (x *ListFinancialEventsByNextTokenResponse) GetResponseMetadata() *core.ResponseMetadata {
	return x.ResponseMetadata
}
func (x *ListFinancialEventsByNextTokenResponse) GetError() *core.ResponseError {
	return x.Error
}

type ListFinancialEventsResult struct {
	NextToken       string
	FinancialEvents *FinancialEvents
}
type FinancialEvents struct {
	ShipmentEventList                      *ShipmentEventList
	RefundEventList                        *RefundEventList
	RetrochargeEventList                   *RetrochargeEventList
	ProductAdsPaymentEventList             *ProductAdsPaymentEventList
	ServiceFeeEventList                    *ServiceFeeEventList
	AdjustmentEventList                    *AdjustmentEventList
	CouponPaymentEventList                 *CouponPaymentEventList
	SellerReviewEnrollmentPaymentEventList *SellerReviewEnrollmentPaymentEventList
	SellerDealPaymentEventList             *SellerDealPaymentEventList
}

// ShipmentEventList *************************************************************************************
type ShipmentEventList struct {
	ShipmentEvents []*ShipmentEvent `xml:"ShipmentEvent"`
}

type ShipmentEvent struct {
	ShipmentItemList           *ShipmentItemList
	ShipmentItemAdjustmentList *ShipmentItemList
	AmazonOrderId              string
	SellerOrderId              string
	MarketplaceName            string
	PostedDate                 time.Time
}

type ShipmentItemList struct {
	ShipmentItems []*ShipmentItem `xml:"ShipmentItem"`
}

type ShipmentItem struct {
	OrderItemId              string
	OrderAdjustmentItemId    string
	SellerSKU                string
	QuantityShipped          int
	ItemTaxWithheldList      *ItemTaxWithheldList
	ItemChargeList           *ItemChargeList
	ItemChargeAdjustmentList *ItemChargeList
	ItemFeeList              *FeeList
	ItemFeeAdjustmentList    *FeeList
	PromotionList            *PromotionList
}

type ItemTaxWithheldList struct {
	ItemTaxWithhelds []*TaxWithheldComponent `xml:"TaxWithheldComponent"`
}

type ItemChargeList struct {
	Charges []*ChargeComponent `xml:"ChargeComponent"`
}

// type ItemFeeList struct {
// 	Fees []*FeeComponent `xml:"FeeComponent"`
// }

type PromotionList struct {
	Promotions []*Promotion `xml:"Promotion"`
}

type TaxWithheldComponent struct {
	TaxCollectionModel string
	TaxesWithheld      *TaxesWithheld
}

type TaxesWithheld struct {
	Taxes []*ChargeComponent `xml:"ChargeComponent"`
}

type Promotion struct {
	PromotionId     string
	PromotionType   string
	PromotionAmount *Amount
}

// RefundEventList *************************************************************************************
type RefundEventList struct {
	ShipmentEvents []*ShipmentEvent `xml:"ShipmentEvent"`
}

// RetrochargeEventList *************************************************************************************
type RetrochargeEventList struct {
}

// ProductAdsPaymentEventList *************************************************************************************
type ProductAdsPaymentEventList struct {
	ProductAdsPaymentEvents []*ProductAdsPaymentEvent `xml:"ProductAdsPaymentEvent"`
}
type ProductAdsPaymentEvent struct {
	InvoiceId        string    `xml:"invoiceId"`
	TransactionType  string    `xml:"transactionType"`
	BaseValue        *Amount   `xml:"baseValue"`
	TaxValue         *Amount   `xml:"taxValue"`
	TransactionValue *Amount   `xml:"transactionValue"`
	PostedDate       time.Time `xml:"postedDate"`
}

// ServiceFeeEventList *************************************************************************************
type ServiceFeeEventList struct {
	ServiceFeeEvents *ServiceFeeEvent `xml:"ServiceFeeEvent"`
}
type ServiceFeeEvent struct {
	AmazonOrderId string
	FeeList       *FeeList
}

// AdjustmentEventList *************************************************************************************
type AdjustmentEventList struct {
	AdjustmentEvents []*AdjustmentEvent `xml:"AdjustmentEvent"`
}
type AdjustmentEvent struct {
	AdjustmentType     string
	PostedDate         time.Time
	AdjustmentAmount   *Amount
	AdjustmentItemList *AdjustmentItemList
}
type AdjustmentItemList struct {
	Items []*AdjustmentItem `xml:"AdjustmentItem"`
}
type AdjustmentItem struct {
	SellerSKU          string
	ProductDescription string
	Quantity           int
	PerUnitAmount      *Amount
	TotalAmount        *Amount
}

// CouponPaymentEventList *************************************************************************************
type CouponPaymentEventList struct {
	CouponPaymentEvents []*CouponPaymentEvent `xml:"CouponPaymentEvent"`
}
type CouponPaymentEvent struct {
	PaymentEventId          string
	SellerCouponDescription string
	CouponId                string
	ClipOrRedemptionCount   int
	PostedDate              time.Time
	TotalAmount             *Amount
	Fee                     *FeeComponent    `xml:"FeeComponent"`
	Charge                  *ChargeComponent `xml:"ChargeComponent"`
}

// SellerReviewEnrollmentPaymentEventList *************************************************************************************
type SellerReviewEnrollmentPaymentEventList struct {
	SellerReviewEnrollmentPaymentEvents []*SellerReviewEnrollmentPaymentEvent `xml:"SellerReviewEnrollmentPaymentEvent"`
}
type SellerReviewEnrollmentPaymentEvent struct {
	ParentASIN   string
	EnrollmentId string
	PostedDate   time.Time
	Fee          *FeeComponent    `xml:"FeeComponent"`
	Charge       *ChargeComponent `xml:"ChargeComponent"`
}

// SellerDealPaymentEventList *************************************************************************************
type SellerDealPaymentEventList struct {
	SellerDealPaymentEvents []*SellerDealPaymentEvent `xml:"SellerDealPaymentEvent"`
}
type SellerDealPaymentEvent struct {
	DealId          string    `xml:"dealId"`
	EventType       string    `xml:"eventType"`
	FeeType         string    `xml:"feeType"`
	DealDescription string    `xml:"dealDescription"`
	TotalAmount     *Amount   `xml:"totalAmount"`
	FeeAmount       *Amount   `xml:"feeAmount"`
	TaxAmount       *Amount   `xml:"taxAmount"`
	PostedDate      time.Time `xml:"postedDate"`
}
