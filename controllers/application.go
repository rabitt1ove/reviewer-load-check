package controllers

import (
    "github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
    c.Redirect(302, "/views")
}
