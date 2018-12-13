package util

import (
	"camp/feed/service"
	"fmt"
	"time"
)

//验证token
func ValidateToken(userId string, token string) bool {
	results, err := service.NewToken().GetToken(userId, token)
	if err != nil {
		return false
	}
	if results == nil {
		return false
	}

	nowTime := int(time.Now().Unix())
	fmt.Println(nowTime - results.OutTime)
	if nowTime-results.OutTime > 1200 {
		return false
	}
	return true
}
