package Loyalty

import entity "github.com/diegosanchez/clean_arch_loyalty/entity"

type ShippingGateway interface {
	ShippingForItem(item *entity.Item) *entity.Shipping
}
