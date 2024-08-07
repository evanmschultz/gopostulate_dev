package controllers

import (
    "github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
    c.String(200, "Welcome to GoPostulate!")
}
