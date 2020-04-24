package dao

import "time"

// CourseDirectory [...]
type CourseDirectory struct {
	ID           int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Csid         int64     `gorm:"index;column:csid;type:int(11)" json:"csid"`
	Title        string    `gorm:"column:title;type:varchar(45)" json:"title"`
	BeginTime    time.Time `gorm:"column:begin_time;type:timestamp" json:"begin_time"`
	EndTime      time.Time `gorm:"column:end_time;type:timestamp" json:"end_time"`
	Tid          int64     `gorm:"column:tid;type:bigint(20)" json:"tid"`
	CsType       int64     `gorm:"column:cs_type;type:int(11)" json:"cs_type"`
	ExamDuration int64     `gorm:"column:exam_duration;type:int(11)" json:"exam_duration"`
	Cid          int64     `gorm:"index;column:cid;type:int(11)" json:"cid"`
	Grade        int64     `gorm:"column:grade;type:int(11)" json:"grade"`
	Subject      int64     `gorm:"column:subject;type:int(11)" json:"subject"`
}

// Teacher [...]
type Teacher struct {
	ID   int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Tid  int64  `gorm:"index;column:tid;type:bigint(20)" json:"tid"`
	Name string `gorm:"column:name;type:varchar(45)" json:"name"`
	Desc string `gorm:"column:desc;type:varchar(1024)" json:"desc"`
	Pic  string `gorm:"column:pic;type:varchar(1024)" json:"pic"`
}

// CourseInfo [...]
type CourseInfo struct {
	ID         int64     `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Cid        int64     `gorm:"index;column:cid;type:int(11) unsigned;not null" json:"cid"`
	CourseName string    `gorm:"column:course_name;type:varchar(45)" json:"course_name"`
	BeginTime  time.Time `gorm:"column:begin_time;type:timestamp" json:"begin_time"`
	EndTime    time.Time `gorm:"column:end_time;type:timestamp" json:"end_time"`
	Grade      int64     `gorm:"column:grade;type:int(11)" json:"grade"`
	Subject    int64     `gorm:"column:subject;type:int(11)" json:"subject"`
	ApplyNum   int64     `gorm:"column:apply_num;type:int(11)" json:"apply_num"`
	PreAmount  int64     `gorm:"column:pre_amount;type:int(11)" json:"pre_amount"`
	AfAmount   int64     `gorm:"column:af_amount;type:int(11)" json:"af_amount"`
}
