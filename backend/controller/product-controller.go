package controller

import (
	"backend/dto"
	"backend/response"
	"backend/service"
	"bytes"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	f "backend/function"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	List(ctx *gin.Context)
	FindID(ctx *gin.Context)
	AdminListByType(ctx *gin.Context)
	AdminUpdate(ctx *gin.Context)
	Insert(ctx *gin.Context)
}

type productController struct{}

func NewProductController() ProductController {
	return &productController{}
}

func (ctrl *productController) List(ctx *gin.Context) {
	typId := f.GetParamUint(ctx.Param("typeId"))
	evtIds := f.GetParamUintArray(ctx.Param("eventIds"))
	lmt := f.GetParamInt(ctx.Param("limit"))
	pag := f.GetParamInt(ctx.Param("page"))
	key := ctx.Param("keyword")

	svc := service.NewProductService()
	itms := svc.List(typId, evtIds, lmt, pag, key)

	res := response.Build("success", itms)
	ctx.JSON(http.StatusOK, res)
}

func (ctrl *productController) FindID(ctx *gin.Context) {
	id := f.GetParamUint(ctx.Param("id"))
	if id == 0 {
		res := response.Error("ID invalid", "")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	svc := service.NewProductService()
	itm := svc.FindID(uint32(id))

	res := response.Build("success", itm)
	ctx.JSON(http.StatusOK, res)
}

func (ctrl *productController) AdminListByType(ctx *gin.Context) {
	typId := f.GetParamUint(ctx.Param("typeId"))
	lmt := f.GetParamInt(ctx.Param("limit"))
	pag := f.GetParamInt(ctx.Param("page"))
	if typId == 0 {
		res := response.Error("ID invalid", "")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	svc := service.NewProductService()
	itms := svc.AdminListByType(uint32(typId), lmt, pag)

	res := response.Build("success", itms)
	ctx.JSON(http.StatusOK, res)
}

func (ctrl *productController) AdminUpdate(ctx *gin.Context) {
	var itm dto.DtoProductForm
	err := ctx.ShouldBind(&itm)
	if err != nil {
		res := response.Error("Failed to update", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	pdtSvc := service.NewProductService()
	isSuc := pdtSvc.AdminUpdate(&itm)

	if !isSuc {
		res := response.Error("Failed to update product", "")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := response.Build("success", isSuc)
	ctx.JSON(http.StatusOK, res)
}

func (ctrl *productController) Insert(ctx *gin.Context) {
	uploadImage(ctx)
}

func uploadImage(ctx *gin.Context) {
	// parse our multipart form, 5 << 20 specifies a maximum upload of 5 MB files
	ctx.Request.ParseMultipartForm(5 << 20)
	// getting image file
	file, hdr, err := ctx.Request.FormFile("image")
	if err != nil {
		response := response.Error("Somethings went wrong when opening file", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	defer file.Close()

	extension := filepath.Ext(hdr.Filename)
	var buf, bufSnd bytes.Buffer

	switch extension {
	case ".png":
		img, err := png.Decode(file)
		if err != nil {
			response := response.Error("Somethings went wrong when converting file", err.Error())
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		png.Encode(&buf, img)
		// resize image
		imgSnd := imaging.Resize(img, 800, 0, imaging.Lanczos)
		png.Encode(&bufSnd, imgSnd)
	case ".jpg", ".jpeg":
		img, err := jpeg.Decode(file)
		if err != nil {
			response := response.Error("Somethings went wrong when converting file", err.Error())
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		jpeg.Encode(&buf, img, &jpeg.Options{Quality: 70})
		// resize image
		imgSnd := imaging.CropAnchor(img, 800, 800, imaging.Center)
		jpeg.Encode(&bufSnd, imgSnd, &jpeg.Options{Quality: 70})
	default:
		response := response.Error("Please input the image file", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	// get extension, microsecond to rename the filename
	micro := strconv.FormatInt(time.Now().UnixNano()/int64(time.Microsecond), 10)

	os.WriteFile("static/images/"+micro+extension, buf.Bytes(), 5<<20)
	os.WriteFile("static/images/"+micro+"-800"+extension, bufSnd.Bytes(), 5<<20)
}
