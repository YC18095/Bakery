package service

import (
	"backend/dto"
	"backend/entity"
	"backend/repository"
)

type ProductService interface {
	List(uint32, string, int, int, string) *dto.DtoProductList
	FindID(id uint32) *dto.DtoProductItem
	Read(uint32) *entity.Product
	AdminListByType(uint32, int, int) *dto.DtoProductList
	AdminUpdate(*dto.DtoProductForm) bool
	Insert(itm *entity.Product) entity.Product
}

type productService struct{}

func NewProductService() ProductService {
	return &productService{}
}

func (svc *productService) List(typId uint32, evtIds string, lmt, pag int, key string) *dto.DtoProductList {
	pdtRep := repository.NewProductRepository()
	pdtLst := dto.DtoProductList{}
	pdtLst.Total = pdtRep.Count(typId, evtIds, key)
	pdtLst.List = pdtRep.List(typId, evtIds, lmt, pag, key)
	return &pdtLst
}

func (svc *productService) FindID(id uint32) *dto.DtoProductItem {
	pdtRep := repository.NewProductRepository()
	evtRep := repository.NewEventRepository()
	pdtItm := pdtRep.FindID(id)
	pdtItm.RelatedItems = pdtRep.FindRelated(id)
	pdtItm.Events = evtRep.List(id, 0)
	return pdtItm
}

func (svc *productService) Read(id uint32) *entity.Product {
	pdtRep := repository.NewProductRepository()
	return pdtRep.Read(id)
}

func (svc *productService) AdminListByType(typId uint32, lmt, pag int) *dto.DtoProductList {
	pdtRep := repository.NewProductRepository()
	pdtLst := dto.DtoProductList{}
	pdtLst.Total = pdtRep.AdminCount("type", typId)
	pdtLst.List = pdtRep.AdminListByType(typId, lmt, pag)
	return &pdtLst
}

func (svc *productService) AdminUpdate(itm *dto.DtoProductForm) bool {
	pdt := entity.Product{ID: itm.ID}
	pdtRep := repository.NewProductRepository()
	return pdtRep.AdminUpdate(&pdt)
}

func (svc *productService) Insert(itm *entity.Product) entity.Product {
	pdtRep := repository.NewProductRepository()
	return pdtRep.Insert(itm)
}
