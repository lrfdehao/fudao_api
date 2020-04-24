package dao

import (
	"log"
	"time"
)

type CourseCountResult struct {
	Subject int64
	Count   int
}

type CourseDetailResult struct {
	ID           int
	Csid         int64
	Title        string
	BeginTime    time.Time
	EndTime      time.Time
	CsType       int64
	ExamDuration int64
	Grade        int64
	Subject      int64
	Teacher      *TeacherResult
	CourseInfo   *CourseInfoResult
}

type CourseInfoResult struct {
	ID         int64
	Cid        int64
	CourseName string
	BeginTime  time.Time
	EndTime    time.Time
	Grade      int64
	Subject    int64
	ApplyNum   int64
	PreAmount  int64
	AfAmount   int64
}

func GetCourseCount(beginTime, endTime time.Time) []*CourseCountResult {
	var result []*CourseCountResult

	err := DbCli.Table("course_directory").Select("subject, count(*) as count").Where("begin_time between ? and ?", beginTime, endTime).Group("subject").Find(&result).Error

	if err != nil {
		log.Print("GetCourseCount error ", err)
	} else {
		log.Printf("GetCourseCount success, resultSize[%v], result[%v]", len(result), result)
	}

	return result
}

func GetCourseDetail(beginTime, endTime time.Time, subject int64) []*CourseDetailResult {
	var courseDirectoryList []*CourseDirectory

	err := DbCli.Where("subject = ? and begin_time between ? and ?", subject, beginTime, endTime).Find(&courseDirectoryList).Error

	if err != nil {
		log.Print("GetCourseDetail error ", err)
	} else {
		log.Printf("GetCourseDetail success, courseDirectoryListSize[%v], courseDirectoryList[%v]", len(courseDirectoryList), courseDirectoryList)
	}

	var result []*CourseDetailResult
	for _, courseDirectory := range courseDirectoryList {
		result = append(result, courseDirectory.toResult())
	}

	return result
}

func GetCouseInfoByCid(cid int64) *CourseInfoResult {
	var courseInfo CourseInfo

	err := DbCli.Where("cid = ?", cid).First(&courseInfo).Error

	if err != nil {
		log.Print("GetCouseInfoByCid error ", err)
	} else {
		log.Printf("GetCouseInfoByCid success, courseInfo[%v]", courseInfo)
	}

	return courseInfo.toResult()
}

func (ci *CourseInfo) toResult() *CourseInfoResult {

	result := CourseInfoResult{}

	result.ID = ci.ID
	result.Cid = ci.Cid
	result.CourseName = ci.CourseName
	result.BeginTime = ci.BeginTime
	result.EndTime = ci.EndTime
	result.Grade = ci.Grade
	result.Subject = ci.Subject
	result.ApplyNum = ci.ApplyNum
	result.PreAmount = ci.PreAmount
	result.AfAmount = ci.AfAmount

	return &result
}

func (cd *CourseDirectory) toResult() *CourseDetailResult {
	cdResult := CourseDetailResult{}

	cdResult.ID = cd.ID
	cdResult.Csid = cd.Csid
	cdResult.Title = cd.Title
	cdResult.BeginTime = cd.BeginTime
	cdResult.EndTime = cd.EndTime
	cdResult.CsType = cd.CsType
	cdResult.ExamDuration = cd.ExamDuration
	cdResult.Grade = cd.Grade
	cdResult.Subject = cd.Subject
	cdResult.Teacher = GetTeacherByTid(cd.Tid)
	cdResult.CourseInfo = GetCouseInfoByCid(cd.Cid)

	return &cdResult
}
