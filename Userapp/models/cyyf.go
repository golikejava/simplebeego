package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type CYYF struct {
	Cid        int `orm:"pk;auto"`
	Name       string
	Address    string
	CreateTime time.Time
}

func init() {
	orm.RegisterModel(new(CYYF))
}

//拿数据（createtime为插入时间，其余数据是post传过来的数据），插入数据mysql
