package controllers

import (
    "github.com/gin-gonic/gin"
    "reviewer-load-check/models"
)

func GetTeamMembers(c *gin.Context) {
    var teamName string = c.Query("teamName")

    requestParam := models.RequestParam{}
    requestParam.Query = requestParam.GetTeamMembersQuery(teamName)
    responseTeamMembers := new(models.ResponseTeamMembers)

    post(requestParam, responseTeamMembers)

    c.JSON(200, responseTeamMembers.Data.Organization.Team.Members.Nodes)
}
