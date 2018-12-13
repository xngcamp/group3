package user

import (
	"camp/feed/api"
	"camp/feed/model"
)

func (token *Token) AddToken(tokenApi *api.Token) (err error) {
	tokenModel := model.NewToken()
	tokenModel.UserID = tokenApi.UserID
	tokenModel.Token = tokenApi.Token
	tokenModel.OutTime = tokenApi.OutTime
	if err = tokenModel.AddToken(); err != nil {
		return
	}

	return
}
