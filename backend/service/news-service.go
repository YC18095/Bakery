package service

import (
	"backend/dto"
	"backend/entity"
	"backend/repository"
)

type NewsService interface {
	List(int, int) *dto.DtoNewsList
	FindID(id uint32) *dto.DtoNewsItem
	Read(uint32) entity.News
	Insert(itm *entity.News) entity.News
}

type newsService struct{}

func NewNewsService() NewsService {
	return &newsService{}
}

func (svc *newsService) List(lmt, pag int) *dto.DtoNewsList {
	evtRep := repository.NewNewsRepository()
	evtLst := dto.DtoNewsList{}
	evtLst.Total = evtRep.Count()
	evtLst.List = evtRep.List(lmt, pag)
	return &evtLst
}

func (svc *newsService) FindID(id uint32) *dto.DtoNewsItem {
	evtRep := repository.NewNewsRepository()
	evtItm := evtRep.FindID(id)
	evtItm.RelatedItems = evtRep.FindRelated(id)
	return evtItm
}

func (svc *newsService) Read(id uint32) entity.News {
	evtRep := repository.NewNewsRepository()
	return evtRep.Read(id)
}

func (svc *newsService) Insert(itm *entity.News) entity.News {
	evtRep := repository.NewNewsRepository()
	return evtRep.Insert(itm)
}
