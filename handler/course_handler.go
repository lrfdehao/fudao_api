package handler

import (
	"crawler_api/dao"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetCourseCount(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	beginTime, err := time.Parse("2006-01-02", c.Query("beginTime"))
	if err != nil {
		c.JSON(400, res(1001, "beginTime param error", ""))
		return
	}

	endTime, err := time.Parse("2006-01-02", c.Query("endTime"))
	if err != nil {
		c.JSON(400, res(1001, "endTime param error", ""))
		return
	}

	resList := dao.GetCourseCount(beginTime, endTime)
	c.JSON(200, res(0, "ok", resList))
}

func GetCourseDetail(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	beginTime, err := time.Parse("2006-01-02", c.Query("beginTime"))
	if err != nil {
		c.JSON(400, res(1001, "beginTime param error", ""))
		return
	}

	endTime, err := time.Parse("2006-01-02", c.Query("endTime"))
	if err != nil {
		c.JSON(400, res(1001, "endTime param error", ""))
		return
	}

	subjectStr := c.Query("subject")
	subject, err := strconv.ParseInt(subjectStr, 10, 64)
	if err != nil {
		c.JSON(400, res(1002, "subjectStr param error", ""))
		return
	}

	resList := dao.GetCourseDetail(beginTime, endTime, subject)

	c.JSON(200, res(0, "ok", resList))
}
