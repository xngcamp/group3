/*
Package service 用于定义服务层代码。
只允许在这里添加对外暴露的接口
*/
package service

import (
	"camp/skel/service/skel"
	"camp/skel/service/user"
)
func NewSkel() *skel.Skel {
	return &skel.Skel{}
}
func NewUser() *user.User{
	return &user.User{}
}