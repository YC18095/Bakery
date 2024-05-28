package dto

type DtoEventList struct {
	Total uint32          `json:"total"`
	List  *[]DtoEventItem `json:"list"`
}

type DtoEventItem struct {
	ID     uint32 `json:"id"`
	Name   string `json:"name"`
	Status uint8  `json:"status,omitempty"`
}
