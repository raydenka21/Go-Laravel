package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"goravel/app/http/controllers"
	"goravel/setting"
)

func Api() {

	rdb := setting.InitRedis()
	categoryController := controllers.NewCategoryController(rdb)
	facades.Route().Prefix("category").Group(func(router route.Router) {
		router.Get("/", categoryController.Index)
		router.Post("/store", categoryController.Store)
		router.Get("/{id}", categoryController.Show)
		router.Put("/update/{id}", categoryController.Update)
		router.Delete("/destroy/{id}", categoryController.Destroy)
	})

}
