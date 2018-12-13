package user

import (
	"camp/feed/api"
	"camp/feed/controller/util"
	"camp/feed/service"
	"camp/lib"
	"encoding/json"
	"github.com/simplejia/clog/api"
	"net/http"
)

type UpdateReq struct {
	Id       string `json:"id" bson:"_id"`
	Nick     string `json:"nick"`
	Sex      byte   `json:"sex"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Email    string `json:"email"`
}

func (updateReq *UpdateReq) Regular() (ok bool) {
	if updateReq == nil {
		return
	}

	if "" == updateReq.Id || "" == updateReq.Token ||
		(updateReq.Sex != 0 && updateReq.Sex != 1 && updateReq.Sex != 2) {
		return
	}

	ok = true
	return
}

type UpdateResp struct{}

// @postfilter("Boss")
func (user *User) Update(w http.ResponseWriter, r *http.Request) {
	fun := "user.User.Update"

	var updateReq *UpdateReq
	if err := json.Unmarshal(user.ReadBody(r), &updateReq); err != nil || !updateReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, updateReq)
		user.ReplyFail(w, lib.CodePara)
		return
	}
	//验证token
	b := util.ValidateToken(updateReq.Id, updateReq.Token)
	if b == false {
		detail := "token无效"
		user.ReplyFailWithDetail(w, lib.CodePara, detail)
		return
	}

	results, err := service.NewUser().GetById(updateReq.Id)
	if err != nil {
		clog.Error("%s user.login err: %v, req: %v", fun, err, updateReq)
		user.ReplyFail(w, lib.CodeSrv)
		return
	}
	if results == nil {
		detail := "there is not this user"
		user.ReplyFailWithDetail(w, lib.CodePara, detail)
		return
	}
	userApi := api.NewUser()
	userApi.ID = updateReq.Id
	userApi.Sex = updateReq.Sex
	userApi.Nick = updateReq.Nick
	userApi.Password = util.Encryption(updateReq.Password)
	userApi.Email = updateReq.Email
	if err := service.NewUser().Update(userApi); err != nil {
		clog.Error("%s feed.Update err: %v, req: %v", fun, err, updateReq)
		user.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &UpdateResp{}
	user.ReplyOk(w, resp)

	return
}
