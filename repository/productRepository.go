package repository

import (
	"context"
	"errors"

	"github.com/macadadi/e-shop/db"
	"github.com/macadadi/e-shop/form"
	"github.com/macadadi/e-shop/model"
)

const (
	allProducts = "SELECT id,name,price FROM products"
	deleteProduct = "DELETE * FROM products WHERE id= $1"
	findById =  allProducts + " WHERE id= $1"
	saveProduct = "INSERT INTO products (name, price, id) values ($1,$2,$3)"
	updateProduct = "UPDATE products SET name = $1, price =$2 WHERE id =$3"
)


type(
	ProductRepository interface{
		ListProducts(ctx context.Context,db db.DB)([]*model.Product,error)
		FindProductByID(ctx context.Context, db db.DB, id int64)(*model.Product, error)
		DeleteProduct(ctx context.Context, db db.DB,id int64)error
		UpdateProduct(ctx context.Context,db db.DB,form *form.ProductForm )(*model.Product,error)
		SaveProduct(ctx context.Context, db db.DB, form *form.ProductForm)error
	}
	AppProductRepository struct{}
)

func NewProductRepository()*AppProductRepository{
	return &AppProductRepository{}
}

func(s *AppProductRepository)ListProducts(ctx context.Context, db db.DB)([]*model.Product,error){
	rows, err := db.QueryContext(ctx ,allProducts)

	if err != nil{
		return  []*model.Product{},err
	}
	var products = make([]*model.Product, 0)

	defer rows.Close()
  	for rows.Next(){
		var prod model.Product
		err := rows.Scan(&prod.Id,
			&prod.Name,
			&prod.Price,
		)
		if err != nil {
			return []*model.Product{},err
		}
		products = append(products, &prod)
	}
	return products,nil
}

func(s *AppProductRepository)FindProductByID(ctx context.Context, db db.DB, id int64)(*model.Product,error){
	var product model.Product
	row := db.QueryRowContext(ctx,findById, id)
	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Price,
	)
	if err != nil{
		return &model.Product{}, errors.New(err.Error())
	}
	return &product,nil
}

func(s *AppProductRepository)DeleteProduct(ctx context.Context, db db.DB, id int64)error{
      _,err := s.FindProductByID(ctx,db, id)

	  if err != nil{
		return errors.New("procuct could not be found")
	  }
	  _,err = db.ExecContext(ctx,deleteProduct,id)
	  if err != nil{
		return errors.New("something went wrong")
	  }
	  return nil
}

func (s *AppProductRepository)UpdateProduct(ctx context.Context, db db.DB, form *form.ProductForm)(*model.Product,error){


	_, err := db.ExecContext(ctx,updateProduct,form.Name,form.Price,form.Id)
	if err != nil{
		return &model.Product{}, errors.New("could not update product")
	}
	product,err := s.FindProductByID(ctx,db,form.Id)
	
	if err != nil{
		return &model.Product{}, errors.New("could not find product")
	}
	return product,nil

}

func(s *AppProductRepository)SaveProduct(ctx context.Context,db db.DB,form *form.ProductForm)error{
     _,err := db.ExecContext(ctx, saveProduct, form.Name,form.Price,form.Id)
	 
	 if err != nil{
		return err
	 }

	 return nil
}