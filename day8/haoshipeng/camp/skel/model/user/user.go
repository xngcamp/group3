/*
Package skel just for demo
*/
package user

import (
	mgo "github.com/globalsign/mgo"
	"camp/skel/api"
	"camp/skel/mongo"
)

// Skel 定义db对应的类型
type User api.User

// Db 返回db name
func (skel *User) Db() (db string) {
	return "skel"
}

// Table 返回table name
func (skel *User) Table() (table string) {
	return "skel"
}

// GetC 返回db col
func (skel *User) GetC() (c *mgo.Collection) {
	db, table := skel.Db(), skel.Table()
	session := mongo.DBS[db]
	sessionCopy := session.Copy()
	c = sessionCopy.DB(db).C(table)
	return
}
