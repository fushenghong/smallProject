package main

import (
	"fmt"
	"os"
)

/*
功能：留言板，能实现增、删、改、查这4个小功能。
-------------欢迎使用祝福信息管理系统-------------
		1.查看所有祝福信息
		2.增加祝福信息
		3.修改祝福信息
		4.删除祝福信息
		5.退出程序
-------------------------------------------------
留言板内容：ID	幸运儿	祝福者	祝福语
*/
//菜单页面
func showMenu() {
	fmt.Println("-----------欢迎使用祝福留言板-------------")
	fmt.Println("1、查看所有祝福信息")
	fmt.Println("2、增加祝福信息")
	fmt.Println("3、修改祝福信息")
	fmt.Println("4、删除祝福信息")
	fmt.Println("5、退出程序")
}

//定义保存祝福信息的结构体
type wishMessage struct {
	wish map[int]*wish //存储祝福信息用map来存储，键是int类型，值是指向wish结构体的指针变量，也即wish结构体的指针对象。
}

//定义留言板内容结构体
type wish struct {
	wishID          int
	wishLucky       string
	wishWisher      string
	wishInformation string
}

//查看所有祝福信息
func (w wishMessage) showAll() {
	fmt.Println("-------------留言板信息-----------------")
	fmt.Println("ID 幸运儿 祝福者 祝福语")
	for _, value := range w.wish {
		fmt.Println(value.wishID, value.wishLucky, value.wishWisher, value.wishInformation)
	}
	fmt.Println("-----------------------------------------")
}

//增加祝福信息
func (w wishMessage) addMessage() {
	var (
		wishID          int
		wishLucky       string
		wishWisher      string
		wishInformation string
	)
	fmt.Println("请输入祝福ID：")
	fmt.Scanln(&wishID)
	fmt.Println("请输入接收祝福的幸运儿：")
	fmt.Scanln(&wishLucky)
	fmt.Println("请输入祝福者：")
	fmt.Scanln(&wishWisher)
	fmt.Println("请输入祝福信息：")
	fmt.Scanln(&wishInformation)
	w.wish[wishID] = &wish{wishID: wishID, wishLucky: wishLucky, wishWisher: wishWisher, wishInformation: wishInformation}
}

//修改祝福信息
func (w wishMessage) changeMessage() {
	var (
		input           int
		num             int
		wishLucky       string
		wishWisher      string
		wishInformation string
	)
	fmt.Println("请输入ID：")
	fmt.Scanln(&input)
	_, ok := w.wish[input]
	if ok != true {
		fmt.Println("数据库中没有该ID的留言")
		return
	}
	fmt.Print("请选择要修改的内容:\n1、修改幸运儿 2、修改祝福者 3、修改祝福语 4、全部\n")
	fmt.Println("请输入数字，以执行数字对应的操作")
	fmt.Scanln(&num)
	switch num {
	case 1:
		fmt.Println("请输入修改后的幸运儿：")
		fmt.Scanln(&wishLucky)
		w.wish[input].wishLucky = wishLucky
	case 2:
		fmt.Println("请输入修改后的祝福者：")
		fmt.Scanln(&wishWisher)
		w.wish[input].wishWisher = wishWisher
	case 3:
		fmt.Println("请输入修改后的祝福语：")
		fmt.Scanln(&wishInformation)
		w.wish[input].wishInformation = wishInformation
	case 4:
		fmt.Println("请输入修改后的幸运儿：")
		fmt.Scanln(&wishLucky)
		fmt.Println("请输入修改后的祝福者：")
		fmt.Scanln(&wishWisher)
		fmt.Println("请输入修改后的祝福语：")
		fmt.Scanln(&wishInformation)
		w.wish[input] = &wish{wishID: input, wishLucky: wishLucky, wishWisher: wishWisher, wishInformation: wishInformation}
	default:
		fmt.Println("输入错误，该数字没有对应的操作")
	}
}

//删除祝福信息
func (w wishMessage) deleteMessage() {
	var input int
	fmt.Println("请输入要删除的ID：")
	fmt.Scanln(&input)

	delete(w.wish, input)
	fmt.Println("删除成功")
}
func main() {
	var w wishMessage                               //声明一个wishMessage类型的变量
	w = wishMessage{wish: make(map[int]*wish, 100)} //对map进行初始化，容量为100
	for {
		//1、展示留言板菜单
		showMenu()
		//2、根据留言板的提示，输入相应的序号
		fmt.Println("请输入上面对应的序号，来执行对应功能：")
		var input int
		fmt.Scanln(&input) //从键盘读取int型数值，并放入变量input中。
		//3、根据输入序号执行相应的功能。
		switch input {
		//查看所有祝福信息
		case 1:
			w.showAll()
		//增加祝福信息
		case 2:
			w.addMessage()
		//修改祝福信息
		case 3:
			w.changeMessage()
		//删除祝福信息
		case 4:
			w.deleteMessage()
		//退出程序
		case 5:
			os.Exit(1)
		default:
			fmt.Println("输入错误，没有该数值对应的功能")
		}
	}
}
