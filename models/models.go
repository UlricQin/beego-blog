package models

// package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Catalog struct {
	Id    int64
	Ident string `orm:"unique"`
	Name  string
}

type Blog struct {
	Id        int64
	Ident     string `orm:"unique"`
	Title     string
	Keywords  string `orm:"null"`
	Content   string `orm:"type(text)"`
	CatalogId int64  `orm:"index"`
	Type      int8   /*0:original, 1:translate, 2:reprint*/
	Status    int8   /*0:draft, 1:release*/
	Views     int64
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
}

type Option struct {
	Id    int
	Name  string
	Value string
}

func (*Catalog) TableEngine() string {
	return engine()
}

func (*Blog) TableEngine() string {
	return engine()
}

func (*Option) TableEngine() string {
	return engine()
}

func engine() string {
	return "INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci"
}

func init() {
	orm.RegisterModelWithPrefix("bb_", new(Catalog), new(Blog), new(Option))
}

// func main() {
// 	orm.RegisterDataBase("default", "mysql", "root:1234@/beego_blog?charset=utf8&loc=Asia%2FShanghai", 30, 200)
// 	orm.RunCommand()
// }
