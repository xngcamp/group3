package user

import (
	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
)

func (token *Token) GetToken() (tokenRet *Token, err error) {
	c := token.GetC()
	defer c.Database.Session.Close()

	err = c.Find(bson.M{"userid": token.UserID, "token": token.Token}).One(&tokenRet)
	if err != nil {
		if err != mgo.ErrNotFound {
			return
		}
		err = nil
		return
	}
	return
}
