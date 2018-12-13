/*
Package model 用于模型层定义，所有db及cache对象封装均定义在这里。
只允许在这里添加对外暴露的接口
*/
package model

import (
	"camp/skel/model/skel"
	"camp/skel/model/user"
)

// NewSkel 构造Skel对象
func NewSkel() *skel.Skel {
	return &skel.Skel{}
}

func NewUser() *user.User{
	return &user.User{}
}


