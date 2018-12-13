package user

import (
	"camp/feed/api"
	"camp/feed/model"
)

func (token *Token) GetToken(userId string, tok string) (tokenApi *api.Token, err error) {
	tokenModel := model.NewToken()
	tokenModel.UserID = userId
	tokenModel.Token = tok
	if tokenModel, err = tokenModel.GetToken(); err != nil {
		return
	}
	if tokenModel == nil {
		return
	}
	tokenApi = (*api.Token)(tokenModel)
	return
}
