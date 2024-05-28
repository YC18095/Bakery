package entity

import "time"

type Product struct {
	ID        uint32    `gorm:"primary_key:auto_increment" json:"id" csv:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name" csv:"name"`
	Images    string    `gorm:"type:text" json:"images,omitempty" csv:"images"`
	Content   string    `gorm:"type:text" json:"content" csv:"content,omitempty"`
	Status    uint8     `json:"status,omitempty" csv:"status,omitempty"` // 0: deleted, 1: active, 2: unactive
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	TypeID    uint32    `json:"typeId,omitempty" csv:"typeId"`
	Events    *[]Event  `gorm:"many2many:product_event" json:"events,omitempty"`
}

func (table *Product) TableName() string {
	return "product"
}
