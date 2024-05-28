package entity

import "time"

type News struct {
	ID        uint32    `gorm:"primary_key:auto_increment" json:"id" csv:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name" csv:"name"`
	Image     string    `gorm:"type:varchar(100)" json:"image" csv:"image"`
	Content   string    `gorm:"type:text" json:"content" csv:"content,omitempty"`
	Status    uint8     `json:"status,omitempty" csv:"status,omitempty"` // 0: deleted, 1: active, 2: unactive
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	EventID   uint32    `json:"eventId,omitempty" csv:"eventId"`
}

func (table *News) TableName() string {
	return "news"
}
