// apis.go

package api

import (
	"ginTut/middleware"
	"ginTut/models" // Adjust the import path
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllUsers fetches all users from the database
// func GetAllUsers(c *gin.Context) {
// 	var users []models.Item // Replace with your actual model name

// 	// Fetch all users from the database
// 	if err := models.DB.Find(&users).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, users)
// }

// GetAllUsers godoc
// @Summary Get all users API
// @Description Get All Users Details From the DataBase
// @Tags Users Data
// @Produce json
// @Success 200 {object} ResponseData
// @Failure 500 {object} ResponseData
// @Router /get-all-users [get]
func GetAllUsers(c *gin.Context) {
	var users []models.Item // Replace with your actual model name

	// Fetch all users from the database
	if err := models.DB.Find(&users).Error; err != nil {
		responseData := middleware.ResponseData{
			ResponseDetail: "Failed to fetch users",
			ResponseCode:   http.StatusInternalServerError,
		}
		responseData.Data.Error = err.Error() // Set the error message

		c.JSON(http.StatusInternalServerError, responseData)
		return
	}

	// If users are found, set them in the response data
	responseData := middleware.ResponseData{
		ResponseDetail: "OK",
		ResponseCode:   http.StatusOK,
	}
	for _, user := range users {
		responseData.Data.Result.Results = append(responseData.Data.Result.Results, user)
	}

	c.JSON(http.StatusOK, responseData)
}
