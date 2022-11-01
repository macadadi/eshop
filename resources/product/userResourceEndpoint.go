package product

import (
	"github.com/gin-gonic/gin"
	"github.com/macadadi/e-shop/db"
	"github.com/macadadi/e-shop/services"
)

func UserEndpoints(route *gin.Engine,db db.DB, s *services.UserService){
	route.GET("/users", GetAllUsers(db, s))
	route.POST("/user", AddUser(db, s))

}