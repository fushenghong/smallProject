package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type List struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   ` json:"status"`
}

//通过gin框架实现，小清单。这种小功能
func main() {
	//创建数据库，用来存放数据
	//create database list
	//与数据库连接
	dsn := "root:123456@(127.0.0.1:3306)/list?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	//创建数据表，需要有一个结构体来映射到该数据表
	db.AutoMigrate(&List{})
	//创建路由引擎
	engine := gin.Default()
	//与html文件进行通信，将前端页面在浏览器端展现出来
	//在将前端页面展示的时候先要进行初始化
	//1、告诉context，html文件在哪里找
	engine.LoadHTMLGlob("./templates/*") //在当前目录下的tmeplates目录下。
	//2、加载静态文件
	engine.Static("./static", "static") //在static目录下来查找静态文件
	engine.GET("/list", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
	//通过路由分组来完成增、删、改、查的功能
	v1Group := engine.Group("/v1")
	{
		//增	实际上就是前端发送一个POST请求，把请求所带的内容存放到数据库List中。
		v1Group.POST("/todo", func(context *gin.Context) {
			var l1 List                  //创建一个list结构体对象
			err := context.BindJSON(&l1) //将结构体指针绑定
			if err != nil {
				fmt.Println(err)
			} else {
				if err = db.Create(&l1).Error; err != nil {
					context.JSON(http.StatusOK, gin.H{"err": "请求没问题，但是在往数据库中写数据时发生了问题。"})
				} else {
					context.JSON(http.StatusOK, nil)
				}
			}
		})
		//删	实际上就是根据请求中的id来删除数据库中的对应数据，id是未知的。
		v1Group.DELETE("/todo/:id", func(context *gin.Context) {
			id, _ := context.Params.Get("id")
			err = db.Where("id=?", id).Delete(&List{}).Error
			if err != nil {
				context.JSON(http.StatusOK, gin.H{"err": err.Error()})
			}
			context.JSON(http.StatusOK, gin.H{"message": "删除成功"})
		})

		//改	实际上本项目没有修改的功能，只是将原来的数据，还是原来的值存放到数据库中，同样的，修改也是根据id进行修改。
		v1Group.PUT("/todo/:id", func(context *gin.Context) {
			id, _ := context.Params.Get("id")
			var l2 List
			err = db.Where("id=?", id).First(&l2).Error
			if err != nil {
				context.JSON(http.StatusOK, gin.H{"err": err.Error()})
			}
			context.BindJSON(&l2)
			err = db.Where("id=?", id).Save(&l2).Error
			if err != nil {
				context.JSON(http.StatusOK, gin.H{"err": err.Error()})
			}
			context.JSON(http.StatusOK, &l2) //将更改后的数据返回到前端，实际上没有更改功能。
		})
		//查	实际上就是把数据库中的所有内容查询出来，在返回到前端页面。
		v1Group.GET("/todo", func(context *gin.Context) {
			var slice1 []List //用一个切片来存放从数据库中读取的数据
			err = db.Find(&slice1).Error
			if err != nil {
				context.JSON(http.StatusOK, gin.H{"err": "请求正常，但是查询数据发生错误喽"})
			}
			context.JSON(http.StatusOK, &slice1)
		})
	}
	//启动web服务
	engine.Run(":9999") //在9999端口来进行一个监听
}
