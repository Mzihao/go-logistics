package v1

import (
	"github.com/gin-gonic/gin"
	"go-logistics/model"
	"go-logistics/service"
	"go-logistics/utils/errmsg"
	"go-logistics/utils/validator"
	"net/http"
)

//type PickUp interface{
//	SearchRouter(id string) (int, map[string]interface{})
//	CreateOrder(address string) (int, map[string]interface{})
//	UpdateOrder(data *PickUp)
//}

// CreateOrder 自提物流下单
// @Tags 自提物流
// @Summary 自提物流下单
// @Description 自提物流下单
// @Accept  json
// @Produce json
// @Param data body model.PickUp true "请求参数data"
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

	pickUp := service.PickUp{Address: data.Address}
	code, result := pickUp.CreateOrder()
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
			"data":    result,
		},
	)
}
