package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-logistics/service"
	"go-logistics/utils/errmsg"
	"net/http"
	"strconv"
)

// QueryExpress 物流查询
// @Tags 物流
// @Summary 物流查询
// @Description 物流查询
// @Accept  json
// @Produce json
// @Param   barcode     path    string     true        "物流单号"
// @Param   carrierCode     query    string     false        "物流公司代码"
// @Router /api/v1/logistics/{barcode} [get]
// @Security ApiKeyAuth
func QueryExpress(c *gin.Context) {
	barcode := c.Param("barcode")
	carrierCode, _ := strconv.Atoi(c.Query("carrierCode"))
	fmt.Println(barcode, carrierCode)
	code, data := service.GetMapleLogistics(barcode)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
