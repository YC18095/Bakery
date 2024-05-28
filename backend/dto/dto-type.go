package dto

import "time"

type DtoTypeList struct {
	Total uint32         `json:"total,omitempty"`
	List  *[]DtoTypeItem `json:"list"`
}

type DtoTypeItem struct {
	ID           uint32            `json:"id"`
	Name         string            `json:"name"`
	Status       uint8             `json:"status,omitempty"`
	CreatedAt    time.Time         `json:"createdAt,omitempty"`
	UpdatedAt    time.Time         `json:"updatedAt,omitempty"`
	ProductCount uint32            `json:"productCount,omitempty"`
	Products     *[]DtoProductItem `json:"products,omitempty"`
}

type DtoTypeUpdate struct {
	ID     uint32 `json:"id,omitempty" form:"id" binding:"required"`
	Name   string `json:"name" form:"name" binding:"required"`
	Status uint8  `json:"status,omitempty" form:"status" binding:"required"`
}

type DtoTypeRemove struct {
	ID uint32 `json:"id,omitempty" form:"id"`
}
