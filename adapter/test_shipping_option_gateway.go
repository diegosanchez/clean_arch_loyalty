package adapter

import entity "github.com/diegosanchez/clean_arch_loyalty/entity"

type TestShippingOptionGateway struct {
}

func NewTestShippingOptionGateway() *TestShippingOptionGateway {
	return new(TestShippingOptionGateway)
}

func (this *TestShippingOptionGateway) ShippingForItem(item *entity.Item) *entity.Shipping {
	return entity.NewShipping(item)
}
