package service

import (
	"backend/dto"
	"backend/entity"
	"backend/repository"
)

type TypeService interface {
	List(int, int) *dto.DtoTypeList
	AdminList(int, int) *dto.DtoTypeList
	AdminFindID(uint32) *dto.DtoTypeItem
	AdminUpdate(*dto.DtoTypeUpdate) bool
	AdminRemove(*dto.DtoTypeRemove) bool
	Insert(itm *entity.Type) entity.Type
}

type typeService struct{}

func NewTypeService() TypeService {
	return &typeService{}
}

func (svc *typeService) List(lmt, pdtLmt int) *dto.DtoTypeList {
	typRep := repository.NewTypeRepository()
	typLst := dto.DtoTypeList{}
	typLst.Total = typRep.Count()
	typLst.List = typRep.List(lmt)
	if pdtLmt > 0 {
		typLstLen := len(*typLst.List)
		pdtRep := repository.NewProductRepository()
		for i := 0; i < typLstLen; i++ {
			itm := &(*typLst.List)[i]
			itm.Products = pdtRep.List((*itm).ID, "", pdtLmt, 0, "")
		}
	}
	return &typLst
}

func (svc *typeService) AdminList(lmt, pag int) *dto.DtoTypeList {
	typRep := repository.NewTypeRepository()
	typLst := dto.DtoTypeList{}
	typLst.List = typRep.AdminList(lmt, pag)
	return &typLst
}

func (svc *typeService) AdminFindID(id uint32) *dto.DtoTypeItem {
	typRep := repository.NewTypeRepository()
	typItm := typRep.AdminFindID(id)
	return typItm
}

func (svc *typeService) AdminUpdate(itm *dto.DtoTypeUpdate) bool {
	typ := entity.Type{ID: itm.ID, Name: itm.Name, Status: itm.Status}
	typRep := repository.NewTypeRepository()
	return typRep.AdminUpdate(&typ)
}

func (svc *typeService) AdminRemove(itm *dto.DtoTypeRemove) bool {
	typ := entity.Type{ID: itm.ID}
	typRep := repository.NewTypeRepository()
	return typRep.AdminUpdate(&typ)
}

func (svc *typeService) Insert(itm *entity.Type) entity.Type {
	typRep := repository.NewTypeRepository()
	return typRep.Insert(itm)
}
