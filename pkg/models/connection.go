package models

import (
	"database/sql"
	"log"

	"context"
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"github.com/yashre-bh/kla-crm-btp/pkg/types"

	_ "github.com/go-sql-driver/mysql"
)

func dsn() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	var config types.Config
	_, err = toml.DecodeFile("config.toml", &config)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Name)
}

func Connection() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn())
	if err != nil {
		log.Printf("Error: %s when opening DB", err)
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		fmt.Printf("Errors %s pinging DB", err)
		return nil, err
	}
	fmt.Printf("Connected to DB successfully\n")
	return db, err
}
