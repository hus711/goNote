//专门用来存放数据模型的
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
	"time"
)

/*
 1.表创建好之后，表名、字段都变成小写的，驼峰法命名变成下划线
 2.
*/

// Scan  sql语句查询使用
type Result struct {
	Name string
	Age  int
}

//使用了自带的模型
type Person struct {
	gorm.Model        //使用了gorm自带的四个字段，id,创建、更新、删除时间
	Name       string `gorm:"unique:not null"` //不为空，唯一
	Age        int    `gorm:"type:varchar(100);not null"`
}

//没有使用自带的模型
type Animal struct {
	AnimalId int       `gorm:"column:beast_id;primary_key;AUTO_INCREMENT"` // set column name to `beast_id`
	Birthday time.Time `gorm:"column:day_of_the_beast"`                    // set column name to `day_of_the_beast`
	Age      int       `gorm:"column:age_of_the_beast"`                    // set column name to `age_of_the_beast`
	CanFly   bool      //演示驼峰命名  数据库中自动命名为 can_fly
}

//数据库表的命名
func (Person) TableName() string {
	return "person"
}

func (animal *Animal) String() string {
	return strconv.Itoa(animal.Age)
}

//在调用create之前会调用，先设置一个默认值给canfly(`gorm:"default:'galeone'"`,使用default之后，插入时候默认忽略这个字段.这时候设置这个无效)
func (animal *Animal) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CanFly", false)
	return nil
}

func (p *Person) AfterCreate() (err error) {
	fmt.Println("数据插入完成:", p.Name)
	return
}
func (p *Person) AfterDelete() (err error) {
	fmt.Println("数据删除完成:", p.Name)
	return
}

func (p *Person) AfterUpdate() (err error) {
	fmt.Println("数据更新完成:", p.Name)
	return
}

func (p *Person) AfterFind() (err error) {
	fmt.Println("数据查询完成:", p.Name)
	return
}
