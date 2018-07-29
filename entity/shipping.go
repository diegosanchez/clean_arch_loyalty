package entity

type Shipping struct {
	item *Item
}

func NewShipping(item *Item) *Shipping {
	result := new(Shipping)

	result.item = item

	return result
}
