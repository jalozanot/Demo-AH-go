package database_client

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fmcarrero/bookstore_utils-go/logger"
	"github.com/jalozanot/demoCeiba/infrastructure/adapters/repository/models"
	"github.com/jinzhu/gorm"
)

const (
	MysqlUsersUsername = "MYSQL_USERS_USERNAME"
	MysqlUsersPassword = "MYSQL_USERS_PASSWORD"
	MysqlUsersHost     = "MYSQL_USERS_HOST"
	MysqlUsersSchema   = "MYSQL_USERS_SCHEMA"
	MysqlUsersPort     = "MYSQL_USERS_PORT"
)

func GetDatabaseInstance() *gorm.DB {

	userName := os.Getenv(MysqlUsersUsername)
	password := os.Getenv(MysqlUsersPassword)
	host := os.Getenv(MysqlUsersHost)
	schema := os.Getenv(MysqlUsersSchema)
	port, _ := strconv.ParseInt(os.Getenv(MysqlUsersPort), 10, 64)

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=UTC", userName, password, host, port, schema)
	fmt.Println(dataSource)
	db, err := gorm.Open("mysql", "ceiba:ceiba@tcp(192.168.1.33:3306)/movie_ceiba?charset=utf8&parseTime=True&loc=UTC")
	if err != nil {
		logger.Error(err.Error(), err)
		_ = db.Close()
		panic("database not working")
	}
	db.SingularTable(true)
	migrateDatabase(db)

	return db
}

func migrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.MovieEntity{})
}
