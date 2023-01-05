package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// gin框架提供给开发者表单实体绑定的功能，可以将表单数据与结构体绑定
type userRegister struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Phone    int    `json:"phone"`
}

func Register(ctx *gin.Context) {
	var userForm = userRegister{}

	// 127.0.0.1:8000/register?username=马亚南&password=123465&phone=15188945949
	/* ShouldBindQuery无论啥请求，都只会绑定查询字符串中的参数
	   if err := ctx.ShouldBindQuery(&userForm); err != nil {
	       log.Fatalln("ShouldBindQuery failed")
	   }
	*/
	/*
	   ShouldBind如果是GET请求会绑定查询字符串中的参数，
	   如果是POST请求，优先绑定form表单或Json字符串中的数据，如果没有也可以绑定查询字符串中的数据
	      err := ctx.ShouldBind(&userForm)
	*/
	// ShouldBindJson只能接收json格式的数据，查询字符串参数或form表单数据都不能接收
	// Bind开头和ShouldBind开头的区别？
	// Bind会在header头中添加400的返回信息，而ShouldBind不会
	err := ctx.BindJSON(&userForm)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(userForm)

	ctx.Writer.Write([]byte(fmt.Sprintf(
		"%s:%s:%d", userForm.UserName, userForm.Password, userForm.Phone,
	)))
}

func main() {
	router := gin.Default()

	router.POST("/register", Register)

	router.Run(":8000")
}
