package Loyalty

type LoyaltyFreeShipping struct {
	itemGateway ItemGateway
	shippingOpt ShippingGateway
	loyaltyDM   LoyaltyDataMapper
}

func NewLoyaltyFreeShipping(itemGW ItemGateway,
	shippingOptGW ShippingGateway,
	loyaltyDM LoyaltyDataMapper) *LoyaltyFreeShipping {

	result := new(LoyaltyFreeShipping)

	result.itemGateway = itemGW
	result.shippingOpt = shippingOptGW
	result.loyaltyDM = loyaltyDM

	return result
}

func (this *LoyaltyFreeShipping) doWork(req, res interface{}) {
}
