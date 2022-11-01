package services

import (
	"context"

	"github.com/macadadi/e-shop/db"
	"github.com/macadadi/e-shop/form"
	"github.com/macadadi/e-shop/model"
	"github.com/macadadi/e-shop/repository"
)

type(
	UserServiceInterface interface{
		AddUser(ctx context.Context, db db.DB, form *form.UserForm)error
		GetAllUsers(ctx context.Context, db db.DB)([]*model.User,error)
	}

	UserService struct{
		UserRepository *repository.UserRepository
	}
)

func NewUserService( repo *repository.UserRepository)*UserService{
	return &UserService{
		UserRepository: repo,
	}
}

func(s *UserService)AddUser(ctx context.Context, db db.DB, form *form.UserForm)(error){	
	err := s.UserRepository.AddUser(ctx, db, form)

	if err != nil{
		return err
	}
	return nil
}

func(s *UserService)GetAllUsers(ctx context.Context, db db.DB)([]*model.User,error){
	users, err := s.UserRepository.GetAllUsers(ctx,db)

	if err != nil{
		return []*model.User{},err
	}
	return users,nil
}
