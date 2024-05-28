package repository

import (
	"backend/database"
	"backend/dto"
	"backend/entity"
	"backend/function"
	"log"

	"gorm.io/gorm"
)

type NewsRepository interface {
	List(int, int) *[]dto.DtoNewsItem
	FindID(id uint32) *dto.DtoNewsItem
	FindRelated(id uint32) *[]dto.DtoNewsItem
	Count() uint32
	Read(uint32) entity.News
	Insert(itm *entity.News) entity.News
}

type newsRepository struct {
	connection *gorm.DB
}

func NewNewsRepository() NewsRepository {
	return &newsRepository{
		connection: database.Connect(),
	}
}

func (rep *newsRepository) List(lmt, pag int) *[]dto.DtoNewsItem {
	// query data from database and assigns into row
	rows, err := rep.connection.Table("news AS n").Select([]string{"n.id", "n.name", "n.image", "SUBSTRING(n.content, 1, 100)", "n.created_at"}).Joins("JOIN event AS e ON n.event_id = e.id").Where("n.status = 1 AND e.status = 1").Order("id desc").Offset(pag * lmt).Limit(lmt).Rows()
	// print error and exit the function
	if err != nil {
		log.Fatalf("Query type table error: %v\n", err)
	}
	defer rows.Close()
	// map rows data into DtoProductItem array and return
	res := function.MapRowsToDtoNews(rows, 1)
	return res
}

func (rep *newsRepository) FindID(id uint32) *dto.DtoNewsItem {
	rows, err := rep.connection.Table("news AS n").Select([]string{"n.id", "n.name", "n.image", "n.content", "n.created_at", "e.id", "e.name"}).Joins("JOIN event AS e ON n.event_id = e.id").Where("n.status = 1 AND e.status = 1 AND n.id = ?", id).Rows()
	// print error and exit the function
	if err != nil {
		log.Fatalf("Query type table error: %v\n", err)
	}
	defer rows.Close()
	// map rows data into DtoProductItem array and return
	res := function.MapRowsToDtoNews(rows, 2)
	return &(*res)[0]
}

func (rep *newsRepository) FindRelated(id uint32) *[]dto.DtoNewsItem {
	rows, err := rep.connection.Raw("SELECT n.id, n.name FROM news AS n WHERE n.status = 1 AND n.id IN((SELECT MAX(id) FROM news WHERE status = 1 AND id < ?),(SELECT MIN(id) FROM news WHERE status = 1 AND id > ?))", id, id).Rows()
	// print error and exit the function
	if err != nil {
		log.Fatalf("Query type table error: %v\n", err)
	}
	defer rows.Close()
	// map rows data into DtoProductItem array and return
	res := function.MapRowsToDtoNews(rows, 3)
	return res
}

func (rep *newsRepository) Count() uint32 {
	var cnt int64
	rep.connection.Table("news AS n").Where("n.status = 1").Count(&cnt)
	return uint32(cnt)
}

func (rep *newsRepository) Read(id uint32) entity.News {
	itm := entity.News{}
	rep.connection.Model(entity.News{}).Where("status = 1 AND id = ?", id).Find(&itm)
	return itm
}

func (rep *newsRepository) Insert(itm *entity.News) entity.News {
	rep.connection.Save(itm)
	return *itm
}
