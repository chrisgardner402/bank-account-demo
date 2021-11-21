package entity

type User struct {
	userid   string
	username string
}

func (u *User) Userid() string {
	return u.userid
}

func (u *User) Username() string {
	return u.username
}
