package entity

type Loyalty struct {
	user *User
}

func NewLoyalty(user *User) *Loyalty {
	result := new(Loyalty)

	result.user = user

	return result
}
