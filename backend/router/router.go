package router

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	ClientRouter()
	AdminRouter()
}

type router struct {
	engine *gin.Engine
}

func NewRouter(ngin *gin.Engine) Router {
	return &router{
		engine: ngin,
	}
}

func (rtr *router) ClientRouter() {
	ping(rtr)
	setType(rtr)
	setEvent(rtr)
	setProduct(rtr)
	setNews(rtr)
}

func (rtr *router) AdminRouter() {
	migrateDate(rtr)
	// setAdminType(rtr)
	// setAdminProduct(rtr)
}
