package service

import (
	"go-logistics/model"
	"go-logistics/utils"
)

type PickUpServer struct {
	ID      string
	Address string
	Status  uint
}

var statusMap = map[uint]string{0: "未知", 1: "备货中", 2: "待取货", 3: "已取货", 4: "订单取消"}

func (p PickUpServer) CreateOrder() (int, map[string]interface{}) {
	id := utils.GetSnowflakeId()
	data := model.PickUp{ID: id, Address: p.Address, Status: 1}
	model.CreatePickUp(&data)
	result := make(map[string]interface{})
	result["id"] = id
	result["address"] = p.Address
	result["status"] = statusMap[1]
	return 200, result
}

func (p PickUpServer) SearchRouter() (int, map[string]interface{}) {
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

func (p PickUpServer) UpdateOrder() (int, map[string]interface{}) {
	result := make(map[string]interface{})
	// 判断单号是否存在
	if model.HasOrder(p.ID) == 4002 {
		return 4002, result
	}
	data := model.PickUp{Status: p.Status, Address: p.Address}
	model.EditPickUp(p.ID, &data)
	id, status, address := model.GetPickUp(p.ID)
	result["id"] = id
	result["address"] = address
	result["status"] = statusMap[status]
	return 200, result
}

func (p PickUpServer) CancelOrder() (int, map[string]interface{}) {
	result := make(map[string]interface{})
	// 判断单号是否存在
	if model.HasOrder(p.ID) == 4002 {
		return 4002, result
	}
	data := model.PickUp{Status: 4}
	model.EditPickUp(p.ID, &data)
	id, status, address := model.GetPickUp(p.ID)
	result["id"] = id
	result["address"] = address
	result["status"] = statusMap[status]
	return 200, result
}
