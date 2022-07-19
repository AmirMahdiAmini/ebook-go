package dto

type VerifyAccount struct {
	Phone string `json:"phone"`
	Code  int    `json:"code"`
}
type Signin struct {
	Username_or_phone string `json:"username_or_phone"`
	Password          string `json:"password"`
}
type ForgotPassword struct {
	Phone string `json:"phone"`
}
type ChangePassowrd struct {
	Code     int    `json:"code"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
