// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Mzihao",
            "email": "pimpkin_mak@163.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/healthcheck": {
            "get": {
                "description": "健康检测",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "健康检测模块"
                ],
                "summary": "健康检测",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.RegularResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/changePassword/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "修改密码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "修改密码",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "请求参数data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.PasswordData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.RegularResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/login": {
            "post": {
                "description": "后台登陆",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录控制模块"
                ],
                "summary": "后台登陆",
                "parameters": [
                    {
                        "description": "请求参数data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.LoginResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/logistics/{barcode}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "物流查询",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "物流"
                ],
                "summary": "物流查询",
                "parameters": [
                    {
                        "type": "string",
                        "description": "物流单号",
                        "name": "barcode",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "物流公司代码",
                        "name": "carrierCode",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.SearchLogisticsResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/logistics/{carrierCode}": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "物流下单",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "物流"
                ],
                "summary": "物流下单",
                "parameters": [
                    {
                        "type": "string",
                        "description": "物流商代码",
                        "name": "carrierCode",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "请求参数data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.LogisticsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.AddLogisticsResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/pickUp": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "自提物流下单",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "自提物流"
                ],
                "summary": "自提物流下单",
                "parameters": [
                    {
                        "description": "请求参数data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.PickUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.PickUpResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/pickUp/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "自提物流查询",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "自提物流"
                ],
                "summary": "自提物流查询",
                "parameters": [
                    {
                        "type": "string",
                        "description": "物流单号",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.PickUpResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "自提物流信息修改",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "自提物流"
                ],
                "summary": "自提物流信息修改",
                "parameters": [
                    {
                        "type": "string",
                        "description": "物流单号",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "请求参数data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.PickUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.PickUpResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "自提物流取消订单",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "自提物流"
                ],
                "summary": "自提物流取消订单",
                "parameters": [
                    {
                        "type": "string",
                        "description": "物流单号",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.PickUpResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/user/add": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "请求参数data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.UserRegistered"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.RegularResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/user/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "查询单个用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "查询单个用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.GetUserInfoResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "编辑用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "编辑用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "请求参数data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.UserInfoData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.RegularResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "删除用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.RegularResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "查询用户列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "查询用户列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "页面容量",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "页数",
                        "name": "page_num",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.GetUserInfoResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schemas.AddLogisticsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "物流信息",
                    "$ref": "#/definitions/schemas.Order"
                },
                "message": {
                    "description": "响应信息",
                    "type": "string",
                    "example": "success"
                },
                "status": {
                    "description": "响应状态",
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "schemas.CargoList": {
            "type": "object",
            "properties": {
                "count": {
                    "description": "数量",
                    "type": "integer",
                    "example": 5
                },
                "height": {
                    "description": "高(单位默认为厘米)",
                    "type": "number",
                    "example": 80
                },
                "length": {
                    "description": "长(单位默认为厘米)",
                    "type": "number",
                    "example": 100
                },
                "name": {
                    "description": "货物名称",
                    "type": "string",
                    "example": "货物名称"
                },
                "uint": {
                    "description": "单位， 如个、台、件",
                    "type": "string",
                    "example": "件"
                },
                "volume": {
                    "description": "宽(单位默认为厘米)",
                    "type": "number",
                    "example": 480000
                },
                "weight": {
                    "description": "重量(单位:kg)",
                    "type": "number",
                    "example": 60
                },
                "width": {
                    "description": "宽(单位默认为厘米)",
                    "type": "number",
                    "example": 60
                }
            }
        },
        "schemas.GetUserInfoResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "用户信息",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.UserInfoData"
                    }
                },
                "message": {
                    "description": "响应信息",
                    "type": "string",
                    "example": "成功"
                },
                "status": {
                    "description": "响应状态",
                    "type": "integer",
                    "example": 200
                },
                "total": {
                    "description": "用户数量",
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "schemas.LoginRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string",
                    "maxLength": 120,
                    "minLength": 6,
                    "example": "****"
                },
                "username": {
                    "description": "用户名",
                    "type": "string",
                    "maxLength": 12,
                    "minLength": 4,
                    "example": "name"
                }
            }
        },
        "schemas.LoginResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "Id      uint   ` + "`" + `json:\"id\" form:\"id\" example:\"1\"` + "`" + `             //用户id\nName    string ` + "`" + `json:\"name\" form:\"name\" example:\"string\"` + "`" + `    //用户名",
                    "type": "string",
                    "example": "成功"
                },
                "status": {
                    "description": "响应状态",
                    "type": "integer",
                    "example": 200
                },
                "token": {
                    "description": "令牌",
                    "type": "string",
                    "example": "asdffgh"
                }
            }
        },
        "schemas.Logistics": {
            "type": "object",
            "properties": {
                "barcode": {
                    "description": "运单号",
                    "type": "string"
                },
                "carrier_code": {
                    "description": "物流公司代码",
                    "type": "string"
                },
                "id": {
                    "description": "订单号",
                    "type": "string"
                },
                "trackInfo": {
                    "description": "物流时效",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.TrackInfo"
                    }
                },
                "weblink": {
                    "description": "物流官网",
                    "type": "string"
                }
            }
        },
        "schemas.LogisticsRequest": {
            "type": "object",
            "properties": {
                "cargoList": {
                    "description": "货物明细",
                    "$ref": "#/definitions/schemas.CargoList"
                },
                "delivery_address": {
                    "description": "到件方详细地址",
                    "type": "string",
                    "example": "沙头街道皇庭广场"
                },
                "delivery_city": {
                    "description": "到件方所在地级行政区名称",
                    "type": "string",
                    "example": "深圳市"
                },
                "delivery_contact": {
                    "description": "到件方联系人",
                    "type": "string",
                    "example": "李华"
                },
                "delivery_county": {
                    "description": "到件方所在县/区级行政区名称",
                    "type": "string",
                    "example": "福田区"
                },
                "delivery_province": {
                    "description": "到件方所在省级行政区名称",
                    "type": "string",
                    "example": "广东省"
                },
                "delivery_tel": {
                    "description": "到件方联系电话",
                    "type": "string",
                    "example": "18452154695"
                },
                "expected_pick_up_time": {
                    "description": "希望上门取件时间。格式：yyyy-MM-dd HH:mm:ss。isDocall=1时有效，其中，预约时间超过晚8点，会自动约成第二天。",
                    "type": "string",
                    "example": "2022-05-21 00:00:00"
                },
                "is_do_call": {
                    "description": "是否需要上门,0-不下call，不会通知小哥上门；1-下call，默认2小时内上门",
                    "type": "integer",
                    "example": 1
                },
                "pay_method": {
                    "description": "付款方式(邮费)：1.寄方付2.收方付3.第三方付",
                    "type": "integer",
                    "example": 1
                },
                "remark": {
                    "type": "string",
                    "example": "备注"
                },
                "send_address": {
                    "description": "寄件方详细地址",
                    "type": "string",
                    "example": "粤海街道科兴科学园"
                },
                "send_city": {
                    "description": "寄件方所在地级行政区名称",
                    "type": "string",
                    "example": "深圳市"
                },
                "send_contact": {
                    "description": "寄件方联系人",
                    "type": "string",
                    "example": "张湘"
                },
                "send_county": {
                    "description": "寄件人所在县/区级行政区名称",
                    "type": "string",
                    "example": "南山区"
                },
                "send_province": {
                    "description": "寄件方所在省级行政区名称",
                    "type": "string",
                    "example": "广东省"
                },
                "send_tel": {
                    "description": "寄件方联系电话",
                    "type": "string",
                    "example": "18452154695"
                }
            }
        },
        "schemas.Order": {
            "type": "object",
            "properties": {
                "barcode": {
                    "description": "运单号",
                    "type": "string"
                },
                "carrier_code": {
                    "description": "物流公司代码",
                    "type": "string"
                },
                "weblink": {
                    "description": "物流官网",
                    "type": "string"
                }
            }
        },
        "schemas.PasswordData": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string",
                    "maxLength": 120,
                    "minLength": 6,
                    "example": "****"
                }
            }
        },
        "schemas.PickUp": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "广东省深圳市福田区沙尾街道金地三路888号"
                },
                "id": {
                    "type": "string",
                    "example": "12345678"
                },
                "status": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "schemas.PickUpRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "广东省深圳市福田区沙尾街道金地三路888号"
                },
                "status": {
                    "type": "integer",
                    "enum": [
                        0,
                        1,
                        2,
                        3,
                        4
                    ],
                    "example": 0
                }
            }
        },
        "schemas.PickUpResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "物流信息",
                    "$ref": "#/definitions/schemas.PickUp"
                },
                "message": {
                    "description": "响应信息",
                    "type": "string",
                    "example": "success"
                },
                "status": {
                    "description": "响应状态",
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "schemas.RegularResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "响应信息",
                    "type": "string",
                    "example": "成功"
                },
                "status": {
                    "description": "响应状态",
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "schemas.SearchLogisticsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "物流信息",
                    "$ref": "#/definitions/schemas.Logistics"
                },
                "message": {
                    "description": "响应信息",
                    "type": "string",
                    "example": "success"
                },
                "status": {
                    "description": "响应状态",
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "schemas.TrackInfo": {
            "type": "object",
            "properties": {
                "date": {
                    "description": "时间",
                    "type": "string"
                },
                "status_description": {
                    "description": "物流描述",
                    "type": "string"
                }
            }
        },
        "schemas.UserInfoData": {
            "type": "object",
            "properties": {
                "role": {
                    "description": "角色码",
                    "type": "integer",
                    "example": 1
                },
                "username": {
                    "description": "用户名",
                    "type": "string",
                    "example": "name"
                }
            }
        },
        "schemas.UserRegistered": {
            "type": "object",
            "required": [
                "password",
                "role",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string",
                    "maxLength": 120,
                    "minLength": 6,
                    "example": "****"
                },
                "role": {
                    "description": "角色码",
                    "type": "integer",
                    "minimum": 2,
                    "example": 2
                },
                "username": {
                    "description": "用户名",
                    "type": "string",
                    "maxLength": 12,
                    "minLength": 4,
                    "example": "name"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Logistics Open Api",
	Description:      "聚合物流服务",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
