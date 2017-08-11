package main

import (
	//sql包提供了保证SQL或类SQL数据库的泛用接口。
	"database/sql"

	"math/rand"

	"time"

	//mt包实现了类似C语言printf和scanf的格式化I/O。
	"fmt"

	//sql驱动
	_ "github.com/Go-SQL-Driver/MySQL"
)

const (
	KC_RAND_KIND_ALL = 3 // 数字、大小写字母
)

func Krand(size int, kind int) []byte {

	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)

	is_all := kind > 2 || kind < 0

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {

		if is_all { // random ikind

			ikind = rand.Intn(3)

		}

		scope, base := kinds[ikind][0], kinds[ikind][1]

		result[i] = uint8(base + rand.Intn(scope))

	}

	return result
}

func main() {

	//sql.Open()函数用来打开一个注册过的数据库驱动，Go-MySQL-Driver中注册了mysql这个数据库驱动，第二个参数是DNS(Data Source Name)，它是Go-MySQL-Driver定义的一些数据库链接和配置信息；user:password@/dbname
	db, err := sql.Open("mysql", "root:123456@/mysql_pythoncz?charset=utf8")
	checkErr(err)

	//db.Prepare()函数用来返回准备要执行的sql操作，然后返回准备完毕的执行状态。
	stmt, err := db.Prepare("INSERT xingzhen001 SET name=?,address=?")
	checkErr(err)

	for i := 0; i < 200; i++ {
		//stmt.Exec()函数用来执行stmt准备好的SQL语句
		res, err := stmt.Exec(string(Krand(6, KC_RAND_KIND_ALL)), string(Krand(12, KC_RAND_KIND_ALL)))
		checkErr(err)

		//
		id, err := res.LastInsertId()
		checkErr(err)

		fmt.Println(id)
	}

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}

}
