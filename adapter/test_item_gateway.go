package adapter

import entity "github.com/diegosanchez/clean_arch_loyalty/entity"

type TestItemGateway struct {
}

func NewTestItemGateway() *TestItemGateway {
	return new(TestItemGateway)
}

func (this *TestItemGateway) ItemById(itemId *entity.ItemId) *entity.Item {
	price := 0
	switch *itemId {
	case *entity.NewItemId("MLA", 12):
		price = 1400
	case *entity.NewItemId("MLA", 50):
		price = 200
	default:
		panic("No item found!")
	}

	return entity.NewItem(itemId, price)
}
