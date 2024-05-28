package function

import (
	"backend/dto"
	"database/sql"
	"log"
	"strings"
)

func MapRowsToDtoType(rows *sql.Rows, num uint8) *[]dto.DtoTypeItem {
	res := []dto.DtoTypeItem{}
	// assign every rows items into DtoTypeItem
	// then append into DtoTypeItem array and return
	for rows.Next() {
		itm := dto.DtoTypeItem{}
		var err error
		switch num {
		case 1:
			err = rows.Scan(&itm.ID, &itm.Name)
		case 2:
			err = rows.Scan(&itm.ID, &itm.Name, &itm.Status, &itm.UpdatedAt, &itm.ProductCount)
		case 3:
			err = rows.Scan(&itm.ID, &itm.Name, &itm.Status, &itm.CreatedAt, &itm.UpdatedAt, &itm.ProductCount)
		}
		if err != nil {
			log.Fatalf("Mapping rows into DtoProductItem error: %v\n", err)
		}
		res = append(res, itm)
	}
	return &res
}

func MapRowsToDtoEvent(rows *sql.Rows, num uint8) *[]dto.DtoEventItem {
	res := []dto.DtoEventItem{}
	// assign every rows items into DtoEventItem
	// then append into DtoEventItem array and return
	for rows.Next() {
		itm := dto.DtoEventItem{}
		var err error
		switch num {
		case 1:
			err = rows.Scan(&itm.ID, &itm.Name)
		}
		if err != nil {
			log.Fatalf("Mapping rows into DtoEventItem error: %v\n", err)
		}
		res = append(res, itm)
	}
	return &res
}

func MapRowsToDtoProduct(rows *sql.Rows, num uint8) *[]dto.DtoProductItem {
	res := []dto.DtoProductItem{}
	// assign every rows items into DtoProductItem
	// then append into DtoProductItem array and return
	for rows.Next() {
		itm := dto.DtoProductItem{}
		img := ""
		var err error
		switch num {
		// map object for product list
		case 1:
			err = rows.Scan(&itm.ID, &itm.Name, &img, &itm.Content)
		// map object for product detail
		case 2:
			itm.Type = &dto.DtoTypeItem{}
			err = rows.Scan(&itm.ID, &itm.Name, &img, &itm.Content, &itm.Type.ID, &itm.Type.Name)
		// map object for related products
		case 3:
			err = rows.Scan(&itm.ID, &itm.Name, &img)
		case 4:
			err = rows.Scan(&itm.ID, &itm.Name, &img, &itm.Status)
		}
		if err != nil {
			log.Fatalf("Mapping rows into DtoProductItem error: %v\n", err)
		}
		imgSpt := strings.Split(img, ",")
		itm.Images = &imgSpt
		res = append(res, itm)
	}
	return &res
}

func MapRowsToDtoNews(rows *sql.Rows, num uint8) *[]dto.DtoNewsItem {
	res := []dto.DtoNewsItem{}
	// assign every rows items into DtoNewsItem
	// then append into DtoNewsItem array and return
	for rows.Next() {
		itm := dto.DtoNewsItem{}
		var err error
		switch num {
		// map object for product list
		case 1:
			err = rows.Scan(&itm.ID, &itm.Name, &itm.Image, &itm.Content, &itm.CreatedAt)
		// map object for product detail
		case 2:
			itm.Event = &dto.DtoEventItem{}
			err = rows.Scan(&itm.ID, &itm.Name, &itm.Image, &itm.Content, &itm.CreatedAt, &itm.Event.ID, &itm.Event.Name)
		// map object for related products
		case 3:
			err = rows.Scan(&itm.ID, &itm.Name)
		}
		if err != nil {
			log.Fatalf("Mapping rows into DtoNewsItem error: %v\n", err)
		}
		res = append(res, itm)
	}
	return &res
}
