package domain

type User struct {
	ID int
	FullName string
	Email string
    Password string
}

func (u *User) GetUserID() int {
    return u.ID
}
