package Loyalty

import entity "github.com/diegosanchez/clean_arch_loyalty/entity"

type LoyaltyDataMapper interface {
	LoyaltyFor(user *entity.User) *entity.Loyalty
}
