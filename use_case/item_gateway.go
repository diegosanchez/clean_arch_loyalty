package Loyalty

import "github.com/diegosanchez/clean_arch_loyalty/entity"

type ItemGateway interface {
	ItemById(itemId *entity.ItemId) *entity.Item
}
