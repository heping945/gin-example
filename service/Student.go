package service

import (
	"gin-example/global"
	"gin-example/model"
	"gin-example/model/request"
)

// @title    CreateStudent
// @description   create a Student
// @param     student               model.Student
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateStudent(student model.Student) (err error) {
	err = global.GVA_DB.Create(&student).Error
	return err
}

// @title    DeleteStudent
// @description   delete a Student
// @auth                     （2020/04/05  20:22）
// @param     student               model.Student
// @return                    error

func DeleteStudent(student model.Student) (err error) {
	err = global.GVA_DB.Delete(student).Error
	return err
}

// @title    UpdateStudent
// @description   update a Student
// @param     student          *model.Student
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateStudent(student *model.Student) (err error) {
	err = global.GVA_DB.Save(student).Error
	return err
}

// @title    GetStudent
// @description   get the info of a Student
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Student        Student

func GetStudent(id uint) (err error, student model.Student) {
	err = global.GVA_DB.Where("id = ?", id).First(&student).Error
	return
}

// @title    GetStudentInfoList
// @description   get Student list by pagination, 分页获取用户列表
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetStudentInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	var students []model.Student
	err = db.Find(&students).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&students).Error
	return err, students, total
}
