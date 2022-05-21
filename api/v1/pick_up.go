package v1

import (
	"github.com/gin-gonic/gin"
	"go-logistics/model"
	"go-logistics/service"
	"go-logistics/utils/errmsg"
	"go-logistics/utils/validator"
	"net/http"
)

// CreateOrder 自提物流下单
// @Tags 自提物流
// @Summary 自提物流下单
// @Description 自提物流下单
// @Accept  json
// @Produce json
// @Param data body schemas.PickUpRequest true "请求参数data"
// @Router /api/v1/pickUp [post]
// @Success 200 {object} schemas.PickUpResponse
// @Security ApiKeyAuth
func CreateOrder(c *gin.Context) {
	var data model.PickUp
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

	pickUp := service.PickUpServer{Address: data.Address}
	code, result := pickUp.CreateOrder()
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
			"data":    result,
		},
	)
}

// SearchRouter 自提物流查询
// @Tags 自提物流
// @Summary 自提物流查询
// @Description 自提物流查询
// @Accept  json
// @Produce json
// @Param   id     path    string     true        "物流单号"
// @Router /api/v1/pickUp/{id} [get]
// @Success 200 {object} schemas.PickUpResponse
// @Security ApiKeyAuth
func SearchRouter(c *gin.Context) {
	id := c.Param("id")
	pickUp := service.PickUpServer{ID: id}
	code, result := pickUp.SearchRouter()
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
			"data":    result,
		},
	)
}

// UpdateOrder 自提物流信息修改
// @Tags 自提物流
// @Summary 自提物流信息修改
// @Description 自提物流信息修改
// @Accept  json
// @Produce json
// @Param   id     path    string     true        "物流单号"
// @Param data body schemas.PickUpRequest true "请求参数data"
// @Router /api/v1/pickUp/{id} [put]
// @Success 200 {object} schemas.PickUpResponse
// @Security ApiKeyAuth
func UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	var data model.PickUp
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

	pickUp := service.PickUpServer{ID: id, Address: data.Address, Status: data.Status}
	code, result := pickUp.UpdateOrder()
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
			"data":    result,
		},
	)
}

// CancelOrder 自提物流取消订单
// @Tags 自提物流
// @Summary 自提物流取消订单
// @Description 自提物流取消订单
// @Accept  json
// @Produce json
// @Param   id     path    string     true        "物流单号"
// @Router /api/v1/pickUp/{id} [delete]
// @Success 200 {object} schemas.PickUpResponse
// @Security ApiKeyAuth
func CancelOrder(c *gin.Context) {
	id := c.Param("id")
	pickUp := service.PickUpServer{ID: id}
	code, result := pickUp.CancelOrder()
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
			"data":    result,
		},
	)
}
