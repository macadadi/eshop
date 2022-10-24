package product

import (
	"github.com/gin-gonic/gin"
)

func Endpoints(route *gin.Engine){
	route.GET("/test",GetProduct())
	route.POST("/test",SaveProduct())
}