package dao

import "log"

type TeacherResult struct {
	ID   int
	Tid  int64
	Name string
	Desc string
	Pic  string
}

func GetTeacherByTid(tid int64) *TeacherResult {

	var teacher Teacher

	err := DbCli.Where("tid = ?", tid).First(&teacher).Error

	if err != nil {
		log.Print("GetTeacherByTid error ", err)
	} else {
		log.Printf("GetTeacherByTid success, result[%v]", teacher)
	}

	return teacher.toResult()
}

func (t *Teacher) toResult() *TeacherResult {
	tResult := TeacherResult{}

	tResult.ID = t.ID
	tResult.Tid = t.Tid
	tResult.Name = t.Name
	tResult.Desc = t.Desc
	tResult.Pic = t.Pic

	return &tResult
}
