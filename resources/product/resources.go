package product

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/macadadi/e-shop/db"
	"github.com/macadadi/e-shop/form"
	"github.com/macadadi/e-shop/services"
)

type User struct{
	Name string `json:"name"`
	Age int       `json:"age"`
}
func GetProduct(db db.DB, s *services.AppProductService)func(c *gin.Context){
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		prod, err := s.GetProduct(ctx, db)

		if err != nil{
			c.JSON(http.StatusBadRequest,gin.H{"err_message":err})
			return	
		}
		
		c.JSON(http.StatusOK, prod)
	}
}

func SaveProduct(db db.DB, s *services.AppProductService)func(c *gin.Context){
	return func(c *gin.Context) {
		var form *form.ProductForm
		userId, err := strconv.ParseInt(c.Param("id"),32,64)
		if err != nil{
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
ctx := c.Request.Context()
		err = c.BindJSON(&form)
		form.User_id = userId
		if err != nil{
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		err = s.SaveProduct(ctx, db, form)

		if err != nil{
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"message":"Successfully added product"})
}
}

func UpdateProduct(db db.DB, s *services.AppProductService)func(c *gin.Context){
	return func(c *gin.Context) {
		var form *form.ProductForm
		ctx := c.Request.Context()
		
		if err := c.BindJSON(&form); err != nil{
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		prod, err := s.UpdateProduct(ctx,db,form)

		if err != nil{
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusAccepted,prod)
	}
}

func GetSingleProduct(db db.DB, s *services.AppProductService)func(c *gin.Context){
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		id, err:= strconv.ParseInt(c.Param("id"),0,64)
		if err != nil{
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		prod, err := s.GetSingleProduct(ctx, db, id)
		if err != nil{
			c.JSON(http.StatusBadRequest,err)
			return
		}
		c.JSON(http.StatusOK,prod)
	}
}