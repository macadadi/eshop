package db

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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
  if err := godotenv.Load(); err != nil{
	log.Fatal("could not load environment file")
  }
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
	db,err := sql.Open("postgres",url)

	if err != nil{
		log.Panic(err)
	}
	err = db.Ping()
	if err != nil{
		log.Panic(err)
	}
	return db
}