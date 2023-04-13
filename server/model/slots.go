package model

type Slot struct {
	ID              int    `json:"slotId" gorm:"primaryKey"`
	Start_time      string `json:"startTime"`
	End_time        string `json:"endTime"`
	Available_space int64  `json:"availableSlots" gorm:"default:50"`
}
