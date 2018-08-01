package entity

type User struct {
	id *UserId
}

func NewUser(id *UserId) *User {
	result := new(User)

	result.id = id

	return result
}
