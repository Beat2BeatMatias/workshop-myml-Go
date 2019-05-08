package myml

import (
    "errors"
    "github.com/gin-gonic/gin"
    "github.com/mercadolibre/workshop-myml/src/api/services/myml"
    "net/http"
    "strconv"
)

func GetOrderInformation(c *gin.Context) {
    orderID := c.Param("orderID")
    order, err := strconv.ParseInt(orderID, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, errors.New("cannot parse order"))
        return
    }

    fullOrder, apiErr := myml.GetOrderInformation(order)
    if apiErr != nil {
        c.JSON(apiErr.Status(), apiErr)
        return
    }

    c.JSON(http.StatusOK, *fullOrder)
}
