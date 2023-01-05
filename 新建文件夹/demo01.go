package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Student struct {
	Id    int
	Name  string
	Age   int
	Class string
}

func main() {
	dsn := "root:123456@(127.0.0.1:3306)/dbexercise?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println("与数据库连接成功")
	defer DB.Close()
	//创建结构体与数据库的连接
	err = DB.AutoMigrate(&Student{}).Error
	if err != nil {
		panic(err)
	}
	fmt.Println("表创建成功")
	//s1 := Student{1, "洪树树", 19, "科技2班"}
	//s2 := Student{2, "洪树1", 19, "科技2班"}
	//s3 := Student{3, "洪树2", 19, "科技2班"}
	//s4 := Student{4, "洪树3", 19, "科技2班"}
	//DB.Debug().Create(&s1)
	//DB.Debug().Create(&s2)
	//DB.Debug().Create(&s3)
	//DB.Debug().Create(&s4)
	//var slice1 []Student
	//DB.Debug().Find(&slice1)
	//fmt.Println(slice1)
	//for _, value := range slice1 {
	//	fmt.Println(value.Id, value.Name, value.Age, value.Class)
	//}
	var user Student
	DB.Debug().Where("id=?", 4).First(&user)
	DB.Debug().Save(&user)
}
