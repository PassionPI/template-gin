package model

type TodoCreateItem struct {
	ID          int    `json:"id"`   // 可选字段
	Done        bool   `json:"done"` // 可选字段
	Title       string `json:"title"`
	Description string `json:"description,omitempty"` // 可选字段
}
