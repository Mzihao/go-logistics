---
title: logistics v1.0.0
---

# logistics

> v1.0.0

# 用户

## PUT 编辑用户

PUT /api/v1/user/{id}

编辑用户

> Body 请求参数

```json
{
  "role": 1,
  "username": "name"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|integer| 是 |用户id|
|body|body|[schemas.UserInfoData](#schemaschemas.userinfodata)| 否 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[schemas.RegularResponse](#schemaschemas.regularresponse)|

## GET 查询单个用户

GET /api/v1/user/{id}

查询单个用户

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|integer| 是 |用户id|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[schemas.GetUserInfoResponse](#schemaschemas.getuserinforesponse)|

## DELETE 删除用户

DELETE /api/v1/user/{id}

删除用户

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|integer| 是 |用户id|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[schemas.RegularResponse](#schemaschemas.regularresponse)|

## POST 用户注册

POST /api/v1/user/add

用户注册

> Body 请求参数

```json
{
  "password": "****",
  "role": 2,
  "username": "name"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|[schemas.UserRegistered](#schemaschemas.userregistered)| 否 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[schemas.RegularResponse](#schemaschemas.regularresponse)|

## GET 查询用户列表

GET /api/v1/users

查询用户列表

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|username|query|string| 否 |用户名|
|page_size|query|string| 否 |页面容量|
|page_num|query|string| 否 |页数|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[schemas.GetUserInfoResponse](#schemaschemas.getuserinforesponse)|

## PUT 修改密码

PUT /api/v1/admin/changePassword/{id}

修改密码

> Body 请求参数

```json
{
  "password": "****"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|integer| 是 |用户id|
|body|body|[schemas.PasswordData](#schemaschemas.passworddata)| 否 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[schemas.RegularResponse](#schemaschemas.regularresponse)|

# 登录控制模块

## POST 后台登陆

POST /api/v1/login

后台登陆

> Body 请求参数

```json
{
  "password": "****",
  "username": "name"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|[schemas.LoginRequest](#schemaschemas.loginrequest)| 否 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[schemas.LoginResponse](#schemaschemas.loginresponse)|

# 物流

## POST 物流下单

POST /api/v1/logistics/{carrierCode}

物流下单

> Body 请求参数

```json
{
  "cargoList": {
    "count": 5,
    "height": 80,
    "length": 100,
    "name": "货物名称",
    "uint": "件",
    "volume": 480000,
    "weight": 60,
    "width": 60
  },
  "delivery_address": "沙头街道皇庭广场",
  "delivery_city": "深圳市",
  "delivery_contact": "李华",
  "delivery_county": "福田区",
  "delivery_province": "广东省",
  "delivery_tel": "18452154695",
  "expected_pick_up_time": "2022-05-21 00:00:00",
  "is_do_call": 1,
  "pay_method": 1,
  "remark": "备注",
  "send_address": "粤海街道科兴科学园",
  "send_city": "深圳市",
  "send_contact": "张湘",
  "send_county": "南山区",
  "send_province": "广东省",
  "send_tel": "18452154695"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|carrierCode|path|string| 是 |物流商代码|
|body|body|[schemas.LogisticsRequest](#schemaschemas.logisticsrequest)| 否 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[schemas.AddLogisticsResponse](#schemaschemas.addlogisticsresponse)|

## GET 物流查询

GET /api/v1/logistics/{barcode}

物流查询

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|barcode|path|string| 是 |物流单号|
|carrierCode|query|string| 否 |物流公司代码|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[schemas.SearchLogisticsResponse](#schemaschemas.searchlogisticsresponse)|

# 自提物流

## POST 自提物流下单

POST /api/v1/pickUp

自提物流下单

> Body 请求参数

```json
{
  "address": "广东省深圳市福田区沙尾街道金地三路888号",
  "status": 0
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|[schemas.PickUpRequest](#schemaschemas.pickuprequest)| 否 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[schemas.PickUpResponse](#schemaschemas.pickupresponse)|

## DELETE 自提物流取消订单

DELETE /api/v1/pickUp/{id}

自提物流取消订单

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |物流单号|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[schemas.PickUpResponse](#schemaschemas.pickupresponse)|

## GET 自提物流查询

GET /api/v1/pickUp/{id}

自提物流查询

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |物流单号|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[schemas.PickUpResponse](#schemaschemas.pickupresponse)|

## PUT 自提物流信息修改

PUT /api/v1/pickUp/{id}

自提物流信息修改

> Body 请求参数

```json
{
  "address": "广东省深圳市福田区沙尾街道金地三路888号",
  "status": 0
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |物流单号|
|body|body|[schemas.PickUpRequest](#schemaschemas.pickuprequest)| 否 |none|

> 返回示例

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[schemas.PickUpResponse](#schemaschemas.pickupresponse)|

# 数据模型

<h2 id="tocS_schemas.LoginResponse">schemas.LoginResponse</h2>

<a id="schemaschemas.loginresponse"></a>
<a id="schema_schemas.LoginResponse"></a>
<a id="tocSschemas.loginresponse"></a>
<a id="tocsschemas.loginresponse"></a>

```json
{
  "id": 1,
  "message": "成功",
  "name": "string",
  "status": 200,
  "token": "asdffgh"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|integer|false|none||用户id|
|message|string|false|none||响应信息|
|name|string|false|none||用户名|
|status|integer|false|none||响应状态|
|token|string|false|none||令牌|

<h2 id="tocS_schemas.SearchLogisticsResponse">schemas.SearchLogisticsResponse</h2>

<a id="schemaschemas.searchlogisticsresponse"></a>
<a id="schema_schemas.SearchLogisticsResponse"></a>
<a id="tocSschemas.searchlogisticsresponse"></a>
<a id="tocsschemas.searchlogisticsresponse"></a>

```json
{
  "data": {
    "barcode": "string",
    "carrier_code": "string",
    "id": "string",
    "trackInfo": [
      {
        "date": "string",
        "status_description": "string"
      }
    ],
    "weblink": "string"
  },
  "message": "success",
  "status": 200
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|data|[schemas.Logistics](#schemaschemas.logistics)|false|none||none|
|message|string|false|none||响应信息|
|status|integer|false|none||响应状态|

<h2 id="tocS_schemas.PickUpResponse">schemas.PickUpResponse</h2>

<a id="schemaschemas.pickupresponse"></a>
<a id="schema_schemas.PickUpResponse"></a>
<a id="tocSschemas.pickupresponse"></a>
<a id="tocsschemas.pickupresponse"></a>

```json
{
  "data": {
    "address": "广东省深圳市福田区沙尾街道金地三路888号",
    "id": "12345678",
    "status": 0
  },
  "message": "success",
  "status": 200
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|data|[schemas.PickUp](#schemaschemas.pickup)|false|none||none|
|message|string|false|none||响应信息|
|status|integer|false|none||响应状态|

<h2 id="tocS_schemas.PickUp">schemas.PickUp</h2>

<a id="schemaschemas.pickup"></a>
<a id="schema_schemas.PickUp"></a>
<a id="tocSschemas.pickup"></a>
<a id="tocsschemas.pickup"></a>

```json
{
  "address": "广东省深圳市福田区沙尾街道金地三路888号",
  "id": "12345678",
  "status": 0
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|address|string|false|none||none|
|id|string|false|none||none|
|status|integer|false|none||none|

<h2 id="tocS_schemas.AddLogisticsResponse">schemas.AddLogisticsResponse</h2>

<a id="schemaschemas.addlogisticsresponse"></a>
<a id="schema_schemas.AddLogisticsResponse"></a>
<a id="tocSschemas.addlogisticsresponse"></a>
<a id="tocsschemas.addlogisticsresponse"></a>

```json
{
  "data": {
    "barcode": "string",
    "carrier_code": "string",
    "weblink": "string"
  },
  "message": "success",
  "status": 200
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|data|[schemas.Order](#schemaschemas.order)|false|none||none|
|message|string|false|none||响应信息|
|status|integer|false|none||响应状态|

<h2 id="tocS_schemas.Logistics">schemas.Logistics</h2>

<a id="schemaschemas.logistics"></a>
<a id="schema_schemas.Logistics"></a>
<a id="tocSschemas.logistics"></a>
<a id="tocsschemas.logistics"></a>

```json
{
  "barcode": "string",
  "carrier_code": "string",
  "id": "string",
  "trackInfo": [
    {
      "date": "string",
      "status_description": "string"
    }
  ],
  "weblink": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|barcode|string|false|none||运单号|
|carrier_code|string|false|none||物流公司代码|
|id|string|false|none||订单号|
|trackInfo|[[schemas.TrackInfo](#schemaschemas.trackinfo)]|false|none||物流时效|
|weblink|string|false|none||物流官网|

<h2 id="tocS_schemas.LoginRequest">schemas.LoginRequest</h2>

<a id="schemaschemas.loginrequest"></a>
<a id="schema_schemas.LoginRequest"></a>
<a id="tocSschemas.loginrequest"></a>
<a id="tocsschemas.loginrequest"></a>

```json
{
  "password": "****",
  "username": "name"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|password|string|false|none||密码|
|username|string|false|none||用户名|

<h2 id="tocS_schemas.LogisticsRequest">schemas.LogisticsRequest</h2>

<a id="schemaschemas.logisticsrequest"></a>
<a id="schema_schemas.LogisticsRequest"></a>
<a id="tocSschemas.logisticsrequest"></a>
<a id="tocsschemas.logisticsrequest"></a>

```json
{
  "cargoList": {
    "count": 5,
    "height": 80,
    "length": 100,
    "name": "货物名称",
    "uint": "件",
    "volume": 480000,
    "weight": 60,
    "width": 60
  },
  "delivery_address": "沙头街道皇庭广场",
  "delivery_city": "深圳市",
  "delivery_contact": "李华",
  "delivery_county": "福田区",
  "delivery_province": "广东省",
  "delivery_tel": "18452154695",
  "expected_pick_up_time": "2022-05-21 00:00:00",
  "is_do_call": 1,
  "pay_method": 1,
  "remark": "备注",
  "send_address": "粤海街道科兴科学园",
  "send_city": "深圳市",
  "send_contact": "张湘",
  "send_county": "南山区",
  "send_province": "广东省",
  "send_tel": "18452154695"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|cargoList|[schemas.CargoList](#schemaschemas.cargolist)|false|none||none|
|delivery_address|string|false|none||到件方详细地址|
|delivery_city|string|false|none||到件方所在地级行政区名称|
|delivery_contact|string|false|none||到件方联系人|
|delivery_county|string|false|none||到件方所在县/区级行政区名称|
|delivery_province|string|false|none||到件方所在省级行政区名称|
|delivery_tel|string|false|none||到件方联系电话|
|expected_pick_up_time|string|false|none||希望上门取件时间。格式：yyyy-MM-dd HH:mm:ss。isDocall=1时有效，其中，预约时间超过晚8点，会自动约成第二天。|
|is_do_call|integer|false|none||是否需要上门,0-不下call，不会通知小哥上门；1-下call，默认2小时内上门|
|pay_method|integer|false|none||付款方式(邮费)：1.寄方付2.收方付3.第三方付|
|remark|string|false|none||none|
|send_address|string|false|none||寄件方详细地址|
|send_city|string|false|none||寄件方所在地级行政区名称|
|send_contact|string|false|none||寄件方联系人|
|send_county|string|false|none||寄件人所在县/区级行政区名称|
|send_province|string|false|none||寄件方所在省级行政区名称|
|send_tel|string|false|none||寄件方联系电话|

<h2 id="tocS_schemas.CargoList">schemas.CargoList</h2>

<a id="schemaschemas.cargolist"></a>
<a id="schema_schemas.CargoList"></a>
<a id="tocSschemas.cargolist"></a>
<a id="tocsschemas.cargolist"></a>

```json
{
  "count": 5,
  "height": 80,
  "length": 100,
  "name": "货物名称",
  "uint": "件",
  "volume": 480000,
  "weight": 60,
  "width": 60
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|count|integer|false|none||数量|
|height|number|false|none||高(单位默认为厘米)|
|length|number|false|none||长(单位默认为厘米)|
|name|string|false|none||货物名称|
|uint|string|false|none||单位， 如个、台、件|
|volume|number|false|none||宽(单位默认为厘米)|
|weight|number|false|none||重量(单位:kg)|
|width|number|false|none||宽(单位默认为厘米)|

<h2 id="tocS_schemas.UserRegistered">schemas.UserRegistered</h2>

<a id="schemaschemas.userregistered"></a>
<a id="schema_schemas.UserRegistered"></a>
<a id="tocSschemas.userregistered"></a>
<a id="tocsschemas.userregistered"></a>

```json
{
  "password": "****",
  "role": 2,
  "username": "name"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|password|string|false|none||密码|
|role|integer|false|none||角色码|
|username|string|false|none||用户名|

<h2 id="tocS_schemas.RegularResponse">schemas.RegularResponse</h2>

<a id="schemaschemas.regularresponse"></a>
<a id="schema_schemas.RegularResponse"></a>
<a id="tocSschemas.regularresponse"></a>
<a id="tocsschemas.regularresponse"></a>

```json
{
  "message": "成功",
  "status": 200
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|message|string|false|none||响应信息|
|status|integer|false|none||响应状态|

<h2 id="tocS_schemas.GetUserInfoResponse">schemas.GetUserInfoResponse</h2>

<a id="schemaschemas.getuserinforesponse"></a>
<a id="schema_schemas.GetUserInfoResponse"></a>
<a id="tocSschemas.getuserinforesponse"></a>
<a id="tocsschemas.getuserinforesponse"></a>

```json
{
  "data": [
    {
      "role": 1,
      "username": "name"
    }
  ],
  "message": "成功",
  "status": 200,
  "total": 1
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|data|[[schemas.UserInfoData](#schemaschemas.userinfodata)]|false|none||用户信息|
|message|string|false|none||响应信息|
|status|integer|false|none||响应状态|
|total|integer|false|none||用户数量|

<h2 id="tocS_schemas.TrackInfo">schemas.TrackInfo</h2>

<a id="schemaschemas.trackinfo"></a>
<a id="schema_schemas.TrackInfo"></a>
<a id="tocSschemas.trackinfo"></a>
<a id="tocsschemas.trackinfo"></a>

```json
{
  "date": "string",
  "status_description": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|date|string|false|none||时间|
|status_description|string|false|none||物流描述|

<h2 id="tocS_schemas.PickUpRequest">schemas.PickUpRequest</h2>

<a id="schemaschemas.pickuprequest"></a>
<a id="schema_schemas.PickUpRequest"></a>
<a id="tocSschemas.pickuprequest"></a>
<a id="tocsschemas.pickuprequest"></a>

```json
{
  "address": "广东省深圳市福田区沙尾街道金地三路888号",
  "status": 0
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|address|string|false|none||none|
|status|integer|false|none||none|

<h2 id="tocS_schemas.Order">schemas.Order</h2>

<a id="schemaschemas.order"></a>
<a id="schema_schemas.Order"></a>
<a id="tocSschemas.order"></a>
<a id="tocsschemas.order"></a>

```json
{
  "barcode": "string",
  "carrier_code": "string",
  "weblink": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|barcode|string|false|none||运单号|
|carrier_code|string|false|none||物流公司代码|
|weblink|string|false|none||物流官网|

<h2 id="tocS_schemas.PasswordData">schemas.PasswordData</h2>

<a id="schemaschemas.passworddata"></a>
<a id="schema_schemas.PasswordData"></a>
<a id="tocSschemas.passworddata"></a>
<a id="tocsschemas.passworddata"></a>

```json
{
  "password": "****"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|password|string|false|none||密码|

<h2 id="tocS_schemas.UserInfoData">schemas.UserInfoData</h2>

<a id="schemaschemas.userinfodata"></a>
<a id="schema_schemas.UserInfoData"></a>
<a id="tocSschemas.userinfodata"></a>
<a id="tocsschemas.userinfodata"></a>

```json
{
  "role": 1,
  "username": "name"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|role|integer|false|none||角色码|
|username|string|false|none||用户名|

