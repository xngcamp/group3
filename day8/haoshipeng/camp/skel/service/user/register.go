package user
import (
	"camp/skel/api"
	"camp/skel/model"
)

// Add 定义新增操作
func (user *User) AddUser(userApi *api.User) (err error) {
	userModel := model.NewUser()
	userModel.Nick= userApi.Nick
	userModel.Sex= userApi.Sex
	userModel.Email= userApi.Email
	userModel.Password= userApi.Password

	if err = userModel.Add(); err != nil {
		return
	}

	return
}
