package user

import (
	"camp/lib"
	"camp/skel/api"
	"camp/skel/service"
	"encoding/json"
	"github.com/simplejia/clog/api"
	"net/http"
)

type AddReq struct {
	Nick string `json:"Nick"`//用户昵称
	Sex int `json:"Sex"` //用户性别
	Email string `json:"Email"` //用户邮箱，做登录用，并且检测是否已经注册
	Password string `json:"Password"`//用户的密码
}


// Add just for demo
// @postfilter("Boss")
func (user *User) AddUser(w http.ResponseWriter, r *http.Request) {
	fun := "user.User.Add"

	var addReq *AddReq
	if err := json.Unmarshal(user.ReadBody(r), &addReq); err != nil  {
		clog.Error("%s param err: %v, req: %v", fun, err, addReq)
		user.ReplyFail(w, lib.CodePara)
		return
	}

	userApi := api.NewUser()
	userApi.Nick = addReq.Nick
	userApi.Sex = addReq.Sex
	userApi.Email = addReq.Email
	userApi.Password = addReq.Password

	if err := service.NewUser().AddUser(userApi); err != nil {
		clog.Error("%s skel.Add err: %v, req: %v", fun, err, addReq)
		user.ReplyFail(w, lib.CodeSrv)
		return
	}
	user.ReplyOk(w, userApi)

}