package tests

import (
	"testing"

	"github.com/syncfuture/mws/products"

	"github.com/stretchr/testify/assert"
)

func TestGetMatchingProduct(t *testing.T) {
	xml, err := _apiSet.Products.GetMatchingProduct(&products.GetMatchingProductQuery{
		ASINList: []string{"B07S58DSMH"},
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, xml)
	t.Log(xml)
}
