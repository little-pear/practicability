package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/pelletier/go-toml"

	"github.com/gin-gonic/gin"
)

var gPort = flag.Int("port", 7890, "server port")

var gDB *SqlDB

// 设置跨域
//*
// server.OPTIONS("/", func(c *gin.Context) {
// 	c.Header("Access-Control-Allow-Origin", "*")
// 	c.String(http.StatusOK, "ok")
// })
//*/

func main() {
	server := gin.Default()

	var err error

	tree, err := toml.LoadFile("config.toml")
	if err != nil {
		log.Panic(err)
	}

	server.Static("ele/static/uploadfile/", "./static/uploadfile/")

	dbName := tree.GetDefault("mysql.username", "username").(string)
	dbPass := tree.GetDefault("mysql.passwd", "passwd").(string)
	dbHost := tree.GetDefault("mysql.Host", "Host").(string)
	dbPort := tree.GetDefault("mysql.Port", "Port").(string)
	dbSchema := tree.GetDefault("mysql.Schema", "Schema").(string)

	log.Printf("username:[%v],passwd:[%v],host:[%v]\n", dbName, dbPass, dbHost)

	gDB, err = OpenDB("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", dbName, dbPass, dbHost, dbPort, dbSchema))
	if err != nil {
		log.Panic(err)
	}
	defer gDB.Close()

	eleAPI := server.Group("/ele")
	{
		eleAPI.GET("Login/register_do", registerDo) //注册接口
		eleAPI.GET("Login/login_do", loginDo)
		eleAPI.GET("User/get_user", getUser)
		eleAPI.GET("User/mod_user", modUser)
		eleAPI.GET("User/feedback", feedback)
		eleAPI.GET("User/plus_money", plusMoney)

		eleAPI.GET("User/add_address", addAddress)
		eleAPI.GET("User/get_address", getAddress)
		eleAPI.GET("User/del_address", delAddress)
		eleAPI.GET("User/get_log", getLog)
		eleAPI.GET("User/get_message", getMessge)
		eleAPI.GET("User/get_cash", getCash)

		eleAPI.GET("Item/get_item", getItem)
		eleAPI.GET("Item/get_one", getOne)
		eleAPI.GET("Item/create_order", createOrder)
		eleAPI.GET("Item/money_pay", moneyPay)
		eleAPI.GET("Item/mod_state", modState)
		eleAPI.GET("Item/get_order")
		eleAPI.GET("Item/commont")

		eleAPI.POST("upload_file", uploadFile)
	}

	err = server.Run(fmt.Sprintf(":%v", *gPort))
	if err == nil {
		log.Fatal(err)
	}
}
