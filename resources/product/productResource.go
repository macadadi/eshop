package product

import (
	"github.com/gin-gonic/gin"
	"github.com/macadadi/e-shop/db"
	"github.com/macadadi/e-shop/services"
)

func Endpoints(route *gin.Engine, db db.DB, s *services.AppProductService){
	route.GET("/product",GetProduct(db,s))
	route.POST("/product",SaveProduct(db,s))
	route.PUT("/product",UpdateProduct(db,s))
	route.GET("/product/:id",GetSingleProduct(db,s))
}