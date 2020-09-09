package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const timeFormat = "2006-01-02 15:04:05"

//User ...
type User struct {
	ID       int    `json:"id,string"`
	NickName string `json:"nickname"`
	Password string `json:"password"`
	HeadImg  string `json:"headimg"`
	Phone    string `json:"phone"`
	Money    string `json:"money"`
	CTime    string `json:"cTime"`
}

/*
CREATE TABLE `cash_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `phone` varchar(45) DEFAULT NULL,
  `type` varchar(45) DEFAULT NULL,
  `money` float DEFAULT NULL,
  `old_money` float DEFAULT NULL,
  `cTime` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
*/

type CastLog struct {
	ID       int     `json:"id,string"`
	Phone    string  `json:"phone"`
	Type     string  `json:"type"`
	Money    float64 `json:"money"`
	OldMoney float64 `json:"old_money"`
	CTime    string  `json:"cTime"`
}

/*
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `use_phone` varchar(45) DEFAULT NULL,
  `name` varchar(45) DEFAULT NULL,
  `sex` varchar(45) DEFAULT NULL,
  `phone` varchar(45) DEFAULT NULL,
  `address` text,
  `time` varchar(45) DEFAULT NULL
*/

type Address struct {
	ID        int    `json:"id,string"`
	UserPhone string `json:"user_phone"`
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	CTime     string `json:"cTime"`
}

type Message struct {
	ID      int    `json:"id,string"`
	Message string `json:"message"`
	CTime   string `json:"cTime"`
}

/*
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `class` int(11) DEFAULT NULL,
  `title` varchar(45) DEFAULT NULL,
  `img` varchar(45) DEFAULT NULL,
  `price` float DEFAULT NULL,
  `is_distribution` int(11) DEFAULT NULL,
  `buy_number` varchar(45) DEFAULT NULL,
  `distance` int(11) DEFAULT NULL,*/
type Item struct {
	ID             int     `json:"id,string"`
	Class          int     `json:"class,string"`
	Title          string  `json:"title"`
	Img            string  `json:"img"`
	Price          float64 `json:"price"`
	IsDistribution int     `json:"isDistribution"`
	BuyNumber      string  `json:"buyNumber"`
	Distance       int     `json:"distance"`
}

//注册
func registerDo(c *gin.Context) {

	phone := c.Query("phone_data")
	password := c.Query("password_data")
	t := time.Now().Format("2006-01-02 15:04:05")
	log.Printf("phone:[%v] passwd:[%v] time:[%v]\n", phone, password, t)

	rows, err := gDB.Select(fmt.Sprintf("select * from ele.user where phone = '%v'", phone))
	if err != nil {
		log.Printf("[Error]: err:[%v]", err)
		c.JSON(http.StatusOK, 0)
		return
	}

	// 已存在
	if rows.NextT() == true {
		c.JSON(http.StatusOK, -2)
		return
	}

	ret, err := gDB.Exec(fmt.Sprintf("insert into ele.user(password,phone,cTime) value('%v','%v','%v')", password, phone, t))
	if err != nil {
		log.Printf("[Error]: err:[%v]", err)
		c.JSON(http.StatusOK, 0)
		return
	}

	id, err := ret.LastInsertId()
	if err != nil {
		log.Printf("[Error]: err:[%v]", err)
		c.JSON(http.StatusOK, 0)
		return
	}

	gDB.Exec(fmt.Sprintf("insert into ele.message(phone,content,cTime) value('%v','%v','%v')",
		phone, "欢迎加入饿了么，开启全新生活，欢迎加入", t))

	// c.Data(http.StatusOK, "", []byte("OK"))
	c.JSON(http.StatusOK, id)
}

