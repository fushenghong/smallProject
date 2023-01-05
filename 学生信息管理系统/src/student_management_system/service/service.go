package service

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"mysql_learn/template"
	"os"
)

var db *sql.DB //定义一个数据库对象的全局变量

// ShowMenu 功能面板
func ShowMenu() {
	println("----------欢迎使用学生信息管理系统----------")
	println("1、增加学生信息")
	println("2、删除学生信息")
	println("3、修改学生信息")
	println("4、查看学生信息")
	println("5、退出系统")
	println("-----------------------------------------")
}

// ChoiceNumberExecutiveFunction 选择数字对应的功能
func ChoiceNumberExecutiveFunction() int {
	fmt.Println("请输入数字，系统来执行该数字对应的功能")
	var inputNumber int
	_, err := fmt.Scanln(&inputNumber)
	if err != nil {
		log.Fatal(err)
	}
	return inputNumber
}

// NumCorrespondOperate 数字对应操作
func NumCorrespondOperate(num int) {
	switch num {
	case 1:
		//增
		AddData()
	case 2:
		//删
		DeleteData()
	case 3:
		//改
		UpdateData()
	case 4:
		//查
		LookStudentMessage()
	case 5:
		os.Exit(1) //退出系统
	}
}

// ConnectDatabase 数据库的连接
func ConnectDatabase() {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接发生错误，err：", err)
		return
	}
	//尝试与数据库建立连接，校验dsn是否正确。
	err = db.Ping()
	if err != nil {
		fmt.Println("使用Ping()发现Ping不同，说明数据库连接失败。")
		return
	}
	fmt.Println("数据库连接成功")
}

// CreateTable 创建数据库中的表
//func CreateTable() {
//	sqlStr := "create table studentMessage (Id integer primary key auto_increment,Name varchar(50),Age integer,Class varchar(50))"
//	_, err := db.Exec(sqlStr)
//	if err != nil {
//		fmt.Println("表创建失败，err：", err)
//		return
//	}
//}

//CloseDatabase  关闭数据库
//func CloseDatabase() {
//	err := db.Close()
//	if err != nil {
//		fmt.Println("关闭数据库出现错误，err：", err)
//		return
//	}
//}

// AddData 往数据库中插入数据
func AddData() {
	var (
		Id    int
		Name  string
		Age   int
		Class string
	)
	println("----------请输入学生的各项基本信息----------")
	println("1、请输入学生学号：")
	fmt.Scanln(&Id)
	println("2、请输入学生姓名")
	fmt.Scanln(&Name)
	println("3、请输入学生年龄")
	fmt.Scanln(&Age)
	println("4、请输入学生班级")
	fmt.Scanln(&Class)
	sqlStr := "insert into student_message(Id,Name,Age,Class) values(?,?,?,?)"
	_, err := db.Exec(sqlStr, Id, Name, Age, Class)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("插入数据成功。")
}

// DeleteData 删除学生信息
func DeleteData() {
	var Id int
	fmt.Println("请输入要删除的学生Id:")
	fmt.Scanln(&Id)
	sqlStr := "delete from student_message where Id=?"
	_, err := db.Exec(sqlStr, Id)
	if err != nil {
		fmt.Println("从数据中删除数据发生错误，err：", err)
		return
	}
	fmt.Printf("Id为：%v的学生个人信息删除成功。\n", Id)
}

//studentMessageMenu 学生个人信息菜单
func studentMessageMenu() {
	println("1、修改学生姓名")
	println("2、修改学生年龄")
	println("3、修改学生班级")
	println("4、修改学生全部信息")
	fmt.Println("请输入相应数字：")
}

// UpdateData 修改学生信息
func UpdateData() {
	var (
		inputNumber int
		Id          int
		Name        string
		Age         int
		Class       string
	)

	fmt.Println("----------请选择要修改哪部分的信息----------")
	studentMessageMenu()
	fmt.Scanln(&inputNumber)
	switch inputNumber {
	case 1:
		fmt.Println("请输入学生Id:")
		fmt.Scanln(&Id)
		fmt.Println("请输入修改后的学生姓名：")
		fmt.Scanln(&Name)
		sqlStr := "update student_message set Name=? where Id=?"
		_, err := db.Exec(sqlStr, Name, Id)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		fmt.Println("学生姓名修改成功")
	case 2:
		fmt.Println("请输入学生Id:")
		fmt.Scanln(&Id)
		fmt.Println("请输入修改后的学生年龄：")
		fmt.Scanln(&Age)
		sqlStr := "update student_message set Age=? where Id=?"
		_, err := db.Exec(sqlStr, Age, Id)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		fmt.Println("学生年龄修改成功")
	case 3:
		fmt.Println("请输入学生Id:")
		fmt.Scanln(&Id)
		fmt.Println("请输入修改后的学生班级：")
		fmt.Scanln(&Class)
		sqlStr := "update student_message set Class=? where Id=?"
		_, err := db.Exec(sqlStr, Class, Id)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		fmt.Println("学生班级修改成功")
	case 4:
		fmt.Println("请输入学生Id:")
		fmt.Scanln(&Id)
		sql := "delete from student_message where Id=?"
		_, err := db.Exec(sql, Id)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		AddData()
	}
}

// LookStudentMessage 查看学生个人信息
func LookStudentMessage() {
	var (
		inputNumber int
		Id          int
	)
	fmt.Println("----------查看学生个人信息----------")
	println("1、查看单个学生的信息")
	println("2、查看所有学生的信息")
	fmt.Println("请输入数字：")
	fmt.Scanln(&inputNumber)
	switch inputNumber {
	case 1:
		fmt.Println("请输入要查询学生的Id:")
		fmt.Scanln(&Id)
		sqlStr := "select Id,Name,Age,Class from student_message where Id=?"
		var u template.StudentMessage
		err := db.QueryRow(sqlStr, Id).Scan(&u.Id, &u.Name, &u.Age, &u.Class)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("----------学生信息如下----------")
		fmt.Printf("Id:%v\tName:%v\tAge:%v\tClass:%v\n", u.Id, u.Name, u.Age, u.Class)
		fmt.Println("-------------------------------")
	case 2:
		sqlStr := "select Id,Name,Age,Class from student_message "
		var u template.StudentMessage //定义一个结构体对象
		rows, err := db.Query(sqlStr) //返回一个指向Rows结构体的一个指针变量。
		if err != nil {
			log.Fatal(err)
		}
		//非常重要，关闭rows释放持有的数据库连接。
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(rows)
		//循环读取结果集中的数据。(rows便是结果集)
		for rows.Next() {
			err := rows.Scan(&u.Id, &u.Name, &u.Age, &u.Class)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Id:%v\tName:%v\tAge:%v\tClass:%v\n", u.Id, u.Name, u.Age, u.Class)
		}
	}
}
