// 自动生成模板Student
package model

import (
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	Name   string `json:"name" gorm:"column:name"`
	Gender bool   `json:"gender" gorm:"column:gender"`
	Age    int    `json:"age" gorm:"column:age"`
}
