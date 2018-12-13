package api
type Feed struct {
	Id int `json:"id"`
	Title string
	Txt   string `json:"txt"`
	Author string
}
func NewFeed() *Feed {
	return &Feed{}
}
