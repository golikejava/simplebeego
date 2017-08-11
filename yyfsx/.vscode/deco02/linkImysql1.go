package main

import (
	//bytes包实现了操作[]byte的常用函数
	"bytes"
	//打印sql数据是否传到这个变量里面
	"fmt"
	//sql包提供了保证SQL或类SQL数据库的泛用接口。
	"database/sql"
	//mysql驱动
	_ "github.com/Go-SQL-Driver/MySQL"
	//提供post接口
	"net/http"
	//输入输出接口
	"io/ioutil"
	//json包
	"encoding/json"
)

type User struct { //定义结构体
	Name    string
	Address string
}

func main() {
	//第一步：数据库取数据
	//定义变量
	var xingzheng002 []User
	db, err := sql.Open("mysql", "root:123456@/mysql_pythoncz?charset=utf8")
	checkErr(err)
	rows, err := db.Query("select name,address from xingzhen001")
	checkErr(err)
	//循环遍历查询结果,并传递给变量xingzheng002
	for rows.Next() {
		a := User{}
		rows.Scan(&a.Name, &a.Address)
		xingzheng002 = append(xingzheng002, a)
		fmt.Println(xingzheng002)
	}
	fmt.Println("================================================================")
	
	

	// //协程数量为10，并发，数据一个个发出去
	// var chan1 chan User
	// for c := 0; c < 10; c++ {
	// 	go func() {
	// 		ch := make(chan string, 1)
	// 		ch <- gangjiasuo()
	// 	}()

	// }

	//第二步：把数据通过post请求发送给服务器
	b, err := json.Marshal(xingzheng002)
	a := bytes.NewBuffer(b)
	checkErr(err)
	fmt.Println(xingzheng002)
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++")
	client := &http.Client{}                    //客户端生成，参数默认
	url := "http://localhost:8080/v1/user"      //beego项目中的地址指向
	req, err := http.NewRequest("POST", url, a) //request请求
	checkErr(err)
	req.Header.Set("Content-Type", "application/json;charset=utf-8") //请求头的设置
	resp, err := client.Do(req)                                      //处理返回结果
	result, err := ioutil.ReadAll(resp.Body)                         //读取resp对象中的数据
	resp.Body.Close()                                                //关闭数据流
	checkErr(err)
	fmt.Println(string(result)) //打印result的结果
	fmt.Println("---------------------------------------------------")
}

//err报错打印
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
