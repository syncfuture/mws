package finances

import (
	"time"

	"github.com/syncfuture/mws/core"
)

type Amount struct {
	CurrencyAmount float32
	CurrencyCode   string
}

type ListFinancialEventsResponse struct {
	ListFinancialEventsResult *ListFinancialEventsResult `xml:"ListFinancialEventsResult"`
	ResponseMetadata          *core.ResponseMetadata
}

type ListFinancialEventsByNextTokenResponse struct {
	ListFinancialEventsResult *ListFinancialEventsResult `xml:"ListFinancialEventsByNextTokenResult"`
	ResponseMetadata          *core.ResponseMetadata
}

type ListFinancialEventsResult struct {
	NextToken       string
	FinancialEvents *FinancialEvents
}
type FinancialEvents struct {
	ShipmentEventList *ShipmentEventList
}

type ShipmentEventList struct {
	ShipmentEvents []*ShipmentEvent `xml:"ShipmentEvent"`
}

type ShipmentEvent struct {
	ShipmentItemList *ShipmentItemList
	AmazonOrderId    string
	SellerOrderId    string
	MarketplaceName  string
	PostedDate       time.Time
}

type ShipmentItemList struct {
	ShipmentItems []*ShipmentItem `xml:"ShipmentItem"`
}

type ShipmentItem struct {
	OrderItemId         string
	SellerSKU           string
	QuantityShipped     int
	ItemTaxWithheldList *ItemTaxWithheldList
	ItemChargeList      *ItemChargeList
	ItemFeeList         *ItemFeeList
	PromotionList       *PromotionList
}

type ItemTaxWithheldList struct {
	ItemTaxWithhelds []*TaxWithheldComponent `xml:"TaxWithheldComponent"`
}

type ItemChargeList struct {
	Charges []*ChargeComponent `xml:"ChargeComponent"`
}

type ItemFeeList struct {
	Fees []*FeeComponent `xml:"FeeComponent"`
}

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

type ChargeComponent struct {
	ChargeType   string
	ChargeAmount *Amount
}

type FeeComponent struct {
	FeeType   string
	FeeAmount *Amount
}

type Promotion struct {
	PromotionId     string
	PromotionType   string
	PromotionAmount *Amount
}

type RefundEventList struct {
}

type RetrochargeEventList struct {
}

type ProductAdsPaymentEventList struct {
}

type ServiceFeeEventList struct {
}

type AdjustmentEventList struct {
}

type CouponPaymentEventList struct {
}

type SellerReviewEnrollmentPaymentEventList struct {
}

type SellerDealPaymentEventList struct {
}
