package tests

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/syncfuture/mws/finances"

	"github.com/stretchr/testify/assert"
	"github.com/subchen/go-xmldom"
	u "github.com/syncfuture/go/util"
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
	filepath.Walk("finances/test", func(file string, fi os.FileInfo, err error) error {
		if u.LogError(err) {
			return err
		}
		if !fi.IsDir() {
			data, err := ioutil.ReadFile(file)
			assert.NoError(t, err)

			if strings.HasSuffix(file, "00001.xml") {
				resp := new(finances.ListFinancialEventsResponse)
				err = xml.Unmarshal(data, resp)
				assert.NoError(t, err)
			} else {
				resp := new(finances.ListFinancialEventsByNextTokenResponse)
				err = xml.Unmarshal(data, resp)
				assert.NoError(t, err)
			}
		}

		return nil
	})

	// for i, file := range files {
	// 	data, err := ioutil.ReadFile(file.Name())
	// 	assert.NoError(t, err)

	// 	if i == 0 {
	// 		resp := new(finances.ListFinancialEventsResponse)
	// 		err = xml.Unmarshal(data, resp)
	// 		assert.NoError(t, err)
	// 	} else {
	// 		resp := new(finances.ListFinancialEventsByNextTokenResponse)
	// 		err = xml.Unmarshal(data, resp)
	// 		assert.NoError(t, err)
	// 	}
	// }
}
