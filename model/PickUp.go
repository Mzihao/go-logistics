package model

import (
	"go-logistics/utils/errmsg"
	"gorm.io/gorm"
)

type PickUp struct {
	gorm.Model
	ID      string `gorm:"primary_key;auto_increment" json:"id"`
	Status  uint   `gorm:"type:int" json:"status"`
	Address string `gorm:"type:varchar(500)" json:"address"`
}

// CreatePickUp 新增自提订单
func CreatePickUp(data *PickUp) int {
	err := DB.Create(&data).Error
	if err != nil {
		return errmsg.Error // 500
	}
	return errmsg.Success
}

// GetPickUp 查询单个物流信息
func GetPickUp(id string) (string, uint, string) {
	var pickUp PickUp
	result := DB.Where("id = ?", id).First(&pickUp)
	if result.RowsAffected > 0 {
		return pickUp.ID, pickUp.Status, pickUp.Address
	}
	return "", 0, ""
}

// EditPickUp 编辑订单信息
func EditPickUp(id string, data *PickUp) int {
	var pick PickUp
	var maps = make(map[string]interface{})
	if data.Status != 0 {
		maps["status"] = data.Status
	}
	if data.Address != "" {
		maps["address"] = data.Address
	}
	err := DB.Model(&pick).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.Error
	}
	return errmsg.Success
}

// 查询订单号是否存在
func HasOrder(id string) int {
	var pickUp PickUp
	DB.Select("id").Where("id = ?", id).First(&pickUp)
	if pickUp.ID > 0 {
		return errmsg.NotFound //4002
	}
	return errmsg.Success
}