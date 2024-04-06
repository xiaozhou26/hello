package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	http.Handle("/", router)
}

func Listen(w http.ResponseWriter, r *http.Request) {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	router.ServeHTTP(w, r)
}