//登陆
func loginDo(c *gin.Context) {
	phone := c.Query("phone_data")
	password := c.Query("password_data")
	rows, err := gDB.Select(fmt.Sprintf("select password from ele.user where phone = '%v'", phone))
	if err != nil {
		log.Printf("[Error]: err:[%v]", err)
		c.JSON(http.StatusOK, 0)
		return
	}

	// 已存在
	if rows.NextT() {
		dbPwd := rows.GetField("password")
		if err != nil {
			log.Printf("[ERROR]: err:[%v]\n", err)
			c.JSON(http.StatusOK, -2)
		}

		if password == dbPwd {
			c.JSON(http.StatusOK, 1)
		} else {
			c.JSON(http.StatusOK, 0)
		}
	} else {
		c.JSON(http.StatusOK, -2)
	}
}

//获取用户信息
func getUser(c *gin.Context) {
	phone := c.Query("phone_data")

	u := dbGetUser(phone)
	if u == nil {
		c.JSON(http.StatusOK, -2)
		return
	}
	c.JSON(http.StatusOK, u)
}

//从DB中获取当前用户信息
func dbGetUser(phone string) *User {
	rows, err := gDB.Select(fmt.Sprintf("select * from ele.user where phone = '%v'", phone))
	if err != nil {
		log.Printf("[Error]: err:[%v]", err)
		return nil
	}

	u := &User{}
	// 已存在
	if rows.NextT() {
		u.ID, _ = strconv.Atoi(rows.GetField("id"))
		u.NickName = rows.GetField("nickName")
		u.Phone = rows.GetField("phone")
		u.Password = rows.GetField("password")
		u.CTime = rows.GetField("cTime")
		u.HeadImg = rows.GetField("headimg")
		u.Money = rows.GetField("money")
		return u
	}
	return nil
}

//修改用户信息
func modUser(c *gin.Context) {
	phone := c.Query("phone_data")
	headImg := c.Query("img_data")
	name := c.Query("nickname_data")

	_, err := gDB.Exec(fmt.Sprintf("update ele.user set nickname='%v', headimg='%v' where phone = '%v'", name, headImg, phone))
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusOK, -2)
		return
	}

	c.JSON(http.StatusOK, 0)
}

//反馈意见
func feedback(c *gin.Context) {
	phone := c.Query("phone_data")
	content := c.Query("content_data")
	t := time.Now().Format("2006-01-02 15:04:05")

	_, err := gDB.Exec(fmt.Sprintf("insert into ele.feedback(phone,content,cTime) value('%v','%v','%v')", phone, content, t))
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusOK, -2)
		return
	}
	c.JSON(http.StatusOK, 0)
}

//充值
func plusMoney(c *gin.Context) {
	phone := c.Query("phone_data")
	money := c.Query("money_data")

	u := dbGetUser(phone)
	if u == nil {
		c.JSON(http.StatusOK, -1)
		return
	}

	dbMoney, _ := strconv.Atoi(u.Money)
	addMoney, _ := strconv.Atoi(money)

	_, err := gDB.Exec(fmt.Sprintf("update ele.user set money='%v' where phone = '%v'", dbMoney+addMoney, phone))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, -1)
		return
	}

	_, err = gDB.Exec(fmt.Sprintf("insert into ele.cash_log (phone, type,  money, old_money,cTime) value('%v','%v','%v','%v','%v')", phone, "plus", addMoney, dbMoney, time.Now().Format("2006-01-02 15:04:05")))
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, dbMoney+addMoney)
}

//添加地址
func addAddress(c *gin.Context) {
	loginPhone := c.Query("phone_data1")
	name := c.Query("name_data")
	sex := c.Query("sex_data")
	phone := c.Query("phone_data2")
	address := c.Query("address_data")
	cTime := time.Now().Format("2006-01-02 15:04:05")

	//插入

	sql := fmt.Sprintf(`insert into ele.address(user_phone,name,sex,phone,address,cTime) 
	value('%v','%v','%v','%v','%v','%v')`, loginPhone, name, sex, phone, address, cTime)

	_, err := gDB.Exec(sql)
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, 0)

}

