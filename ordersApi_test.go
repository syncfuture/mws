package mws

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListOrders(t *testing.T) {
	api := NewOrdersApi()
	xml, err := api.ListOrders(&ListOrdersQuery{
		CreatedAfter: time.Now().UTC().Format(time.RFC3339),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, xml)
}