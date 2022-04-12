package model

import (
	"go-logistics/utils/errmsg"
	"gorm.io/gorm"
)

type Logistics struct {
	gorm.Model
	ID             uint   `gorm:"primary_key;auto_increment" json:"id"`
	TrackingNumber string `gorm:"type:varchar(100)" json:"tracking_number"`
	CarrierCode    string `gorm:"type:varchar(100)" json:"carrier_code"`
}

// CreateLogistics 新增物流信息
func CreateLogistics(data *Logistics) int {
	err := DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS
}

// GetLogisticsById 查询单个分类信息by id
func GetLogisticsById(id int) (Logistics, int) {
	var logistics Logistics
	DB.Where("id = ?", id).First(&logistics)
	return logistics, errmsg.SUCCESS
}

// GetLogisticsByTrackingNumber 查询单个分类信息by id
func GetLogisticsByTrackingNumber(trackingNumber string) (Logistics, int) {
	var logistics Logistics
	DB.Where("tracking_number = ?", trackingNumber).First(&logistics)
	return logistics, errmsg.SUCCESS
}
