package router

import (
	"gin-example/api/v1"
	"github.com/gin-gonic/gin"
)

func InitStudentRouter(Router *gin.RouterGroup) {
	StudentRouter := Router.Group("student")
	{
		StudentRouter.POST("createStudent", v1.CreateStudent)   // 新建Student
		StudentRouter.DELETE("deleteStudent", v1.DeleteStudent) //删除Student
		StudentRouter.PUT("updateStudent", v1.UpdateStudent)    //更新Student
		StudentRouter.GET("findStudent", v1.FindStudent)        // 根据ID获取Student
		StudentRouter.GET("getStudentList", v1.GetStudentList)  //获取Student列表
	}
}
