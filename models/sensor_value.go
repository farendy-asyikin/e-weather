package models

import "time"

type SensorValue struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	SensorID  uint64    `gorm:"not null" json:"sensor_id"`
	Sensor    Sensor    `gorm:"foreignkey:SensorID" json:"sensor"`
	Value     float64   `gorm:"not null" json:"value"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}
