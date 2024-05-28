package repository

import (
	"backend/database"
	"backend/dto"
	"backend/entity"
	"backend/function"
	"log"
	"strconv"

	"gorm.io/gorm"
)

type EventRepository interface {
	List(uint32, int) *[]dto.DtoEventItem
	Count() uint32
	Insert(itm *entity.Event) entity.Event
}

type eventRepository struct {
	connection *gorm.DB
}

func NewEventRepository() EventRepository {
	return &eventRepository{
		connection: database.Connect(),
	}
}

func (rep *eventRepository) List(pdtId uint32, lmt int) *[]dto.DtoEventItem {
	whr := "e.status = 1"
	if pdtId > 0 {
		whr += " AND p.status = 1 AND p.id = " + strconv.Itoa(int(pdtId))
	}
	// query data from database and assigns into row
	rows, err := rep.connection.Table("event AS e").Select([]string{"e.id", "e.name"}).Joins("JOIN product_event AS pe ON e.id = pe.event_id JOIN product AS p ON p.id = pe.product_id").Where(whr).Order("id desc").Limit(lmt).Group("e.id").Rows()
	// print error and exit the function
	if err != nil {
		log.Fatalf("Query type table error: %v\n", err)
	}
	defer rows.Close()
	// map rows data into DtoProductItem array and return
	res := function.MapRowsToDtoEvent(rows, 1)
	return res
}

func (rep *eventRepository) Count() uint32 {
	var cnt int64
	rep.connection.Table("event AS e").Where("e.status = 1").Count(&cnt)
	return uint32(cnt)
}

func (rep *eventRepository) Insert(itm *entity.Event) entity.Event {
	rep.connection.Save(itm)
	return *itm
}
