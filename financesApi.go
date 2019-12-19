package mws

type financesAPI struct {
	apiBase
}

func NewFinancesAPI(seller string) *financesAPI {
	r := new(financesAPI)
	r.Seller = seller
	r.Module = ConfigProvider.GetStringDefault("Orders.Module", "Finances")
	r.Version = ConfigProvider.GetStringDefault("Orders.Version", "2015-05-01")
	return r
}

type ListFanancialEventsQuery struct {
	queryBase
	MaxResultsPerPage     string
	AmazonOrderId         string
	FinancialEventGroupId string
	PostedAfter           string
	PostedBefore          string
}

func (x *financesAPI) ListFinancialEvents(query *ListFanancialEventsQuery) (string, error) {
	client := x.newClient("ListFinancialEvents")

	client.setParameter("MaxResultsPerPage", query.MaxResultsPerPage)
	client.setParameter("PostedAfter", query.PostedAfter)

	return client.Get()
}

type ListFinancialEventsByNextTokenQuery struct {
	queryBase
	NextToken string
}

func (x *financesAPI) ListFinancialEventsByNextToken(query *ListFinancialEventsByNextTokenQuery) (string, error) {
	client := x.newClient("ListFinancialEventsByNextToken")

	client.setParameter("NextToken", query.NextToken)

	return client.Get()
}
