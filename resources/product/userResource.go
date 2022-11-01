package product

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/macadadi/e-shop/db"
	"github.com/macadadi/e-shop/form"
	"github.com/macadadi/e-shop/services"
)

func AddUser(db db.DB, s *services.UserService)func(c *gin.Context){
	return func(c *gin.Context) {
		var form *form.UserForm
		ctx := c.Request.Context()
		if err := c.BindJSON(&form); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":"You provided an invalid form"})
			return
		}
		form.Created_at = time.Now()
		if err := s.AddUser(ctx,db, form); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
     c.JSON(http.StatusCreated, gin.H{"message":"successfully added user"})
	}
}

func GetAllUsers(db db.DB, s *services.UserService)func(c *gin.Context){
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		users, err := s.GetAllUsers(ctx,db)

		if err != nil{
			c.JSON(http.StatusBadRequest,err)
			return
		}
		c.JSON(http.StatusOK,users)
	}
}