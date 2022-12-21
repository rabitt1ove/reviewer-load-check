package controllers

import (
    "github.com/gin-gonic/gin"
    "reviewer-load-check/models"
)

func GetPullRequests(c *gin.Context) {
    var githubId string = c.PostForm("githubId")
    var state string    = c.PostForm("state")
    var from string     = c.PostForm("createdFrom")
    var to string       = c.PostForm("createdTo")

    requestParam := models.RequestParam{}
    requestParam.Query = requestParam.GetPullRequestsQuery(githubId, state, from, to)
    responseResponsePullRequests := new(models.ResponsePullRequests)

    post(requestParam, responseResponsePullRequests)

    responsePullRequestsData := models.TabulatePullRequest(responseResponsePullRequests, githubId, state)

    c.JSON(200, responsePullRequestsData)
}
