package entity

type Item struct {
	price int
	id    *ItemId
}

func NewItem(id *ItemId, price int) *Item {
	result := new(Item)

	result.id = id
	result.price = price

	return result
}

func (this *Item) ShippingBasedOnPrice(shipping Shipping) Shipping {
	if this.price < 1200 {
		return shipping
	}
	return NewFreeShipping(this)
}
