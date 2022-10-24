package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/macadadi/e-shop/resources/product"
)
func main(){

	rout := gin.New()
	rout.Use(gin.Logger())


	product.Endpoints(rout)

	rout.NoRoute(func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusNotFound,gin.H{"err_message":"end point not found test now 789 found keep trying dcdcd"})
	})
	server := &http.Server{
		Addr: "localhost:8080",
		Handler: rout,
		
	}
	log.Fatal(server.ListenAndServe())
}