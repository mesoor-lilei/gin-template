package model

import "time"

type Model struct {
	Id        uint64 `gorm:"primaryKey"`
	CreatedBy uint64
	CreatedAt time.Time
	UpdatedBy uint64
	UpdatedAt time.Time
}
