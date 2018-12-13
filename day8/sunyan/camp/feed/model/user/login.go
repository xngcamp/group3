package user

import (
	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
)

func (user *User) Login() (userRet *User, err error) {
	c := user.GetC()
	defer c.Database.Session.Close()

	err = c.Find(bson.M{"email": user.Email, "password": user.Password}).One(&userRet)
	if err != nil {
		if err != mgo.ErrNotFound {
			return
		}
		err = nil
		return
	}

	return
}
