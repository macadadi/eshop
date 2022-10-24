package product

import (
	"context"

	"github.com/macadadi/e-shop/db"
	"github.com/macadadi/e-shop/model"
)

const (
	allProducts = "SELECT id,name,price FROM products"
	findById =  allProducts + "WHERE id= ?"
)


type(
	ProductRepository interface{
		ListProducts(ctx context.Context,db db.DB)([]*model.Product,error)
		FindProductByID(ctx context.Context, db db.DB, id int64)(*model.Product, error)
		DeleteProduct(ctx context.Context, db db.DB,id int64)error
		UpdateProduct(ctx context.Context,db db.DB, )(*model.Product,error)
	}
	AppProductRepository struct{}
)

func NewProductRepository()*AppProductRepository{
	return &AppProductRepository{}
}

func(s *AppProductRepository)ListProducts(ctx context.Context, db db.DB)([]*model.Product,error){
	rows, error := db.QueryContext(ctx ,allProducts)

	if error != nil{
		return  []*model.Product{},error
	}
	var products = make([]*model.Product, 0)
	defer rows.Close()
  	for rows.Next(){
		var product *model.Product
		error := rows.Scan(&product.Id,
			&product.Name,
			&product.Price,
		)
		if error != nil {
			return []*model.Product{},error
		}
		products = append(products, product)
	}
	return products,nil
}

func( s *AppProductRepository)FindProductByID(ctx context.Context, db db.DB, id int64)(*model.Product,error){
	var product *model.Product
	row := db.QueryRowContext(ctx,findById, id)
	error := row.Scan(
		&product.Id,
		&product.Name,
		&product.Price,
	)
	if error != nil{
		return &model.Product{},error
	}
	return product,nil
}