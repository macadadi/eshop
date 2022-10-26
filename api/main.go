package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/macadadi/e-shop/db"
	"github.com/macadadi/e-shop/repository"
	"github.com/macadadi/e-shop/resources/product"
	"github.com/macadadi/e-shop/services"
)
const Port =":3001"
func main(){

   	db:= db.InitDB()
	defer db.Close()

	route := gin.New()
	route.Use(gin.Logger())
	
	productRepository := repository.NewProductRepository()

	productService := services.NewProductService(productRepository)

	product.Endpoints(route,db,productService)

	route.NoRoute(func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusNotFound,gin.H{"err_message":"end point not found try again later "})
	})

	server := http.Server{
		Addr: Port,
		Handler: route,
	}
	
	log.Fatal(server.ListenAndServe())
}