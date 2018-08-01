package Loyalty

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

import "github.com/diegosanchez/clean_arch_loyalty/entity"

func Test_dummy(t *testing.T) {
	req := make(map[string]interface{})
	req["item_id"] = entity.NewItemId("MLA", 12)

	assert.NotNil(t, req["item_id"].(*entity.ItemId))
}
