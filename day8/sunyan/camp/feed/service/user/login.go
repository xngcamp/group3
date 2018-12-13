package user

import (
	"camp/feed/api"
	"camp/feed/model"
)

func (user *User) Login(email string, pwd string) (userApi *api.User, err error) {
	userModel := model.NewUser()
	userModel.Email = email
	userModel.Password = pwd
	if userModel, err = userModel.Login(); err != nil {
		return
	}
	if userModel == nil {
		return
	}
	userApi = (*api.User)(userModel)
	return
}
