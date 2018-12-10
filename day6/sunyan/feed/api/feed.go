package api

// Skel 用于示例
type Feed struct {
	ID  string `json:"id" bson:"_id"`
	Txt string `json:"txt" bson:"_txt"`
}

func NewFeed() *Feed {
	return &Feed{}
}