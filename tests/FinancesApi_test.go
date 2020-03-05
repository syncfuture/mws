package tests

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/syncfuture/armos/go/report/xmlutil"
	"github.com/syncfuture/mws/finances"

	"github.com/stretchr/testify/assert"
	"github.com/subchen/go-xmldom"
	"github.com/syncfuture/go/u"
)

func TestListFinancialEvents(t *testing.T) {
	r, err := _apiSet.Finances.ListFinancialEvents(&finances.ListFanancialEventsQuery{
		MaxResultsPerPage: "100",
		PostedAfter:       "2020-01-08T06:19:00Z",
		PostedBefore:      "2020-01-08T06:22:00Z",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, r)

	doc := xmldom.Must(xmldom.ParseXML(r))
	token := xmlutil.GetNodeText(doc.Root, "ListFinancialEventsResult/NextToken")
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
