package adapter

import entity "github.com/diegosanchez/clean_arch_loyalty/entity"

type TestLoyaltyDM struct{}

func NewTestLoyaltyDM() *TestLoyaltyDM {
	return new(TestLoyaltyDM)
}

func (this *TestLoyaltyDM) LoyaltyFor(user *entity.User) *entity.Loyalty {
	return entity.NewLoyalty(user)
}
