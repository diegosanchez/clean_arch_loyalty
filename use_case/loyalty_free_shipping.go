package Loyalty

import "github.com/diegosanchez/clean_arch_loyalty/entity"

type LoyaltyFreeShipping struct {
	itemGW        ItemGateway
	shippingOptGW ShippingGateway
	loyaltyDM     LoyaltyDataMapper
}

func NewLoyaltyFreeShipping(itemGW ItemGateway,
	shippingOptGW ShippingGateway,
	loyaltyDM LoyaltyDataMapper) *LoyaltyFreeShipping {

	result := new(LoyaltyFreeShipping)

	result.itemGW = itemGW
	result.shippingOptGW = shippingOptGW
	result.loyaltyDM = loyaltyDM

	return result
}

func (this *LoyaltyFreeShipping) doWork(req, res map[string]interface{}) {

	item := this.itemGW.ItemById(
		req["item_id"].(*entity.ItemId))

	// User should be retrieved through GW
	// loyalty := this.loyaltyDM.LoyaltyFor(
	// 	entity.NewUser(req["user_id"].(*entity.UserId)))

	shippingOpt := this.shippingOptGW.ForItem(item)

	dataBag := shippingOpt.AsMapForReview(entity.NewDatabag())

	dataBag.AsMap(&res)

}
