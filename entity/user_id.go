package entity

type UserId struct {
	id int
}

func NewUserId(id int) *UserId {
	result := new(UserId)

	result.id = id

	return result
}
