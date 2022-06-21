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

type UserRegistered struct {
	Username string `json:"username" form:"username" example:"name" validate:"required,min=4,max=12"`  // 用户名
	Password string `json:"password" form:"password" example:"****" validate:"required,min=6,max=120"` //密码
	Role     int    `json:"role" form:"role" example:"2" validate:"required,gte=2"`                    // 角色码
}

type GetUserInfoResponse struct {
	Status  int            `json:"status" form:"status" example:"200"`  //响应状态
	Data    []UserInfoData `json:"data" form:"data"`                    //用户信息
	Total   int            `json:"total" form:"total" example:"1"`      //用户数量
	Message string         `json:"message" form:"message" example:"成功"` //响应信息
}

type PasswordData struct {
	Password string `json:"password" form:"password" example:"****" validate:"required,min=6,max=120"` //密码
}

type LoginRequest struct {
	Username string `json:"username" form:"username" example:"name" validate:"required,min=4,max=12"`  // 用户名
	Password string `json:"password" form:"password" example:"****" validate:"required,min=6,max=120"` //密码
}

type TrackInfo struct {
	Date              string `json:"date"`               // 时间
	StatusDescription string `json:"status_description"` // 物流描述
}

type Logistics struct {
	Id          string      `json:"id"`           // 订单号
	Barcode     string      `json:"barcode"`      // 运单号
	CarrierCode string      `json:"carrier_code"` // 物流公司代码
	TrackInfo   []TrackInfo `json:"trackInfo"`    // 物流时效
	Weblink     string      `json:"weblink"`      // 物流官网
}

type SearchLogisticsResponse struct {
	Status  int       `json:"status" form:"status" example:"200"`       //响应状态
	Message string    `json:"message" form:"message" example:"success"` //响应信息
	Data    Logistics // 物流信息
}

type Order struct {
	Barcode     string `json:"barcode"`      // 运单号
	CarrierCode string `json:"carrier_code"` // 物流公司代码
	Weblink     string `json:"weblink"`      // 物流官网
}

type AddLogisticsResponse struct {
	Status  int    `json:"status" form:"status" example:"200"`       //响应状态
	Message string `json:"message" form:"message" example:"success"` //响应信息
	Data    Order  // 物流信息
}

type User struct {
}

type LoginResponse struct {
	Status int `json:"status" form:"status" example:"200"` //响应状态
	//Id      uint   `json:"id" form:"id" example:"1"`             //用户id
	//Name    string `json:"name" form:"name" example:"string"`    //用户名
	Message string `json:"message" form:"message" example:"成功"`  //响应信息
	Token   string `json:"token" form:"token" example:"asdffgh"` //令牌
}

type PickUp struct {
	ID      string `json:"id" form:"id" example:"12345678"`
	Status  uint   `json:"status" form:"status" example:"0"`
	Address string `json:"address" form:"address" example:"广东省深圳市福田区沙尾街道金地三路888号"`
}

type PickUpRequest struct {
	Status  uint   `json:"status" form:"status" example:"0" validate:"oneof=0 1 2 3 4" label:"状态码"`
	Address string `json:"address" form:"address" example:"广东省深圳市福田区沙尾街道金地三路888号"`
}

type PickUpResponse struct {
	Status  int    `json:"status" form:"status" example:"200"`       //响应状态
	Message string `json:"message" form:"message" example:"success"` //响应信息
	Data    PickUp // 物流信息
}

type CargoList struct {
	Name   string  `json:"name" form:"name" example:"货物名称"`         // 货物名称
	Unit   string  `json:"uint" form:"uint" example:"件"`            // 单位， 如个、台、件
	Count  int     `json:"count" form:"count" example:"5"`          // 数量
	Length float32 `json:"length" form:"length" example:"100.0"`    // 长(单位默认为厘米)
	Height float32 `json:"height" form:"height" example:"80.0"`     // 高(单位默认为厘米)
	Width  float32 `json:"width" form:"width" example:"60.0"`       // 宽(单位默认为厘米)
	Volume float32 `json:"volume" form:"volume" example:"480000.0"` // 宽(单位默认为厘米)
	Weight float32 `json:"weight" form:"weight" example:"60.0"`     // 重量(单位:kg)
}

type LogisticsRequest struct {
	SendContact        string    `json:"send_contact" form:"send_contact" example:"张湘"`                                    // 寄件方联系人
	SendTel            string    `json:"send_tel" form:"send_tel" example:"18452154695"`                                   // 寄件方联系电话
	SendProvince       string    `json:"send_province" form:"send_province" example:"广东省"`                                 // 寄件方所在省级行政区名称
	SendCity           string    `json:"send_city" form:"send_city" example:"深圳市"`                                         // 寄件方所在地级行政区名称
	SendCounty         string    `json:"send_county" form:"send_county" example:"南山区"`                                     // 寄件人所在县/区级行政区名称
	SendAddress        string    `json:"send_address" form:"send_address" example:"粤海街道科兴科学园"`                             // 寄件方详细地址
	IsDoCall           uint      `json:"is_do_call" form:"is_do_call" example:"1"`                                         // 是否需要上门,0-不下call，不会通知小哥上门；1-下call，默认2小时内上门
	ExpectedPickUpTime string    `json:"expected_pick_up_time" form:"expected_pick_up_time" example:"2022-05-21 00:00:00"` // 希望上门取件时间。格式：yyyy-MM-dd HH:mm:ss。isDocall=1时有效，其中，预约时间超过晚8点，会自动约成第二天。
	DeliveryContact    string    `json:"delivery_contact" form:"delivery_contact" example:"李华"`                            // 到件方联系人
	DeliveryTel        string    `json:"delivery_tel" form:"delivery_tel" example:"18452154695"`                           // 到件方联系电话
	DeliveryProvince   string    `json:"delivery_province" form:"delivery_province" example:"广东省"`                         // 到件方所在省级行政区名称
	DeliveryCity       string    `json:"delivery_city" form:"delivery_city" example:"深圳市"`                                 // 到件方所在地级行政区名称
	DeliveryCounty     string    `json:"delivery_county" form:"delivery_county" example:"福田区"`                             // 到件方所在县/区级行政区名称
	DeliveryAddress    string    `json:"delivery_address" form:"delivery_address" example:"沙头街道皇庭广场"`                      // 到件方详细地址
	PayMethod          uint      `json:"pay_method" form:"pay_method" example:"1"`                                         // 付款方式(邮费)：1.寄方付2.收方付3.第三方付
	Remark             string    `json:"remark" form:"remark" example:"备注"`
	CargoList          CargoList // 货物明细
}
