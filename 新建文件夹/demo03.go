package main

import (
	"fmt"
	"strconv"
)

type myfun func(int, int) string //type关键字定义了一个有2个形参，一个返回值的函数类型

func fun1() myfun { //没有参数传入，但是返回一个函数类型，即将fun1()函数体里面的匿名函数返回，因匿名函数赋值给变量fun了。所以通过fun即可调用匿名函数。
	fun := func(a, b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}
	return fun //将匿名函数返回到函数调用处
}

func main() {
	testfun := fun1()              //调用fun1()函数，返回一个2个参数，1个返回值的函数类型。并将值赋给变量testfun。
	fmt.Println(testfun(100, 200)) //通过调用testfun()函数，即调用了fun()函数，即匿名函数功能。匿名函数实现了将
	//2个int类型的值转换为string类型的值。最后的结果是100200
}
