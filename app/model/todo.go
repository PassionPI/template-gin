package model

type TodoCreateItem struct {
	Username    string `json:"username" binding:"required,min=4,max=32"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DeadLine    string `json:"dead_line"`
}
