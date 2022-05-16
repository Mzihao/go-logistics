package schemas

import "go-logistics/model"

type GetUsersResponse struct {
	Status  int          `json:"status" form:"status" example:"200"`  //响应状态
	Data    []model.User `json:"data" form:"data"`                    //用户信息
	Total   int          `json:"total" form:"total" example:"1"`      //用户数量
	Message string       `json:"message" form:"message" example:"成功"` //响应信息
}

type RegularResponse struct {
	Status  int    `json:"status" form:"status" example:"200"`  //响应状态
	Message string `json:"message" form:"message" example:"成功"` //响应信息
}

type UserInfoData struct {
	Username string `json:"username" form:"username" example:"name"` // 用户名
	Role     int    `json:"role" form:"role" example:"1"`            // 角色码
}

type GetUserInfoResponse struct {
	Status  int            `json:"status" form:"status" example:"200"`  //响应状态
	Data    []UserInfoData `json:"data" form:"data"`                    //用户信息
	Total   int            `json:"total" form:"total" example:"1"`      //用户数量
	Message string         `json:"message" form:"message" example:"成功"` //响应信息
}

type PasswordData struct {
	Password string `json:"password" form:"password" example:"****"` //密码
}

type TrackInfo struct {
	Date              string `json:"date"`               // 时间
	StatusDescription string `json:"status_description"` // 物流描述
}

type Logistics struct {
	CarrierCode string      `json:"carrier_code"` // 物流公司代码
	TrackInfo   []TrackInfo `json:"trackInfo"`    // 物流时效
	Weblink     string      `json:"weblink"`      // 物流官网
}

type LogisticsResponse struct {
	Status  int       `json:"status" form:"status" example:"200"`       //响应状态
	Message string    `json:"message" form:"message" example:"success"` //响应信息
	Data    Logistics // 物流信息
}

type LoginResponse struct {
	Status  int    `json:"status" form:"status" example:"200"`   //响应状态
	Id      uint   `json:"id" form:"id" example:"1"`             //用户id
	Name    string `json:"name" form:"name" example:"string"`    //用户名
	Message string `json:"message" form:"message" example:"成功"`  //响应信息
	Token   string `json:"token" form:"token" example:"asdffgh"` //令牌
}

type PickUp struct {
	ID      string `json:"id" form:"id" example:"12345678"`
	Status  uint   `json:"status" form:"status" example:"0"`
	Address string `json:"address" form:"address" example:"广东省深圳市福田区沙尾街道金地三路"`
}

type PickUpResponse struct {
	Status  int    `json:"status" form:"status" example:"200"`       //响应状态
	Message string `json:"message" form:"message" example:"success"` //响应信息
	Data    PickUp // 物流信息
}
