package app

import (
    "github.com/gin-gonic/gin"
    "github.com/mercadolibre/workshop-myml/src/api/controllers/myml"
    "github.com/mercadolibre/workshop-myml/src/api/controllers/ping"
)

func StartApp() {
    router := gin.Default()
    router.GET("/ping", ping.Ping)
    router.GET("/myml/:orderID", myml.GetOrderInformation)
    router.Run(":8080")
}
