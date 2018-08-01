package Loyalty

import entity "github.com/diegosanchez/clean_arch_loyalty/entity"

type ShippingGateway interface {
	ForItem(item *entity.Item) entity.Shipping
}
