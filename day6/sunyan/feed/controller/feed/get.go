package feed

import (
	"net/http"

	"camp/feed/api"
	"camp/feed/service"
	"camp/lib"

	"github.com/simplejia/clog/api"
)

// GetResp 定义输出
type GetResp struct {
	Feed []*api.Feed `json:"feed,omitempty"`
}

// Get just for demo
// @postfilter("Boss")
func (feed *Feed) Get(w http.ResponseWriter, r *http.Request) {
	fun := "feed.Feed.Get"

	feedApi, err := service.NewFeed().GetAll()
	if err != nil {
		clog.Error("%s skel.Get err: %v, req: %v", fun, err, nil)
		feed.ReplyFail(w, lib.CodeSrv)
		return
	}

	resp := &GetResp{
		Feed: feedApi,
	}
	feed.ReplyOk(w, resp)

	// 进行一些异步处理的工作
	go lib.Updates(feedApi, lib.GET, nil)

	return
}
