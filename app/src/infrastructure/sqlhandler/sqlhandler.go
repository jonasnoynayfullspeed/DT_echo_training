package sqlhandler

import (
	"app/src/entities"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
	DB *gorm.DB
}

func MigrateEntities(DB *gorm.DB) {
	migrator := DB.Migrator()
	if !migrator.HasTable(entities.Blog{}) {
		err := DB.AutoMigrate(&entities.Blog{})
		if err != nil {
			log.Fatalln("Problem with migrating entities to mysql database")
		}
	}
}

func NewSqlHandler() *SqlHandler {
	sqlHandler := new(SqlHandler)
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DATABASE")
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname)
	DB, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatalln(connection + "database can't connect")
	}
	MigrateEntities(DB)

	sqlHandler.DB = DB
	return sqlHandler
}
