package product

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct{
	Name string `json:"name"`
	Age int       `json:"age"`
}
func GetProduct()func(c *gin.Context){
	return func(c *gin.Context) {
	
		
		c.JSON(http.StatusOK, gin.H{"message":"Adadi f sammy new  the value is here test all"})
	}
}

func SaveProduct()func(c *gin.Context){
	return func(c *gin.Context) {
		
		data,_ := ioutil.ReadAll(c.Request.Body)
		fmt.Print(string(data))


	}
}