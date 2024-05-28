package controller

import (
	f "backend/function"
	"backend/response"
	"backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewsController interface {
	List(ctx *gin.Context)
	FindID(ctx *gin.Context)
}

type newsController struct{}

func NewNewsController() NewsController {
	return &newsController{}
}

func (ctrl *newsController) List(ctx *gin.Context) {
	lmt := f.GetParamInt(ctx.Param("limit"))
	pag := f.GetParamInt(ctx.Param("page"))

	svc := service.NewNewsService()
	itms := svc.List(lmt, pag)

	res := response.Build("success", itms)
	ctx.JSON(http.StatusOK, res)
}

func (ctrl *newsController) FindID(ctx *gin.Context) {
	id := f.GetParamUint(ctx.Param("id"))
	if id == 0 {
		res := response.Error("ID invalid", "")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	svc := service.NewNewsService()
	itm := svc.FindID(uint32(id))

	res := response.Build("success", itm)
	ctx.JSON(http.StatusOK, res)
}
