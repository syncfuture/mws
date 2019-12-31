package mws

import (
	"encoding/xml"
	"io/ioutil"
	"testing"

	"github.com/syncfuture/mws/finances"

	"github.com/stretchr/testify/assert"
	"github.com/subchen/go-xmldom"
)

func TestListFinancialEvents(t *testing.T) {
	r, err := _apiSet.Finances.ListFinancialEvents(&finances.ListFanancialEventsQuery{
		MaxResultsPerPage: "2",
		PostedAfter:       "2019-12-11T00:00:00",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, r)
	t.Log(r)

	doc := xmldom.Must(xmldom.ParseXML(r))
	token := doc.Root.GetChild("ListFinancialEventsResult").GetChild("NextToken").Text
	t.Log(token)
	if token != "" {
		r, err = _apiSet.Finances.ListFinancialEventsByNextToken(&finances.ListFinancialEventsByNextTokenQuery{
			NextToken: token,
		})
		assert.NoError(t, err)
		assert.NotEmpty(t, r)
	}
}

func TestUnmarshal(t *testing.T) {
	data, err := ioutil.ReadFile("finances/1.xml")
	assert.NoError(t, err)

	resp1 := new(finances.ListFinancialEventsResponse)
	err = xml.Unmarshal(data, resp1)
	assert.NoError(t, err)

	data, err = ioutil.ReadFile("finances/2.xml")
	assert.NoError(t, err)

	resp2 := new(finances.ListFinancialEventsByNextTokenResponse)
	err = xml.Unmarshal(data, resp2)
	assert.NoError(t, err)
}
