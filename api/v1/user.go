package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-logistics/model"
	"go-logistics/utils/errmsg"
	"go-logistics/utils/validator"
	"net/http"
	"strconv"
)

// AddUser 用户注册
// @Tags 用户
// @Summary 用户注册
// @Description 用户注册
// @Accept  json
// @Produce json
// @Param data body model.User true "请求参数data"
// @Router /api/v1/user/add [post]
// @Success 200 {object} schemas.RegularResponse
func AddUser(c *gin.Context) {
	var data model.User
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

	code := model.CheckUser(data.Username)
	if code == errmsg.Success {
		model.CreateUser(&data)
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// GetUserInfo 查询单个用户
// @Tags 用户
// @Summary 查询单个用户
// @Description 查询单个用户
// @Accept  json
// @Produce json
// @Param   id     path    int     true        "用户id"
// @Router /api/v1/user/{id} [get]
// @Success 200 {object} schemas.GetUserInfoResponse
// @Security ApiKeyAuth
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var maps = make(map[string]interface{})
	data, code := model.GetUser(id)
	maps["username"] = data.Username
	maps["role"] = data.Role
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    maps,
			"total":   1,
			"message": errmsg.GetErrMsg(code),
		},
	)

}

// GetUsers 查询用户列表
// @Tags 用户
// @Summary 查询用户列表
// @Description 查询用户列表
// @Accept  json
// @Produce json
// @Param   username     query    string     false        "用户名"
// @Param   page_size     query    string     false        "页面容量"
// @Param   page_num     query    string     false        "页数"
// @Router /api/v1/users [get]
// @Success 200 {object} schemas.GetUserInfoResponse
// @Security ApiKeyAuth
func GetUsers(c *gin.Context) {
	fmt.Println("111")
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	pageNum, _ := strconv.Atoi(c.Query("page_num"))
	username := c.Query("username")
	data, total := model.GetUsers(username, pageSize, pageNum)
	code := errmsg.Success
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// EditUser 编辑用户
// @Tags 用户
// @Summary 编辑用户
// @Description 编辑用户
// @Accept  json
// @Produce json
// @Param   id     path    int     true        "用户id"
// @Param data body schemas.UserInfoData true "请求参数data"
// @Router /api/v1/user/{id} [put]
// @Success 200 {object} schemas.RegularResponse
// @Security ApiKeyAuth
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.CheckUpUser(id, data.Username)
	if code == errmsg.Success {
		model.EditUser(id, &data)
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// ChangeUserPassword 修改密码
// @Tags 用户
// @Summary 修改密码
// @Description 修改密码
// @Accept  json
// @Produce json
// @Param   id     path    int     true        "用户id"
// @Param data body schemas.PasswordData true "请求参数data"
// @Router /api/v1/admin/changePassword/{id} [put]
// @Success 200 {object} schemas.RegularResponse
// @Security ApiKeyAuth
func ChangeUserPassword(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.ChangePassword(id, &data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// DeleteUser 删除用户
// @Tags 用户
// @Summary 删除用户
// @Description 删除用户
// @Accept  json
// @Produce json
// @Param   id     path    int     true        "用户id"
// @Router /api/v1/user/{id} [delete]
// @Success 200 {object} schemas.RegularResponse
// @Security ApiKeyAuth
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteUser(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
