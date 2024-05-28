package router

import (
	"backend/controller"
)

func migrateDate(rtr *router) {
	ctrl := controller.NewDataController()
	ctrl.Start()
}

func setAdminType(rtr *router) {
	typRtr := rtr.engine.Group("api/admin/type")
	{
		typCtrl := controller.NewTypeController()
		typRtr.GET("/list/:limit/:page", typCtrl.AdminList)
		typRtr.GET("/detail/:id", typCtrl.AdminFindID)
		typRtr.POST("/edit", typCtrl.AdminUpdate)
		typRtr.POST("/remove", typCtrl.AdminRemove)
	}
}

func setAdminProduct(rtr *router) {
	pdtRtr := rtr.engine.Group("api/admin/product")
	{
		pdtCtrl := controller.NewProductController()
		pdtRtr.GET("/listbytype/:typeId/:limit/:page", pdtCtrl.AdminListByType)
		pdtRtr.POST("/edit", pdtCtrl.AdminUpdate)
		pdtRtr.POST("/upload", pdtCtrl.Insert)
	}
}