//获取地址
func getAddress(c *gin.Context) {
	phone := c.Query("phone_data")
	rows, err := gDB.Select(fmt.Sprintf("select * from ele.address where user_phone = '%v' order by id desc", phone))
	if err != nil {
		log.Println(err)
	}

	ads := make([]*Address, 0, rows.ResNum())

	for rows.NextT() {
		l := &Address{}
		l.ID, _ = strconv.Atoi(rows.GetField("id"))
		l.Phone = rows.GetField("phone")
		l.Name = rows.GetField("name")
		l.Sex = rows.GetField("sex")
		l.Address = rows.GetField("address")
		l.CTime = rows.GetField("cTime")
		ads = append(ads, l)
	}

	c.JSON(http.StatusOK, ads)
}

//删除地址
func delAddress(c *gin.Context) {
	id := c.Query("id_data")

	sql := fmt.Sprintf("delete from ele.address where id = %v", id)

	_, err := gDB.Select(sql)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, -1)
	}

	c.JSON(http.StatusOK, 0)
}

//获取充值记录
func getLog(c *gin.Context) {
	phone := c.Query("phone_data")
	rows, err := gDB.Select(fmt.Sprintf("select * from ele.cash_log where phone = '%v' order by id desc", phone))
	if err != nil {
		log.Println(err)
	}

	logs := make([]*CastLog, 0, rows.ResNum())

	for rows.NextT() {
		l := &CastLog{}
		l.ID, _ = strconv.Atoi(rows.GetField("id"))
		l.Phone = rows.GetField("phone")
		l.Money, _ = strconv.ParseFloat(rows.GetField("money"), 64)
		l.Type = rows.GetField("type")
		l.OldMoney, _ = strconv.ParseFloat(rows.GetField("old_money"), 64)
		l.CTime = rows.GetField("cTime")
		logs = append(logs, l)
	}

	c.JSON(http.StatusOK, logs)
}

//获取消息信息
func getMessge(c *gin.Context) {
	phone := c.Query("phone_data")
	rows, err := gDB.Select(fmt.Sprintf("select * from ele.message where phone = '%v' order by id desc", phone))
	if err != nil {
		log.Println(err)
	}

	msgs := make([]*Message, 0, rows.ResNum())

	for rows.NextT() {
		s := &Message{}
		s.ID, _ = strconv.Atoi(rows.GetField("id"))
		s.Message = rows.GetField("content")
		s.CTime = rows.GetField("cTime")
		msgs = append(msgs, s)
	}

	c.JSON(http.StatusOK, msgs)
}

//提现
func getCash(c *gin.Context) {
	phone := c.Query("phone_data")
	money := c.Query("money_data")

	u := dbGetUser(phone)
	if u == nil {
		c.JSON(http.StatusOK, -1)
		return
	}

	dbMoney, _ := strconv.Atoi(u.Money)
	cashMoney, _ := strconv.Atoi(money)

	sql := fmt.Sprintf("update ele.user set money='%v' where phone = '%v'", dbMoney-cashMoney, phone)
	_, err := gDB.Exec(sql)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, -1)
		return
	}

	_, err = gDB.Exec(fmt.Sprintf("insert into ele.cash_log (phone, type,  money, old_money,cTime) value('%v','%v','%v','%v','%v')", phone, "get", cashMoney, dbMoney, time.Now().Format("2006-01-02 15:04:05")))
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, dbMoney-cashMoney)
	// c.JSON(http.StatusOK, 0)
}

//首页获取商品
func getItem(c *gin.Context) {

	class := c.Query("class_data")

	filter := ""
	if class != "" {
		filter += fmt.Sprintf(" and c.name = '%v'", class)
	}

	sql := fmt.Sprintf(`
SELECT 
    it.*
FROM
    ele.item it
        LEFT JOIN
    ele.class c ON c.id = it.class
WHERE
    1 = 1 %v
ORDER BY it.id DESC`, filter)

	rows, err := gDB.Select(sql)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, -1)
		return
	}

	its := make([]*Item, 0, rows.ResNum())

	for rows.NextT() {
		it := &Item{}

		it.ID, _ = strconv.Atoi(rows.GetField("id"))
		it.Class, _ = strconv.Atoi(rows.GetField("class"))
		it.Title = rows.GetField("title")
		it.Img = rows.GetField("img")
		it.Price, _ = strconv.ParseFloat(rows.GetField("price"), 64)
		it.IsDistribution, _ = strconv.Atoi(rows.GetField("is_distribution"))
		it.BuyNumber = rows.GetField("buy_number")
		it.Distance, _ = strconv.Atoi(rows.GetField("distance"))

		its = append(its, it)
	}

	c.JSON(http.StatusOK, its)
}

