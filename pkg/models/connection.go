package models

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/yashre-bh/kla-crm-btp/pkg/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

func dsn() string {
	var config types.Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name)

}

func Connect() (*gorm.DB, error) {
	connection, err := gorm.Open(mysql.Open(dsn()), &gorm.Config{})
	if connection == nil || err != nil {
		fmt.Println("Error connecting to database")
		return connection, err
	}

	connection.AutoMigrate(&types.Checkpoint{})
	connection.AutoMigrate(&types.Employee{})

	return connection, err
}
