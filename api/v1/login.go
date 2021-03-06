package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-logistics/middleware"
	"go-logistics/model"
	"go-logistics/utils/errmsg"
	"net/http"
	"time"
)

// Login 后台登陆
// @Tags 登录控制模块
// @Summary 后台登陆
// @Description 后台登陆
// @Accept  json
// @Produce json
// @Param data body schemas.LoginRequest true "请求参数data"
// @Router /api/v1/login [post]
// @Success 200 {object} schemas.LoginResponse
func Login(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	var token string
	var code int

	formData, code = model.CheckLogin(formData.Username, formData.Password)

	if code == errmsg.Success {
		setToken(c, formData)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			//"name":    formData.Username,
			//"id":      formData.ID,
			"message": errmsg.GetErrMsg(code),
			"token":   token,
		})
	}

}

// token生成函数
func setToken(c *gin.Context, user model.User) {
	j := middleware.NewJWT()
	claims := middleware.MyClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 604800,
			Issuer:    "GinBlog",
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.Error,
			"message": errmsg.GetErrMsg(errmsg.Error),
			"token":   token,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    user.Username,
		"id":      user.ID,
		"message": errmsg.GetErrMsg(200),
		"token":   "Bearer " + token,
	})
	return
}
