package healthcheck

import (
	"github.com/gin-gonic/gin"
	"go-logistics/utils/errmsg"
	"net/http"
)

// HealthCheck 健康检测
// @Tags 健康检测模块
// @Summary 健康检测
// @Description 健康检测
// @Accept  json
// @Produce json
// @Router /api/healthcheck [get]
// @Success 200 {object} schemas.RegularResponse
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": errmsg.GetErrMsg(200),
	})
}
