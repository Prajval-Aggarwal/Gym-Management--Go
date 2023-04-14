package model

type Slot struct {
	Slot_Id         int    `json:"slotId" gorm:"primaryKey"`
	Start_time      string `json:"startTime"`
	End_time        string `json:"endTime"`
	Available_space int64  `json:"availableSlots" gorm:"default:50"`
}
