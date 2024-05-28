package controller

import (
	f "backend/function"
	"backend/response"
	"backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventController interface {
	List(ctx *gin.Context)
}

type eventController struct{}

func NewEventController() EventController {
	return &eventController{}
}

func (ctrl *eventController) List(ctx *gin.Context) {
	pdtId := f.GetParamUint(ctx.Param("productId"))
	lmt := f.GetParamInt(ctx.Param("limit"))

	svc := service.NewEventService()
	itms := svc.List(pdtId, lmt)

	res := response.Build("success", itms)
	ctx.JSON(http.StatusOK, res)
}
