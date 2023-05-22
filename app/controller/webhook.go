package controller

import (
	"fmt"
	"net/http"

	"app.land.x/app/model"
	"app.land.x/pkg/req"

	"github.com/gin-gonic/gin"
)

// WebhookGithubPR is a webhook handler for Github pull request
//
// 主要接收Merge Preview & Merge Master Pull Request事件
func (ctrl *Controller) WebhookGithubPR(c *gin.Context) {
	payload, err := req.JSON[model.PullRequestPayload](c)

	if err != nil {
		c.String(http.StatusOK, "OK")
		return
	}

	branch := payload.PullRequest.Base.Ref
	if payload.Action == "closed" && payload.PullRequest.Merged && branch == "preview" {
		repository := payload.Repository.Name
		prTitle := payload.PullRequest.Title
		prNumber := payload.PullRequest.Number
		fmt.Printf(
			"Pull request %s (#%d) merged from %s to %s branch\n",
			prTitle,
			prNumber,
			repository,
			branch,
		)

	}

	c.String(http.StatusOK, "OK")
}
