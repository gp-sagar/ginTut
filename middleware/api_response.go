package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Data struct {
		Result struct {
			Results []interface{} `json:"results"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"data"`
	ResponseDetail string `json:"responseDetail"`
	ResponseCode   int    `json:"responseCode"`
}

// JSONResponseMiddleware formats responses in the specified JSON format
func JSONResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a new response data structure
		responseData := ResponseData{
			ResponseDetail: "OK",
			ResponseCode:   http.StatusOK,
		}

		// Call the next handler
		c.Next()

		// Check if response is not yet written (e.g., if the route handler did not write a response)
		if !c.Writer.Written() {
			responseData.Data.Error = nil // Ensure no error data if not set
			c.JSON(http.StatusOK, responseData)
		} else {
			// Marshal the response body and set it in responseData
			var responseBody interface{}
			if err := c.ShouldBindJSON(&responseBody); err == nil {
				responseData.Data.Result.Results = append(responseData.Data.Result.Results, responseBody)
			}

			// Send the final response
			c.JSON(http.StatusOK, responseData)
		}
	}
}
