package mws

type financesApi struct {
	api
}

type ListFanancialEventsQuery struct {
	MaxResultsPerPage     string
	AmazonOrderId         string
	FinancialEventGroupId string
	PostedAfter           string
	PostedBefore          string
}

func NewFinancesApi() *financesApi {
	r := new(financesApi)
	r.Module = ConfigProvider.GetStringDefault("Orders.Module", "Finances")
	r.Version = ConfigProvider.GetStringDefault("Orders.Version", "2015-05-01")
	return r
}

func (x *financesApi) ListFinancialEvents(query *ListFanancialEventsQuery) (string, error) {
	api := newAPI(x.Module, x.Version, "ListFinancialEvents")

	api.setParameter("MaxResultsPerPage", query.MaxResultsPerPage)
	api.setParameter("PostedAfter", query.PostedAfter)

	return api.Get()
}

func (x *financesApi) ListFinancialEventsByNextToken(nextToken string) (string, error) {
	api := newAPI(x.Module, x.Version, "ListFinancialEventsByNextToken")

	api.setParameter("NexToken", nextToken)

	return api.Get()
}
