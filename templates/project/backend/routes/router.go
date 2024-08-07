package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/evanmschultz/gopostulate_dev/app/controllers"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()
    router.GET("/", controllers.Index)
    return router
}
