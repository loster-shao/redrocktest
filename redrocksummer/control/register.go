package control

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"redrocksummer/SQL"
)

type user struct{
	gorm.Model
	Username string
	Password string
}

func Register(username, password string/*c *gin.Context*/)  {
	////post接受数据
	//username := c.PostForm("username")
	//password := c.PostForm("password")

	//数据库连接


	//用户名不可重复
	var u []user
	SQL.DB.Where("username=?", username).Find(&u)
	//if len(u) != 0{
	//	c.JSON(502,
	//		gin.H{"status": http.StatusBadGateway,
	//			"message": "用户名重复，请重新修改"})
	//	return
	//}

	//注册
	SQL.DB.AutoMigrate(&user{})
	SQL.DB.Create(&user{
		Username:username,
		Password:password})
	//c.JSON(200,
	//	gin.H{
	//	"status":  http.StatusOK,
	//	"message": "注册成功"})
}
