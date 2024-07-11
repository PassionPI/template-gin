package model

type Credentials struct {
	Username string `json:"username" binding:"required,min=4,max=32"`
	Password string `json:"password" binding:"required,min=100"`
}
