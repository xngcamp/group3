package user

import (
	"camp/feed/api"
	"camp/feed/mongo"
	mgo "github.com/globalsign/mgo"
)

type Token api.Token

// Db 返回db name
func (token *Token) Db() (db string) {
	return "feed"
}

// Table 返回table name
func (token *Token) Table() (table string) {
	return "token"
}

// GetC 返回db col
func (token *Token) GetC() (c *mgo.Collection) {
	db, table := token.Db(), token.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}
