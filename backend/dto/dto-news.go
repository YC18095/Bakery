package dto

import "time"

type DtoNewsList struct {
	Total uint32         `json:"total"`
	List  *[]DtoNewsItem `json:"list"`
}

type DtoNewsItem struct {
	ID           uint32         `json:"id"`
	Name         string         `json:"name"`
	Image        string         `json:"image,omitempty"`
	Content      string         `json:"content,omitempty"`
	Status       uint8          `json:"status,omitempty"`
	CreatedAt    time.Time      `json:"createdAt,omitempty"`
	RelatedItems *[]DtoNewsItem `json:"relatedItems,omitempty"`
	Event        *DtoEventItem  `json:"event,omitempty"`
}
