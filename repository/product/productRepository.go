package product

import (
	"context"
	"errors"

	"github.com/macadadi/e-shop/db"
	"github.com/macadadi/e-shop/form"
	"github.com/macadadi/e-shop/model"
)

const (
	allProducts = "SELECT id,name,price FROM products"
	deleteProduct = "DELETE * FROM products WHERE id= ?"
	findById =  allProducts + "WHERE id= ?"
	saveProduct = "INSERT INTO products(name, price) values (?,?)"
	updateProduct = "UPDATE products SET name = ?, price =? WHERE id =?"
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

func(s *AppProductRepository)FindProductByID(ctx context.Context, db db.DB, id int64)(*model.Product,error){
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
     _,err := db.ExecContext(ctx, saveProduct, form.Name,form.Price)
	 
	 if err != nil{
		return err
	 }

	 return nil
}