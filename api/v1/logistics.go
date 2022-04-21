package v1

import (
	"github.com/gin-gonic/gin"
	"go-logistics/service"
	"go-logistics/utils/errmsg"
	"net/http"
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
// @Success 200 {object} schemas.LogisticsResponse
// @Security ApiKeyAuth
func QueryExpress(c *gin.Context) {
	barcode := c.Param("barcode")
	carrierCode := c.Query("carrierCode")
	serviceMap := make(map[string]func(string) (int, map[string]interface{}))
	serviceMap["bld-express"] = service.MapleLogisticsService
	server := serviceMap[carrierCode]
	code, data := server(barcode)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
