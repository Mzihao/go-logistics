package service

import (
	"go-logistics/model"
	"go-logistics/utils"
)

type PickUp struct {
	id      string
	Address string
	status  uint
}

// status{0: 未知, 1: 待取货, 2: 已取货, 3: 取消订单}

func (p PickUp) CreateOrder() (int, map[string]interface{}) {
	id := utils.GetSnowflakeId()
	data := model.PickUp{ID: id, Address: p.Address, Status: 1}
	model.CreatePickUp(&data)
	result := make(map[string]interface{})
	result["id"] = id
	result["address"] = p.Address
	result["status"] = 1
	return 200, result
}
