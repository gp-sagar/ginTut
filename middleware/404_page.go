package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NotFoundHandler handles 404 errors by serving a custom HTML page
func NotFoundHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.HTML(http.StatusNotFound, "404.html", gin.H{})
    }
}
