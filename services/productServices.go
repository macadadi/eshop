package services

import (
	"context"

	"github.com/macadadi/e-shop/db"
	"github.com/macadadi/e-shop/form"
	"github.com/macadadi/e-shop/model"
	"github.com/macadadi/e-shop/repository"
)

type(
	ProductService interface{
		GetProduct(ctx context.Context,db db.DB)([]*model.Product,error)
		SaveProduct(ctx context.Context,db db.DB, form *form.ProductForm)error
		UpdateProduct(ctx context.Context, db db.DB, form *form.ProductForm)(*model.Product,error)
		GetSingleProduct(ctx context.Context, db db.DB, id int)(*model.Product,error)
	}
	AppProductService struct{
		Repository  *repository.AppProductRepository
	}
)

func NewProductService(r *repository.AppProductRepository)*AppProductService{
	return &AppProductService{
		Repository: r,
	}
}
func(p *AppProductService)GetProduct(ctx context.Context,db db.DB)([]*model.Product,error){
 
	prod, err := p.Repository.ListProducts(ctx, db)
	if err != nil{
		return []*model.Product{},err
	}
	return prod,nil
}

func(p *AppProductService)SaveProduct(ctx context.Context,db db.DB, form *form.ProductForm)(error){
 
	err := p.Repository.SaveProduct(ctx,db,form)
	if err != nil{
		return err
	}
	return nil
}

func (p *AppProductService)UpdateProduct(ctx context.Context, db db.DB, form *form.ProductForm)(*model.Product,error){

	prod,err := p.Repository.UpdateProduct(ctx, db, form)
	if err != nil{
		return &model.Product{},err
	}

	return prod,nil
}

func(s *AppProductService)GetSingleProduct(ctx context.Context, db db.DB, id int64)(*model.Product,error){
	prod, err := s.Repository.FindProductByID(ctx,db,id)

	if err != nil{
		return &model.Product{}, err
	}

	return prod,nil
}