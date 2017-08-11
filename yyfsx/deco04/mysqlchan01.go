package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/Go-SQL-Driver/MySQL"
)

type User struct {
	Name    string
	Address string
}

var UserArray []User

func main() {
	//创建数组

	//数据库查询数据
	db, err := sql.Open("mysql", "root:123456@/mysql_pythoncz?charset=utf8")
	checkErr(err)
	rows, err := db.Query("select name,address from xingzhen001")
	checkErr(err)
	for rows.Next() {
		b := User{}
		rows.Scan(&b.Name, &b.Address)
		UserArray = append(UserArray, b)
	}

	CSYYF001()
	time.Sleep(1e9)

}

func CSYYF001() {

	channo := make(chan User)
	var temp User
	for d := 0; d < 10; d++ {

		go func(c chan User) {
			for f := 0; f < 20; f++ {
				temp = <-c
				x, err := json.Marshal(temp)
				checkErr(err)
				y := bytes.NewBuffer(x)
				client := &http.Client{}
				url := "http://localhost:8080/v1/user"
				req, err := http.NewRequest("POST", url, y)
				checkErr(err)
				req.Header.Set("Content-Type", "application/json;charset=utf-8") //请求头的设置
				resp, err := client.Do(req)                                      //处理返回结果
				//result, err := ioutil.ReadAll(resp.Body)                         //读取resp对象中的数据
				resp.Body.Close() //关闭数据流
				checkErr(err)
				//fmt.Println(string(result))
				fmt.Println(<-c)
			}
		}(channo)
	}

	for i := range UserArray {
		channo <- UserArray[i]
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
