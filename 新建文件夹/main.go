package main

/*
连接数据库，然后根据用户的请求去实现增、删、改、查

*/
import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //一定不要忘记gorm的数据库驱动，_只是为了用该包下的init函数。
	"net/http"
	//"main.go/tool"  如果想用tool包下的结构体、函数、方法、接口，使用模块名.包名的方式，引入该包
)

//这是一个todo模型
type ToDO struct {
	Id     int    `json:"id"` //与前端进行数据交互的时候，要用json格式。通过tag标签，实现将结构体数据转换为前端的Json数据。
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var (
	DB *gorm.DB //定义一个全局变量，该变量为指向DB结构体的指针变量，通过该指针变量可以实现对数据库的各种操作，相当于数据库引擎。
)

func InitMysql() (err error) {
	dsn := "root:123456@(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return nil
}
func main() {
	//创建数据库
	//create database bubble;	创建一个名为bubble的数据库
	//连接数据库
	err := InitMysql()
	if err != nil {
		panic(err)
	}
	//在执行完main函数里面所有的程序之后，执行defer语句
	defer DB.Close()
	//模型绑定
	DB.AutoMigrate(&ToDO{}) //传一个空结构体指针类型，在数据库中创建对应的表。
	//1、先创建一个gin框架引擎
	r := gin.Default()
	//告诉gin框架去哪里找index.html文件
	r.LoadHTMLGlob("./templates/*")
	//告诉gin框架引用的模板文件的静态文件去哪里找
	r.Static("/static", "static")
	//3、前段发出一个get请求，我们后端给出相应的响应
	r.GET("/slit", func(c *gin.Context) { //接口名为slit
		c.HTML(200, "index.html", nil) //StatusOK  = 200,可以看到statusok就是常量200,200表示状态正常
	}) //第3个参数可以任意填，不影响最后的结果
	v1Group := r.Group("v1") //首先是路由分组，路由分组完成后，在{}里面写一些增、删、改、查的操作
	{
		//待办事项
		//添加
		v1Group.POST("/todo", func(c *gin.Context) {
			//前端页面填写待办事项，点击提交，会发送请求到这里
			//1、从请求中把数据拿出来
			var todo ToDO //todo为结构体实例或者说是对象
			err := c.BindJSON(&todo)
			if err != nil {
				return
			}
			//2、存入到数据库中
			err = DB.Create(&todo).Error //这个地方是错误处理，创建完表之后，把错误信息返回。
			if err != nil {
				c.JSON(http.StatusOK, gin.H{ //gin.H，这个H相当于map[string]interface{}，调用JSON将map类型的数据转换成为json格式并返回给前端，这个gin.H就把map[string]interface{}给代替了。
					"err": err.Error(), //给前端返回一个错误,之所以返回一个200是说明请求是正常的，但是在将数据存放到数据库时，发生了错误
				})
			} else {
				c.JSON(http.StatusOK, &todo) //把todo结构体实例转换成Json格式来进行返回，返回到前端。
			}
			//这样实际上是把存入数据和返回响应放在一起了。
			//3、返回一个响应
		})
		//查看所有的待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			//查询to_dos这个表里所有的数据
			var todolist []ToDO            //该切片中存放的数据是todo结构体的实例或者说是对象
			err = DB.Find(&todolist).Error //将查询到的数据都存放到切片todolist中。
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(), //表示get请求是正常的，但是在从数据库中查询数据时发生了错误。
				})
			} else {
				c.JSON(http.StatusOK, &todolist) //在将所有的值以Json格式数据，在前端展示。
			}
		})
		//查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		//修改某一个待办事项
		v1Group.PUT("/todo/:id", func(c *gin.Context) { //:id相当于不知道具体的id值是什么，这是一个变量用来存放id值。
			id, _ := c.Params.Get("id")
			var todo ToDO
			err = DB.Where("id=?", id).First(&todo).Error //根据请求中的id值，将从数据库中查询出来的数据，放入todo结构体实例中。
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": "这个是无效的id"})
			}
			c.BindJSON(&todo)                           //从数据库中把数据拿出来之后，在对数据做一个更改。
			if err = DB.Save(&todo).Error; err != nil { //没有对数据进行任何的更新。
				c.JSON(http.StatusOK, gin.H{"error": "保存数据失败"})
			} else {
				c.JSON(http.StatusOK, todo) //没有对原来的数据做任何跟新，原来的数据是多少仍然是多少。
			}
		})
		//删除某一个待办事项
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, _ := c.Params.Get("id") //c.Params.Get()主要功能就是得到请求中的id值，并赋值给变量id。
			err = DB.Where("id=?", id).Delete(&ToDO{}).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": "这个是无效的id", //说明delete请求有效，但是在删除具体的值的时候，发生了错误。
				})
			} else {
				c.JSON(http.StatusOK, gin.H{"id": "id对应的数据库中的数据删除成功"})
			}
		})
	}
	//2、运行这个gin框架
	r.Run() //没有定义端口号，所以使用默认的端口号，8080

}
