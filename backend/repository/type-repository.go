package repository

import (
	"backend/database"
	"backend/dto"
	"backend/entity"
	"backend/function"
	"log"

	"gorm.io/gorm"
)

type TypeRepository interface {
	List(int) *[]dto.DtoTypeItem
	AdminList(int, int) *[]dto.DtoTypeItem
	AdminFindID(uint32) *dto.DtoTypeItem
	Count() uint32
	AdminUpdate(*entity.Type) bool
	Insert(itm *entity.Type) entity.Type
}

type typeRepository struct {
	connection *gorm.DB
}

func NewTypeRepository() TypeRepository {
	return &typeRepository{
		connection: database.Connect(),
	}
}

func (rep *typeRepository) List(lmt int) *[]dto.DtoTypeItem {
	// query data from database and assigns into row
	rows, err := rep.connection.Table("type AS t").Select([]string{"t.id", "t.name"}).Where("t.status = 1").Order("id desc").Limit(lmt).Rows()
	// print error and exit the function
	if err != nil {
		log.Fatalf("Query product table error: %v\n", err)
	}
	defer rows.Close()
	// map rows data into DtoProductItem array and return
	res := function.MapRowsToDtoType(rows, 1)
	return res
}

func (rep *typeRepository) AdminList(lmt, pag int) *[]dto.DtoTypeItem {
	// query data from database and assigns into row
	rows, err := rep.connection.Table("type AS t").Select([]string{"t.id", "t.name", "t.status", "t.updated_at", "(SELECT COUNT(id) FROM product WHERE type_id = t.id AND status > 0)"}).Where("t.status > 0").Order("id desc").Limit(lmt).Rows()
	// print error and exit the function
	if err != nil {
		log.Fatalf("Query product table error: %v\n", err)
	}
	defer rows.Close()
	// map rows data into DtoProductItem array and return
	res := function.MapRowsToDtoType(rows, 2)
	return res
}

func (rep *typeRepository) AdminFindID(id uint32) *dto.DtoTypeItem {
	// query data from database and assigns into row
	rows, err := rep.connection.Table("type AS t").Select([]string{"t.id", "t.name", "t.status", "t.created_at", "t.updated_at", "(SELECT COUNT(id) FROM product WHERE type_id = t.id AND status > 0)"}).Where("t.status > 0 AND t.id = ?", id).Rows()
	// print error and exit the function
	if err != nil {
		log.Fatalf("Query product table error: %v\n", err)
	}
	defer rows.Close()
	// map rows data into DtoProductItem array and return
	res := function.MapRowsToDtoType(rows, 3)
	return &(*res)[0]
}

func (rep *typeRepository) Count() uint32 {
	var ttl int64
	// count the type and return
	rep.connection.Table("type AS t").Where("t.status = 1").Count(&ttl)
	return uint32(ttl)
}

func (rep *typeRepository) AdminUpdate(typ *entity.Type) bool {
	res := rep.connection.Updates(&typ)
	return res.RowsAffected > 0
}

func (rep *typeRepository) Insert(itm *entity.Type) entity.Type {
	rep.connection.Save(&itm)
	return *itm
}
