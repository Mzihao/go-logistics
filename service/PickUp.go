package service

import (
	"go-logistics/model"
	"go-logistics/utils"
)

type PickUp struct {
	ID      string
	Address string
	status  uint
}

var statusMap = map[uint]string{0: "未知", 1: "备货中", 2: "待取货", 3: "已取货", 4: "取消订单"}

func (p PickUp) CreateOrder() (int, map[string]interface{}) {
	id := utils.GetSnowflakeId()
	data := model.PickUp{ID: id, Address: p.Address, Status: 1}
	model.CreatePickUp(&data)
	result := make(map[string]interface{})
	result["id"] = id
	result["address"] = p.Address
	result["status"] = statusMap[1]
	return 200, result
}

func (p PickUp) SearchRouter() (int, map[string]interface{}) {
	result := make(map[string]interface{})
	id, status, address := model.GetPickUp(p.ID)
	if id == "" {
		return 4002, result
	}
	result["id"] = id
	result["address"] = address
	result["status"] = statusMap[status]
	return 200, result
}

//func (p PickUp) UpdateOrder() (int, map[string]interface{}) {
//	result := make(map[string]interface{})
//	data := model.PickUp{ID: p.ID, Status: p.status, Address: p.Address}
//	model.EditPickUp(p.ID, &data)
//}
