package controller

import (
	"backend/dto"
	f "backend/function"
	"backend/response"
	"backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TypeController interface {
	List(ctx *gin.Context)
	AdminList(ctx *gin.Context)
	AdminFindID(ctx *gin.Context)
	AdminUpdate(ctx *gin.Context)
	AdminRemove(ctx *gin.Context)
}

type typeController struct{}

func NewTypeController() TypeController {
	return &typeController{}
}

func (ctrl *typeController) List(ctx *gin.Context) {
	lmt := f.GetParamInt(ctx.Param("limit"))
	pdtLmt := f.GetParamInt(ctx.Param("productLimit"))

	typSvc := service.NewTypeService()
	itms := typSvc.List(lmt, pdtLmt)

	res := response.Build("success", itms)
	ctx.JSON(http.StatusOK, res)
}

func (ctrl *typeController) AdminList(ctx *gin.Context) {
	lmt := f.GetParamInt(ctx.Param("limit"))
	pag := f.GetParamInt(ctx.Param("page"))

	typSvc := service.NewTypeService()
	itms := typSvc.AdminList(lmt, pag)

	res := response.Build("success", itms)
	ctx.JSON(http.StatusOK, res)
}

func (ctrl *typeController) AdminFindID(ctx *gin.Context) {
	id := f.GetParamUint(ctx.Param("id"))
	if id == 0 {
		res := response.Error("ID invalid", "")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	typSvc := service.NewTypeService()
	itm := typSvc.AdminFindID(id)

	res := response.Build("success", itm)
	ctx.JSON(http.StatusOK, res)
}

func (ctrl *typeController) AdminUpdate(ctx *gin.Context) {
	var itm dto.DtoTypeUpdate
	err := ctx.ShouldBind(&itm)
	if err != nil {
		res := response.Error("Failed to update", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	typSvc := service.NewTypeService()
	isSuc := typSvc.AdminUpdate(&itm)

	if !isSuc {
		res := response.Error("Failed to update Type", "")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := response.Build("success", isSuc)
	ctx.JSON(http.StatusOK, res)
}

func (ctrl *typeController) AdminRemove(ctx *gin.Context) {
	var itm dto.DtoTypeRemove
	err := ctx.ShouldBind(&itm)
	if err != nil {
		res := response.Error("Failed to remove", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	typSvc := service.NewTypeService()
	isSuc := typSvc.AdminRemove(&itm)

	if !isSuc {
		res := response.Error("Failed to remove Type", "")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := response.Build("success", isSuc)
	ctx.JSON(http.StatusOK, res)
}
