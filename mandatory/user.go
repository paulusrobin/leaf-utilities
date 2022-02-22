package leafMandatory

type User struct {
	login bool
	id    uint64
	email string
}

func (u User) ID() uint64 {
	return u.id
}

func (u User) Email() string {
	return u.email
}

func (u User) IsLogin() bool {
	return u.login
}

func (u User) JSON() map[string]interface{} {
	return map[string]interface{}{
		"is_login": u.IsLogin(),
		"id":       u.ID(),
		"email":    u.Email(),
	}
}
