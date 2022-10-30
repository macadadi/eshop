package form

type ProductForm struct{
	Name string   `json:"name" binding:"required"`
	Price int64  	`json:"price" binding:"required"`
	Id   int64  	`json:"id" `
}