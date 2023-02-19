package input

type SignUpUserInput struct {
	Username string `json:"username" binding:"required,min=5,max=64"`
	Password string `json:"password" binding:"required,min=8,max=128"`
}

type SignInUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
