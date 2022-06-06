package v1

import (
	"github.com/gin-gonic/gin"
	"go-logistics/model"
	"go-logistics/service"
	"go-logistics/utils/errmsg"
	"go-logistics/utils/validator"
	"net/http"
)

type SearchRouterInterface interface {
	SearchRouter(barcode string) (int, map[string]interface{})
}

type PlaceOrderInterface interface {
	CreateOrder() (int, map[string]string)
}

// QueryExpress 物流查询
// @Tags 物流
// @Summary 物流查询
// @Description 物流查询
// @Accept  json
// @Produce json
// @Param   barcode     path    string     true        "物流单号"
// @Param   carrierCode     query    string     false        "物流公司代码"
// @Router /api/v1/logistics/{barcode} [get]
// @Success 200 {object} schemas.SearchLogisticsResponse
// @Security ApiKeyAuth
func QueryExpress(c *gin.Context) {
	barcode := c.Param("barcode")
	carrierCode := c.Query("carrierCode")

	// carrierCode为空则为自营单号，到数据库查询barcode和carrierCode
	if carrierCode == "" {
		barcode, carrierCode = model.GetLogisticsById(barcode)
		// barcode为空则抛出异常(查无此单号)
		if barcode == "" {
			c.JSON(http.StatusOK, gin.H{
				"status":  4002,
				"data":    nil,
				"message": errmsg.GetErrMsg(4002),
			})
			return
		}
	}

	//定义服务转发映射
	serviceMap := make(map[string]SearchRouterInterface)
	serviceMap["bld-express"] = service.MapleLogisticsServer{}
	serviceMap["zeek"] = service.ZeekServer{}
	serviceMap["sf-express"] = service.ShunFengService{}
	serviceMap["kuaidi100"] = service.Kuaidi100Server{}
	// ...可扩展

	// 获取服务
	// 服务不存在
	if val, ok := serviceMap[carrierCode]; !ok {
		c.JSON(http.StatusOK, gin.H{
			"status":  4003,
			"data":    nil,
			"message": errmsg.GetErrMsg(4003),
		})
		return
	}

	server, err := serviceMap[carrierCode]
	//server := service.MapleLogistics{}
	// if !err {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status":  4003,
	// 		"data":    nil,
	// 		"message": errmsg.GetErrMsg(4003),
	// 	})
	// 	return
	// }

	// 查询
	code, data := server.SearchRouter(barcode)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// PlaceOrder 物流下单
// @Tags 物流
// @Summary 物流下单
// @Description 物流下单
// @Accept  json
// @Produce json
// @Param   carrierCode     path    string     true        "物流商代码"
// @Param data body schemas.LogisticsRequest true "请求参数data"
// @Router /api/v1/logistics/{carrierCode} [post]
// @Success 200 {object} schemas.AddLogisticsResponse
// @Security ApiKeyAuth
func PlaceOrder(c *gin.Context) {
	carrierCode := c.Param("carrierCode")
	var data model.Logistics
	var msg string
	var validCode int
	_ = c.ShouldBindJSON(&data)

	msg, validCode = validator.Validate(&data)
	if validCode != errmsg.Success {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  validCode,
				"message": msg,
			},
		)
		c.Abort()
		return
	}

	//定义服务转发映射
	serviceMap := make(map[string]PlaceOrderInterface)
	serviceMap["sf-express"] = service.ShunFengService{}
	// ...可扩展

	// 获取服务
	server, err := serviceMap[carrierCode]
	if !err {
		c.JSON(http.StatusOK, gin.H{
			"status":  4003,
			"data":    nil,
			"message": errmsg.GetErrMsg(4003),
		})
		return
	}

	// 下单
	code, result := server.CreateOrder()
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    result,
		"message": errmsg.GetErrMsg(code),
	})
}
