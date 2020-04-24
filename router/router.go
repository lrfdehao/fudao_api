package router

import (
	"crawler_api/handler"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	r.GET("/course/count", handler.GetCourseCount)
	r.GET("/course/detail", handler.GetCourseDetail)
}
