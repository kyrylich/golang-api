package output

import "time"

type SignInUserOutput struct {
	Token string `json:"token"`
}

type SignUpUserOutput struct {
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
}
