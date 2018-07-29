package entity

type Item struct {
	id *ItemId
}

func NewItem(id *ItemId) *Item {
	result := new(Item)

	result.id = id

	return result
}
