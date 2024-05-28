package entity

import "time"

type Type struct {
	ID        uint32     `gorm:"primary_key:auto_increment" json:"id" csv:"id"`
	Name      string     `gorm:"type:varchar(100)" json:"name" csv:"name"`
	Status    uint8      `json:"status,omitempty" csv:"status,omitempty"` // 0: deleted, 1: active, 2: unactive
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	Products  *[]Product `json:"products,omitempty"`
}

func (table *Type) TableName() string {
	return "type"
}
