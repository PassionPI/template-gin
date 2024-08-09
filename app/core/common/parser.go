package common

import (
	"app-ink/app/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) PaginationQuery(c *gin.Context) *model.Pagination {
	page := c.Query("page")
	size := c.Query("size")

	pagination := model.Pagination{
		Page: 0,
		Size: 10,
	}

	sizeInt, err := strconv.Atoi(size)
	if err == nil && sizeInt > 0 {
		pagination.Size = sizeInt
	}

	pageInt, err := strconv.Atoi(page)
	if err == nil {
		pagination.Page = pageInt
	}

	return &pagination
}
