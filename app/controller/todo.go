package controller

import (
	"app-ink/app/model"
	"app-ink/pkg/util"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) TodoList(c *gin.Context) {
	username, exists := c.Get("username")

	if !exists {
		util.Err(c, "Failed to get username")
		return
	}

	ctx := c.Request.Context()

	pagination := ctrl.core.Common.Parser.PaginationQuery(c)

	params := ctrl.core.Dep.Pg.TodoFindByUsernameParamsCreate()
	params.SetUsername(username.(string))
	params.SetPagination(pagination)

	todos, err := ctrl.core.Dep.Pg.TodoFindByUsername(ctx, params)

	if err != nil {
		util.Err(c, "Failed to get todo list: "+err.Error())
		return
	}

	util.Ok(c, todos)
}

func (ctrl *Controller) TodoAdd(c *gin.Context) {
	username, exists := c.Get("username")

	if !exists {
		util.Err(c, "Failed to get username")
		return
	}

	ctx := c.Request.Context()

	body, err := util.JSON[model.TodoCreateItem](c)

	if err != nil {
		util.Err(c, err.Error())
		return
	}

	err = ctrl.core.Dep.Pg.TodoInsert(ctx, username.(string), &body)

	if err != nil {
		util.Err(c, "Failed to insert todo"+err.Error())
		return
	}

	util.Ok(c, "Done")
}

func (ctrl *Controller) TodoUpdate(c *gin.Context) {

	ctx := c.Request.Context()

	body, err := util.JSON[model.TodoUpdateItem](c)

	if err != nil {
		util.Err(c, err.Error())
		return
	}

	err = ctrl.core.Dep.Pg.TodoUpdateById(ctx, &body)

	if err != nil {
		util.Err(c, "Failed to update todo")
		return
	}

	util.Ok(c, "Done")
}
