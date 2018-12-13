package feed

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"net/http"

	"camp/feed/api"
	"camp/feed/service"
	"camp/lib"

	"github.com/simplejia/clog/api"
)

// AddReq 定义输入
type AddReq struct {
	Txt string `json:"txt"`
}

// Regular 用于参数校验
func (addReq *AddReq) Regular() (ok bool) {
	if addReq == nil {
		return
	}

	if "" == addReq.Txt {
		return
	}

	ok = true
	return
}

// AddResp 定义输出
type AddResp struct {
	data map[string]interface{}
}

// Add just for demo
// @postfilter("Boss")
func (feed *Feed) Add(w http.ResponseWriter, r *http.Request) {
	fun := "feed.Feed.Add"

	var addReq *AddReq
	if err := json.Unmarshal(feed.ReadBody(r), &addReq); err != nil || !addReq.Regular() {
		clog.Error("%s param err: %v, req: %v", fun, err, addReq)
		feed.ReplyFail(w, lib.CodePara)
		return
	}
	feedApi := api.NewFeed()
	id := uuid.Must(uuid.NewV4())
	feedApi.ID = id.String()
	feedApi.Txt = addReq.Txt
	if err := service.NewFeed().Add(feedApi); err != nil {
		clog.Error("%s skel.Add err: %v, req: %v", fun, err, addReq)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &AddResp{map[string]interface{}{"data": feedApi}}
	feed.ReplyOk(w, resp)

	// 进行一些异步处理的工作
	go lib.Updates(feedApi, lib.ADD, nil)

	return
}
