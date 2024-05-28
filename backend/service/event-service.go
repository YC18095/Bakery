package service

import (
	"backend/dto"
	"backend/entity"
	"backend/repository"
)

type EventService interface {
	List(uint32, int) *dto.DtoEventList
	Insert(itm *entity.Event) entity.Event
}

type eventService struct{}

func NewEventService() EventService {
	return &eventService{}
}

func (svc *eventService) List(pdtId uint32, lmt int) *dto.DtoEventList {
	evtRep := repository.NewEventRepository()
	evtLst := dto.DtoEventList{}
	evtLst.Total = evtRep.Count()
	evtLst.List = evtRep.List(pdtId, lmt)
	return &evtLst
}

func (svc *eventService) Insert(itm *entity.Event) entity.Event {
	evtRep := repository.NewEventRepository()
	return evtRep.Insert(itm)
}
