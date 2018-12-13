package user

import (
	"camp/skel/service"
	"camp/lib"
	"encoding/json"
	"github.com/simplejia/clog/api"
	"net/http"
)

type LoginReq struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token"`
}


// @postfilter("Boss")
func (user *User) Login(w http.ResponseWriter, r *http.Request)  {
	fun := "user.User.Login"

	var loginReq *LoginReq

	if err := json.Unmarshal(user.ReadBody(r), &loginReq) ; err !=nil{
		clog.Error("%s param err: %v, req: %v", fun, err, loginReq)
		user.ReplyFail(w, lib.CodePara)
		return
	}

	userApi,err := service.NewUser().GetUserInfo(loginReq.Email)
	if err != nil {
		user.ReplyFail(w,lib.CodeFound)
		return
	}
	if userApi.Password != loginReq.Password {
		http.Error(w,"密码错误",404)
	}


	resp := &LoginResp{userApi.Email}
	user.ReplyOk(w,resp)

}