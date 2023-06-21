package user

func (u *User) Auth(username string, password string) (bool, error) {
	return true, nil
}
