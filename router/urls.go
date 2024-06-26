package router

import (
	"ginTut/views"
	"ginTut/views/api"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	r.GET("/", views.ShowIndexPage)
	r.GET("/create", views.ShowCreatePage)
	r.POST("/create", views.CreateItem)
	r.GET("/update/:id", views.ShowUpdatePage)
	r.POST("/update", views.UpdateItem)
	r.GET("/delete/:id", views.DeleteItem)
	r.GET("/get-all-users", api.GetAllUsers)
}
