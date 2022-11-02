package form

import "time"

type ProductForm struct{
	Id   int64  	`json:"id" `
	Name string   `json:"name" binding:"required"`
	Price int64  	`json:"price" binding:"required"`
	User_id int64  	`json:"user_id"`
}

type UserForm  struct{
	Id   int64     `json:"id"`
	Full_name  string `json:"full_name" binding:"required"`
	Created_at  time.Time 
	Country_code int  `json:"country_code" binding:"required"`
}