//首页获取商品
func getOne(c *gin.Context) {

	id := c.Query("id_data")

	sql := fmt.Sprintf(`
SELECT 
    it.*
FROM
    ele.item it
        LEFT JOIN
    ele.class c ON c.id = it.class
WHERE
    1 = 1 and it.id = %v
ORDER BY it.id DESC`, id)

	rows, err := gDB.Select(sql)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, -1)
		return
	}

	for rows.NextT() {
		it := &Item{}

		it.ID, _ = strconv.Atoi(rows.GetField("id"))
		it.Class, _ = strconv.Atoi(rows.GetField("class"))
		it.Title = rows.GetField("title")
		it.Img = rows.GetField("img")
		it.Price, _ = strconv.ParseFloat(rows.GetField("price"), 64)
		it.IsDistribution, _ = strconv.Atoi(rows.GetField("is_distribution"))
		it.BuyNumber = rows.GetField("buy_number")
		it.Distance, _ = strconv.Atoi(rows.GetField("distance"))

		c.JSON(http.StatusOK, it)
		return
	}

	c.JSON(http.StatusOK, 0)
}

func createOrder(c *gin.Context) {
	phone := c.Query("phone_data")
	itemID := c.Query("item_id")
	money := c.Query("money_data")
	t := time.Now().Format(timeFormat)
	/*
			CREATE TABLE `order` (
		  `id` int(11) NOT NULL AUTO_INCREMENT,
		  `phone` varchar(45) DEFAULT NULL,
		  `item_id` int(11) DEFAULT NULL,
		  `money` float DEFAULT NULL,
		  `state` tinyint(2) DEFAULT '0',
		  `cTime` varchar(45) DEFAULT NULL,
		  PRIMARY KEY (`id`)*/

	sql := fmt.Sprintf(`insert into ele.order(phone,item_id,money,cTime)value('%v','%v','%v','%v')`,
		phone, itemID, money, t)

	ret, err := gDB.Exec(sql)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, -1)
		return
	}

	id, _ := ret.LastInsertId()
	c.JSON(http.StatusOK, id)

}

func moneyPay(c *gin.Context) {
	phone := c.Query("phone_data")
	money := c.Query("money_data")

	u := dbGetUser(phone)
	if u == nil {
		c.JSON(http.StatusOK, -1)
		return
	}

	dbMoney, _ := strconv.Atoi(u.Money)
	cashMoney, _ := strconv.Atoi(money)

	sql := fmt.Sprintf("update ele.user set money='%v' where phone = '%v'", dbMoney-cashMoney, phone)
	_, err := gDB.Exec(sql)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, -1)
		return
	}

	_, err = gDB.Exec(fmt.Sprintf("insert into ele.cash_log (phone, type,  money, old_money,cTime) value('%v','%v','%v','%v','%v')", phone, "use", cashMoney, dbMoney, time.Now().Format("2006-01-02 15:04:05")))
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, dbMoney-cashMoney)
}

func modState(c *gin.Context) {
	id := c.Query("id_data")
	_, err := gDB.Exec(fmt.Sprintf("update ele.order set state=1 where id = '%v'", id))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, -1)
		return
	}
	c.JSON(http.StatusOK, 0)
}

//上传文件
func uploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("uploadkey")

	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	//文件的名称
	filename := header.Filename

	fmt.Println(file, err, filename)
	//创建文件
	out, err := os.Create("static/uploadfile/" + filename)
	//注意此处的 static/uploadfile/ 不是/static/uploadfile/
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, "static/uploadfile/"+filename)
}
