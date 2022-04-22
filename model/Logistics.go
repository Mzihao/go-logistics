package model

import (
	"go-logistics/utils/errmsg"
	"gorm.io/gorm"
)

type Logistics struct {
	gorm.Model
	ID             string `gorm:"primary_key;auto_increment" json:"id"`
	TrackingNumber string `gorm:"type:varchar(100)" json:"tracking_number"`
	CarrierCode    string `gorm:"type:varchar(100)" json:"carrier_code"`
}

// CreateLogistics 新增物流信息
func CreateLogistics(data *Logistics) int {
	err := DB.Create(&data).Error
	if err != nil {
		return errmsg.Error // 500
	}
	return errmsg.Success
}

// GetLogisticsById 查询单个分类信息by id
func GetLogisticsById(id string) (string, string) {
	var logistics Logistics
	result := DB.Where("id = ?", id).First(&logistics)
	if result.RowsAffected > 0 {
		return logistics.TrackingNumber, logistics.CarrierCode
	}
	return "", ""
}

// GetLogisticsByTrackingNumber 查询单个分类信息by id
func GetLogisticsByTrackingNumber(trackingNumber string, carrierCode string) int64 {
	var logistics Logistics
	//DB.Where("tracking_number = ?", trackingNumber).First(&logistics)
	result := DB.Where(&Logistics{TrackingNumber: trackingNumber, CarrierCode: carrierCode}).First(&logistics)
	//result.RowsAffected // 返回找到的记录数
	//result.Error        // returns error or nil
	return result.RowsAffected
}
