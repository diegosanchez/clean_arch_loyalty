package adapter

import entity "github.com/diegosanchez/clean_arch_loyalty/entity"

type TestItemGateway struct {
}

func NewTestItemGateway() *TestItemGateway {
	return new(TestItemGateway)
}

func (this *TestItemGateway) ItemById(itemId *entity.ItemId) *entity.Item {
	return entity.NewItem(itemId)
}
