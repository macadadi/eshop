package db

import (
	"context"
	"database/sql"
	"log"
	"os"
)

type(
	DB interface{
		Close()error
		Ping()error
		ExecContext(ctx context.Context, query string, args ...interface{})(sql.Result,error)
		QueryContext(ctx context.Context,query string, args ...interface{})(*sql.Rows,error)
		QueryRowContext(ctx context.Context, query string, args ...interface{})*sql.Row
	}

	appDatabase struct{
		*sql.DB
	}
)

func InitDB()DB{
	url := os.Getenv("DATABASE_URL")
	if url == ""{
		log.Fatal("No url in environment")
	}
	return initDbWithUrl(url)
}

func initDbWithUrl(s string)DB{
	appDB := initAppDB(s)
	db := &appDatabase{
		DB: appDB,
	}
	return db
}

func initAppDB(url string)*sql.DB{
	db,err := sql.Open("postgresql",url)

	if err != nil{
		log.Panic(err)
	}
	return db
}