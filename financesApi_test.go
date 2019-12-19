package mws

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/subchen/go-xmldom"
)

func TestListFinancialEvents(t *testing.T) {
	client := NewFinancesAPI("LF")
	r, err := client.ListFinancialEvents(&ListFanancialEventsQuery{
		MaxResultsPerPage: "1",
		PostedAfter:       "2019-12-11T00:00:00",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, r)
	t.Log(r)

	doc := xmldom.Must(xmldom.ParseXML(r))
	token := doc.Root.GetChild("ListFinancialEventsResult").GetChild("NextToken").Text
	t.Log(token)
	if token != "" {
		r, err = client.ListFinancialEventsByNextToken(&ListFinancialEventsByNextTokenQuery{
			NextToken: token,
		})
		assert.NoError(t, err)
		assert.NotEmpty(t, r)
	}
}
