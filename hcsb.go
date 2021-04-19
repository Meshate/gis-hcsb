package gis_hcsb

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
)

var (
	DB *gorm.DB
)

func Setup() {
	var err error
	mysqlName := viper.GetString("mysql.user")
	mysqlKey := viper.GetString("mysql.pass")
	mysqlAddress := viper.GetString("mysql.address")
	mysqlDatabase := viper.GetString("mysql.database")
	str := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", mysqlName, mysqlKey, mysqlAddress, mysqlDatabase)
	log.Println(str)
	DB, err = gorm.Open("mysql", str)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("<----------------------database connected---------------------->")
	}
}
