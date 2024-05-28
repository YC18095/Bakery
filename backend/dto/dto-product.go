package dto

type DtoProductList struct {
	Total uint32            `json:"total"`
	List  *[]DtoProductItem `json:"list"`
}

type DtoProductItem struct {
	ID           uint32            `json:"id"`
	Name         string            `json:"name"`
	Images       *[]string         `json:"images"`
	Content      string            `json:"content,omitempty"`
	Status       uint8             `json:"status,omitempty"`
	RelatedItems *[]DtoProductItem `json:"relatedItems,omitempty"`
	Type         *DtoTypeItem      `json:"type,omitempty"`
	Events       *[]DtoEventItem   `json:"events,omitempty"`
}

type DtoProductForm struct {
	ID     uint32 `json:"id,omitempty" form:"id"`
	Name   string `json:"name,omitempty" form:"name"`
	Status uint8  `json:"status,omitempty" form:"status"`
	TypeID uint32 `json:"typeId,omitempty" form:"typeId"`
}
