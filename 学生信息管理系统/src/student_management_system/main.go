package main

import (
	"mysql_learn/service"
)

/*
实现一个学生信息管理系统
学生信息：学号
		姓名
		年龄
		班级
系统面板：
----------欢迎使用学生信息管理系统----------
1、增加学生信息
2、删除学生信息
3、修改学生信息
	1、修改学生姓名
	2、修改学生年龄
	3、修改学生班级
4、查看学生信息
	1、查看单个学生的信息
	2、查看所有学生的信息
5、退出系统
-----------------------------------------
*/
func main() {
	service.ConnectDatabase() //连接数据库
	for {
		service.ShowMenu()                             //功能菜单
		num := service.ChoiceNumberExecutiveFunction() //选择数字
		service.NumCorrespondOperate(num)              //执行数字对应的函数
	}
}
