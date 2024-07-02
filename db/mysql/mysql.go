package mysql

import (
	"example/config"
	"example/model"
	"fmt"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

type Client struct {
	*gorm.DB
}

var client *Client

func Initialize() {
	mysqlInfo := config.GetMysqlInfo()

	parse := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlInfo.Username, mysqlInfo.Password, mysqlInfo.Host, mysqlInfo.Port, mysqlInfo.DBName)

	db, err := gorm.Open(mysql.Open(parse), &gorm.Config{})
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	isMigrate, _ := strconv.ParseBool(mysqlInfo.IsMigrate)
	if isMigrate {
		err = model.CreateEconnectTable(db)
		if err != nil {
			panic(err)
		}
	}

	client = &Client{db}
}

func GetDB() *Client {
	if client == nil {
		Initialize()
	}

	return client
}
