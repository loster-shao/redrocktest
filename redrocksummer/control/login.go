package control

import (
	"fmt"
	"redrocksummer/SQL"
)


//登录
func Login(username, password string/*c *gin.Context*/){
	//username := c.PostForm("username")
	//password := c.PostForm("password")
	fmt.Println("login")
	var u user
	SQL.DB.Where("username = ?", username).First(&u)
	fmt.Println("?")
	//tokens := token.Create(u.Username, int(u.Model.ID))//创建token
	if u.Username == ""{
		fmt.Println("no")
		Register(username, password)
	}
	if password != u.Password {
		//v.Write
		//	c.JSON(200,gin.H{"status:": http.StatusOK, "message":"登录成功", "token": &tokens})
		//}else {
		//	c.JSON(400,gin.H{"status:": http.StatusBadRequest, "message": "用户名或密码错误"})
	}
}