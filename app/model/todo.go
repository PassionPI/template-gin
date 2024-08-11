package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type TodoScanItem struct {
	ID          int              `json:"id"`
	Done        bool             `json:"done"`
	Title       pgtype.Text      `json:"title"`
	Description pgtype.Text      `json:"description"`
	DeadLine    pgtype.Timestamp `json:"deadline"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

type TodoCreateItem struct {
	Title       string `json:"title"`                 // 标题
	DeadLine    string `json:"deadline,omitempty"`    // 可选字段
	Description string `json:"description,omitempty"` // 可选字段
}

type TodoUpdateItem struct {
	ID          int    `json:"id"`
	Done        bool   `json:"done"`
	Title       string `json:"title"`                 // 标题
	DeadLine    string `json:"deadline,omitempty"`    // 可选字段
	Description string `json:"description,omitempty"` // 可选字段
	Username    string `json:"username,omitempty"`    // 可选字段
}
