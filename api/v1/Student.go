package v1

import (
	"fmt"
	"gin-example/global/response"
	"gin-example/model"
	"gin-example/model/request"
	resp "gin-example/model/response"
	"gin-example/service"
	"github.com/gin-gonic/gin"
)

// @Tags Student
// @Summary 创建Student
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Student true "创建Student"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /student/createStudent [post]
func CreateStudent(c *gin.Context) {
	var student model.Student
	_ = c.ShouldBindJSON(&student)
	err := service.CreateStudent(student)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Student
// @Summary 删除Student
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Student true "删除Student"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /student/deleteStudent [delete]
func DeleteStudent(c *gin.Context) {
	var student model.Student
	_ = c.ShouldBindJSON(&student)
	err := service.DeleteStudent(student)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Student
// @Summary 更新Student
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Student true "更新Student"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /student/updateStudent [put]
func UpdateStudent(c *gin.Context) {
	var student model.Student
	_ = c.ShouldBindJSON(&student)
	err := service.UpdateStudent(&student)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Student
// @Summary 用id查询Student
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Student true "用id查询Student"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /student/findStudent [get]
func FindStudent(c *gin.Context) {
	var student model.Student
	_ = c.ShouldBindQuery(&student)
	err, restudent := service.GetStudent(student.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"restudent": restudent}, c)
	}
}

// @Tags Student
// @Summary 分页获取Student列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Student列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /student/getStudentList [get]
func GetStudentList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetStudentInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
