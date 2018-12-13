package user

import (
	"camp/feed/api"
	"camp/feed/controller/util"
	"camp/feed/service"
	"camp/lib"
	"encoding/json"
	"github.com/satori/go.uuid"
	"github.com/simplejia/clog/api"
	"net/http"
	"time"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (loginReq *LoginReq) Regular() (ok bool) {
	if loginReq == nil {
		return
	}

	if "" == loginReq.Email || "" == loginReq.Password {
		return
	}

	ok = true
	return
}

type LoginResp struct {
	Token string `json:"token"`
}

func (user *User) Login(w http.ResponseWriter, r *http.Request) {
	fun := "feed.User.Login"

	var loginReq LoginReq
	if err := json.Unmarshal(user.ReadBody(r), &loginReq); err != nil || !loginReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, loginReq)
		user.ReplyFail(w, lib.CodePara)
		return
	}
	userApi := api.NewUser()
	userApi.Email = loginReq.Email
	userApi.Password = util.Encryption(loginReq.Password)
	results, err := service.NewUser().Login(userApi.Email, userApi.Password)
	if err != nil {
		clog.Error("%s user.login err: %v, req: %v", fun, err, loginReq)
		user.ReplyFail(w, lib.CodeSrv)
		return
	}
	if results == nil {
		detail := "Incorrect user name or password"
		user.ReplyFailWithDetail(w, lib.CodePara, detail)
		return
	}

	//生成token
	tokenApi := api.NewToken()
	tokenApi.UserID = results.ID
	tokenApi.Token = uuid.Must(uuid.NewV4()).String()
	tokenApi.OutTime = int(time.Now().Unix())
	if err := service.NewToken().AddToken(tokenApi); err != nil {
		clog.Error("%s user.register err: %v, req: %v", fun, err, loginReq)
		user.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &LoginResp{tokenApi.Token}
	user.ReplyOk(w, resp)

	return
}
