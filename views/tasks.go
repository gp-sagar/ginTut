package views

import (
	"ginTut/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ShowIndexPage(c *gin.Context) {
	var items []models.Item
	models.DB.Find(&items)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Welcome Page",
		"Items": items,
	})
}

func ShowCreatePage(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", gin.H{
		"title": "Create User",
	})
}

func CreateItem(c *gin.Context) {
	name := c.PostForm("name")
	item := models.Item{Name: name}
	models.DB.Create(&item)
	c.Redirect(http.StatusMovedPermanently, "/")
}

func ShowUpdatePage(c *gin.Context) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid UUID")
		return
	}
	var item models.Item
	if err := models.DB.First(&item, "id = ?", uuid).Error; err != nil {
		c.String(http.StatusNotFound, "Item not found")
		return
	}
	c.HTML(http.StatusOK, "update.html", gin.H{
		"ID":    item.ID,
		"Name":  item.Name,
		"title": "Update User",
	})
}

func UpdateItem(c *gin.Context) {
	id := c.PostForm("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid UUID")
		return
	}
	name := c.PostForm("name")
	var item models.Item
	if err := models.DB.First(&item, "id = ?", uuid).Error; err != nil {
		c.String(http.StatusNotFound, "Item not found")
		return
	}
	item.Name = name
	models.DB.Save(&item)
	c.Redirect(http.StatusMovedPermanently, "/")
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid UUID")
		return
	}
	models.DB.Delete(&models.Item{}, "id = ?", uuid)
	c.Redirect(http.StatusMovedPermanently, "/")
}
