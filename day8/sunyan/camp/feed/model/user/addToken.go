package user

func (token *Token) AddToken() (err error) {
	c := token.GetC()
	defer c.Database.Session.Close()

	err = c.Insert(token)
	if err != nil {
		return
	}

	return
}
