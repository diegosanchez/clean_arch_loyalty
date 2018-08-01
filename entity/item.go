package entity

type Item struct {
	price int64
	id    *ItemId
}

func NewItem(id *ItemId, price int64) *Item {
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
