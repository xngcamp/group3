package feed

import (
	"camp/feed/api"
	"camp/feed/model"
)

func (feed *Feed) Update(feedApi *api.Feed) (err error) {
	feedModel := model.NewFeed()
	feedModel.ID = feedApi.ID
	feedModel.Txt = feedApi.Txt
	if err = feedModel.Update(); err != nil {
		return
	}

	return
}
