package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_item_id_pattern(t *testing.T) {
	itemId := NewItemId("MLA", 123)

	assert.Equal(t, "item_mla_123.json",
		itemId.RenderPattern("item_{site}_{nro}.json"), "itemid render - item_mla_123.json")
}
