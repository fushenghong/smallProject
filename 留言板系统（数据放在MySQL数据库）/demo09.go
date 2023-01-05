package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

/*
功能：留言板，能实现增、删、改、查这4个小功能（注：把信息存到数据库当中）。
-------------欢迎使用留言板系统-------------
		1.查看所有祝福信息
		2.增加祝福信息
		3.修改祝福信息
		4.删除祝福信息
		5.退出程序
-------------------------------------------------
留言板内容：ID	幸运儿	祝福者	祝福语
*/
//定义一个结构体，来映射成表，把输入的信息存到表中。表中的一条数据相当于一个结构体对象或者说是实例。
type WishMessage struct {
	ID          int
	Lucky       string
	WishWisher  string
	Information string
}
//注意：无论是结构体名还是字段名，首字母都要大写，小写gorm就不能正常翻译成SQL语句了。这是个惨痛的教训。
func ShowMunu() {
	fmt.Println("---------欢迎使用留言板系统--------")
	fmt.Println("1、查看所有祝福信息")
	fmt.Println("2、增加祝福信息")
	fmt.Println("3、修改祝福信息")
	fmt.Println("4、删除祝福信息")
	fmt.Println("5、退出程序")
	fmt.Println("-----------------------------------")
	fmt.Println("请输入相应数字，来执行对应功能：")
}

//增加祝福信息
func AddMessage() *WishMessage {
	var (
		ID          int
		Lucky       string
		WishWisher  string
		Information string
	)
	fmt.Println("请输入ID:")
	fmt.Scanln(&ID)
	fmt.Println("请输入幸运儿:")
	fmt.Scanln(&Lucky)
	fmt.Println("请输入祝福者:")
	fmt.Scanln(&WishWisher)
	fmt.Println("请输入祝福语:")
	fmt.Scanln(&Information)
	return &WishMessage{ID, Lucky, WishWisher, Information}
}

func main() {
	for {
		//连接数据库
		db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/message_board?charset=utf8mb4&parseTime=True&loc=Local")
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		//关闭数据库
		defer db.Close()
		//在数据库中创建一张二维表，用来存储数据。
		db.AutoMigrate(&WishMessage{})
		//1、展示菜单
		ShowMunu()
		var input int
		fmt.Scanln(&input)
		//2、根据数字执行对应的功能
		switch input {
		case 1:
			//从数据库表中读取数据，将表中的数据全都读取出来，展示到控制台
			var slice1 []WishMessage
			db.Find(&slice1)
			fmt.Println("--------留言板信息如下--------")
			fmt.Println("ID 幸运儿 祝福者 祝福语")
			for _, value := range slice1 {
				fmt.Println(value.ID, value.Lucky, value.WishWisher, value.Information)
			}
		case 2:
			//往数据库表中添加数据
			db.Create(AddMessage())
		case 3:
			//根据相应的数字修改数据库中的信息
			var (
				input2      int
				ID1         int
				lucky1      string
				wishwisher1 string
				information1 string
			)
			fmt.Println("请输入ID：")
			fmt.Scanln(&ID1)
			fmt.Print("请选择要修改的内容:\n1、修改幸运儿 2、修改祝福者 3、修改祝福语 4、全部\n")
			fmt.Scanln(&input2)
			//根据数字来执行相应的功能
			switch input2 {
			case 1:
				fmt.Println("修改之后的幸运儿为：")
				fmt.Scanln(&lucky1)
				db.Model(&WishMessage{}).Where("id = ?", ID1).Update("lucky", lucky1)
			case 2:
				fmt.Println("修改之后的祝福者为：")
				fmt.Scanln(&wishwisher1)
				db.Model(&WishMessage{}).Where("id = ?", ID1).Update("wish_wisher", wishwisher1)
			case 3:
				fmt.Println("修改之后的祝福语为：")
				fmt.Scanln(&information1)
				db.Model(&WishMessage{}).Where("id = ?", ID1).Update("information", information1)
			case 4:
				db.Where("id=?", ID1).Delete(WishMessage{})
				db.Create(AddMessage())
			default:
				fmt.Println("数字输入错误，无该数字对应的功能。")
			}
		case 4:
			//将数据从数据库表中删除
			fmt.Println("请输入要删除的ID：")
			var input1 int
			fmt.Scanln(&input1)
			db.Where("id=?", input1).Delete(WishMessage{})
			fmt.Println("删除成功")
		case 5:
			os.Exit(1)
		default:
			fmt.Println("输入错误，该数字没有对应的功能，是无效的。")
		}
	}
}
