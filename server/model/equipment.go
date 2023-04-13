package model

import (
	"time"

	"gorm.io/gorm"
)

type Equipment struct {
	Model_No   string         `json:"modelNo" gorm:"default:uuid_generate_v4();unique;primaryKey"` //PK
	Equip_Name string         `json:"equipName"`
	Quantity   int64          `json:"quantity"`
	Weight     float64        `json:"weight"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
