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

type ProductRepository interface {
	List(uint32, string, int, int, string) *[]dto.DtoProductItem
	FindID(id uint32) *dto.DtoProductItem
	FindRelated(id uint32) *[]dto.DtoProductItem
	Count(uint32, string, string) uint32
	AdminListByType(uint32, int, int) *[]dto.DtoProductItem
	AdminCount(string, uint32) uint32
	Read(uint32) *entity.Product
	AdminUpdate(*entity.Product) bool
	Insert(itm *entity.Product) entity.Product
}

type productRepository struct {
	connection *gorm.DB
}

func NewProductRepository() ProductRepository {
	return &productRepository{
		connection: database.Connect(),
	}
}

func (rep *productRepository) List(typId uint32, evtIds string, lmt, pag int, key string) *[]dto.DtoProductItem {
	// set query condition query all type
	whr := "p.status = 1 AND t.status = 1 AND e.status = 1"
	// set query condition query with specific type
	if typId > 0 {
		whr += " AND t.id = " + strconv.Itoa(int(typId))
	}
	// set query condition query with specific event(s)
	if len(evtIds) > 0 {
		whr += " AND e.id IN(" + evtIds + ")"
	}
	// set query condition query with like
	if key != "" {
		whr += " AND p.name LIKE '%" + key + "%'"
	}
	// query data from database and assigns into row
	rows, err := rep.connection.Table("product AS p").Select([]string{"p.id", "p.name", "SUBSTRING_INDEX(p.images,',',1)", "SUBSTRING(p.content,1,100)"}).Joins("JOIN type AS t ON p.type_id = t.id JOIN product_event AS pe ON p.id = pe.product_id JOIN event AS e ON e.id = pe.event_id").Where(whr).Order("id desc").Offset(pag * lmt).Limit(lmt).Group("p.id").Rows()
	// print error and exit the function
	if err != nil {
		log.Fatalf("Query product table error: %v\n", err)
	}
	defer rows.Close()
	// map rows data into DtoProductItem array and return
	res := function.MapRowsToDtoProduct(rows, 1)
	return res
}

func (rep *productRepository) FindID(id uint32) *dto.DtoProductItem {
	// query data from database and assigns into row
	// find by unique id, so will get only 1 row
	rows, err := rep.connection.Table("product AS p").Select([]string{"p.id", "p.name", "p.images", "p.content", "t.id", "t.name"}).Joins("JOIN type AS t ON p.type_id = t.id").Where("p.status = 1 AND t.status = 1 AND p.id = ?", id).Rows()
	// print error and exit the function
	if err != nil {
		log.Fatalf("Query product table error: %v\n", err)
	}
	defer rows.Close()
	// map rows data into DtoProductItem and return
	res := function.MapRowsToDtoProduct(rows, 2)
	return &(*res)[0]
}

func (rep *productRepository) FindRelated(id uint32) *[]dto.DtoProductItem {
	subQue := "SELECT id, name, SUBSTRING_INDEX(images,',',1) FROM product WHERE status = 1"
	rows, err := rep.connection.Raw("SELECT * FROM (("+subQue+" AND id < ? ORDER BY id DESC LIMIT 3) UNION ("+subQue+" AND id > ? ORDER BY id ASC LIMIT 3)) AS p ORDER BY p.id ASC LIMIT 3", id, id).Rows()
	// print error and exit the function
	if err != nil {
		log.Fatalf("Query type table error: %v\n", err)
	}
	defer rows.Close()
	// map rows data into DtoProductItem array and return
	res := function.MapRowsToDtoProduct(rows, 3)
	return res
}

func (rep *productRepository) Count(typId uint32, evtIds string, key string) uint32 {
	var ttl int64
	// set query condition query all type
	whr := "p.status = 1 AND t.status = 1 AND e.status = 1"
	// set query condition query with specific type
	if typId > 0 {
		whr += " AND p.type_id = " + strconv.Itoa(int(typId))
	}
	// set query condition query with specific event(s)
	if len(evtIds) > 0 {
		whr += " AND e.id IN(" + evtIds + ")"
	}
	// set query condition query with like
	if key != "" {
		whr += " AND p.name LIKE '%" + key + "%'"
	}
	// count the product and return
	rep.connection.Table("product AS p").Joins("JOIN type AS t ON p.type_id = t.id JOIN product_event AS pe ON p.id = pe.product_id JOIN event AS e ON e.id = pe.event_id").Where(whr).Group("p.id").Count(&ttl)
	return uint32(ttl)
}

// admin
func (rep *productRepository) AdminListByType(typId uint32, lmt, pag int) *[]dto.DtoProductItem {
	// query data from database and assigns into row
	rows, err := rep.connection.Table("product AS p").Select([]string{"p.id", "p.name", "p.images", "p.status"}).Joins("JOIN type AS t ON p.type_id = t.id").Where("p.status > 0 AND t.status > 0 AND t.id = ?", typId).Order("id desc").Offset(pag * lmt).Limit(lmt).Group("p.id").Rows()
	// print error and exit the function
	if err != nil {
		log.Fatalf("Query product table error: %v\n", err)
	}
	defer rows.Close()
	// map rows data into DtoProductItem array and return
	res := function.MapRowsToDtoProduct(rows, 4)
	return res
}

func (rep *productRepository) AdminCount(ent string, entId uint32) uint32 {
	var ttl int64
	whr, jon := "p.status > 0", ""
	// set query condition query with switch
	switch ent {
	case "type":
		whr += " AND p.type_id = " + strconv.Itoa(int(entId))
		jon += "JOIN type AS t ON p.type_id = t.id"
	case "event":
		whr += " AND p.event_id = " + strconv.Itoa(int(entId))
		jon += "JOIN product_event AS pe ON p.id = pe.product_id JOIN event AS e ON e.id = pe.event_id"
	}
	// count the product and return
	rep.connection.Table("product AS p").Joins(jon).Where(whr).Group("p.id").Count(&ttl)
	return uint32(ttl)
}

func (rep *productRepository) Read(id uint32) *entity.Product {
	itm := entity.Product{}
	rep.connection.Where("status = 1 AND id = ?", id).Take(&itm)
	return &itm
}

func (rep *productRepository) AdminUpdate(pdt *entity.Product) bool {
	res := rep.connection.Updates(&pdt)
	return res.RowsAffected > 0
}

func (rep *productRepository) Insert(itm *entity.Product) entity.Product {
	rep.connection.Save(&itm)
	return *itm
}
