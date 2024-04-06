package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 设置路由，只响应GET请求，并在访问根URL时返回"hello"
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello")
	})
	// 启动服务器，监听默认的8080端口
	router.Run() // 默认监听在0.0.0.0:8080
}
