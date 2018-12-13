package api
type User struct {
	Nick string `json:"Nick"`
	Sex int `json:"Sex"`
	Email string `json:"Email"`
	Password string `json:"Password"`
}
func NewUser() *User {
	return &User{}
}
