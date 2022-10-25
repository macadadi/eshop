package form

type ProductForm struct{
	Name string   `json:"name"`
	Price int64  	`json:"price"`
	Id   int64  	`json:"id"`
}