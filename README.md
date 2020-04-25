# FUDAO api

fudao 平台后台代码

## 技术选型

- web 开发框架： `go-gin`。 
- 数据库： `MySQL` + `gorm`。

爬虫地址: [fudao_crawler](https://github.com/lrfdehao/fudao_frontend)

前端地址: [fudao_frontend](https://github.com/lrfdehao/fudao_frontend)

## 目录结构

```
fudao_api
├── fudao_api
├── dao
│   ├── course.go
│   ├── db.go
│   ├── model.go
│   └── teacher.go
├── handler
│   ├── base.go
│   └── course_handler.go
├── router
│   └── router.go
├── go.mod
├── go.sum
└── main.go
```

## 接口文档

封装公共返回数据 `Res` :

```
type Res struct {
	Retcode int         `json:"retcode"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func res(code int, desc string, data interface{}) Res {
	var res Res
	res.Retcode = code
	res.Msg = desc
	res.Data = data
	return res
}
```

### 结构说明

属性    | 类型        | 说明
--      |-            | -
retcode | int         | 错误码
msg     | string      | 错误信息
data    | interface{} | 响应具体数据



### 查看每日科目课程数量

#### HTTP调用

##### 请求地址
> GET http://localhost:8888/course/count?beginTime=BEGINTIME&endTime=ENDTIME

##### 请求参数

属性      | 类型   | 默认 | 必填 | 说明
--        | -      | -    | -    | -         
beginTime | string |      | 是   | 开始日期。格式为 yyyy-mm-dd
endTime   | string |      | 是   | 结束日期。格式为 yyyy-mm-dd

##### 返回值
##### Object
返回JSON数据包

属性    | 类型           | 说明
--      | -              | -
retcode | int            | 错误码
msg     | string         | 错误信息
data    | Array.Object | 课程科目信息数量列表

data的结构

属性    | 类型 | 说明
--      | -    | -
Subject | int  | 科目code
Count   | int  | 开课数量

### 查看每日科目课程详情

#### HTTP调用
##### 请求地址
> Get http://localhost:8888/course/detail?beginTime=BEGINTIME&endTime=ENDTIME&subject=SUBJECT

##### 请求参数
属性      | 类型   | 默认 | 必填 | 说明
--        | -      | -    | -    | -
beginTime | string |      | 是   | 开始日期。格式为 yyyy-mm-dd
endTime   | string |      | 是   | 结束日期。格式为 yyyy-mm-dd
subject   | int    |      | 是   | 科目code
##### 返回值
##### Object
属性    | 类型           | 说明
--      | -              | -
retcode | int            | 错误码
msg     | string         | 错误信息
data    | Array.Object | 课程详情信息

data的结构
属性         | 类型   | 说明
--           | -      |- 
ID           | int    | 课程单条目录ID
Csid         | int    | fudao-课程目录唯一ID
Title        | string | 课程标题
BeginTime    | string | 开始日期。例: "2020-04-15T08:00:00+08:00"
EndTime      | string | 结束日期。例: "2020-04-15T08:00:00+08:00"
CsType       | int    | 目录类型，目前 0 代表课程， 1 代表考试
ExamDuration | int    |
Grade        | int    | 所属年级，0代表测试数据
Subject      | int    | 科目code
Teacher      | Object | 教师详情
CourseInfo   | Object | 所属课程包详情

Teacher的结构
属性 | 类型   | 说明
--   |-       |-
ID   | int    | 教师ID
Tid  | int    | fudao-教师唯一ID
Name | string | 教师名称
Desc | string | 教师介绍
Pic  | string | 教师头像

CourseInfo的结构
属性      | 类型   | 说明
--        | -      | -
ID        | int    | 课程包ID
Cid       | int    | fudao-课程包唯一ID
CourseName| string | 课程包名称
BeginTime | string | 开始日期。例: "2020-04-15T08:00:00+08:00"
EndTime   | string | 结束日期。例: "2020-04-15T08:00:00+08:00"
Grade     | int    | 所属年级，0代表测试数据
Subject   | int    | 科目code
ApplyNum  | int    | 报名人数
PreAmount | int    | 原价， 单位: 分
AfAmount  | int    | 现价， 单位: 分

#### retcode的合法值
Teacher的结构
值   | 说明                        | 最低版本
--   |-                            | -
0    | 正常                        | 
1001 | beginTime 或者 endTime 错误 | 
1002 | subject 错误                | 

#### 返回数据示例
```
{
    "retcode": 0,
    "msg": "ok",
    "data": [
        {
            "ID": 86,
            "Csid": 272636,
            "Title": "第一节",
            "BeginTime": "2020-04-15T08:00:00+08:00",
            "EndTime": "2020-04-15T09:20:00+08:00",
            "CsType": 0,
            "ExamDuration": 0,
            "Grade": 0,
            "Subject": 6001,
            "Teacher": {
                "ID": 8,
                "Tid": 882000201,
                "Name": "肖主讲",
                "Desc": "我是一名英语老师，跟我学习英语吧",
                "Pic": "http://pub.idqqimg.com/pc/misc/files/20191127/0cd1c693d1e84a38b0d9b5d4a822c326.jpg"
            },
            "CourseInfo": {
                "ID": 28,
                "Cid": 212187,
                "CourseName": "lyh课题流量课现网003",
                "BeginTime": "2020-04-15T08:00:00+08:00",
                "EndTime": "2020-04-15T09:20:00+08:00",
                "Grade": 0,
                "Subject": 6001,
                "ApplyNum": 0,
                "PreAmount": 1,
                "AfAmount": 0
            }
        },
        {
            "ID": 2505,
            "Csid": 274835,
            "Title": "17",
            "BeginTime": "2020-04-15T09:17:05+08:00",
            "EndTime": "2020-04-15T17:16:06+08:00",
            "CsType": 0,
            "ExamDuration": 0,
            "Grade": 0,
            "Subject": 6001,
            "Teacher": {
                "ID": 14,
                "Tid": 552297168,
                "Name": "陈俊文",
                "Desc": "积极向上",
                "Pic": "http://pub.idqqimg.com/pc/misc/files/20190312/4d3d1f269f124b6697b48ae16ec45b7d.jpg"
            },
            "CourseInfo": {
                "ID": 263,
                "Cid": 205699,
                "CourseName": "javen语文课测试课",
                "BeginTime": "2020-03-17T10:50:28+08:00",
                "EndTime": "2020-04-17T17:26:55+08:00",
                "Grade": 0,
                "Subject": 6001,
                "ApplyNum": 0,
                "PreAmount": 0,
                "AfAmount": 0
            }
        },
        {
            "ID": 2506,
            "Csid": 274922,
            "Title": "18",
            "BeginTime": "2020-04-15T19:21:52+08:00",
            "EndTime": "2020-04-15T23:59:59+08:00",
            "CsType": 0,
            "ExamDuration": 0,
            "Grade": 0,
            "Subject": 6001,
            "Teacher": {
                "ID": 14,
                "Tid": 552297168,
                "Name": "陈俊文",
                "Desc": "积极向上",
                "Pic": "http://pub.idqqimg.com/pc/misc/files/20190312/4d3d1f269f124b6697b48ae16ec45b7d.jpg"
            },
            "CourseInfo": {
                "ID": 263,
                "Cid": 205699,
                "CourseName": "javen语文课测试课",
                "BeginTime": "2020-03-17T10:50:28+08:00",
                "EndTime": "2020-04-17T17:26:55+08:00",
                "Grade": 0,
                "Subject": 6001,
                "ApplyNum": 0,
                "PreAmount": 0,
                "AfAmount": 0
            }
        },
        {
            "ID": 8880,
            "Csid": 272470,
            "Title": "第一节",
            "BeginTime": "2020-04-15T11:00:00+08:00",
            "EndTime": "2020-04-15T12:00:00+08:00",
            "CsType": 0,
            "ExamDuration": 0,
            "Grade": 6001,
            "Subject": 6001,
            "Teacher": {
                "ID": 8,
                "Tid": 882000201,
                "Name": "肖主讲",
                "Desc": "我是一名英语老师，跟我学习英语吧",
                "Pic": "http://pub.idqqimg.com/pc/misc/files/20191127/0cd1c693d1e84a38b0d9b5d4a822c326.jpg"
            },
            "CourseInfo": {
                "ID": 807,
                "Cid": 213163,
                "CourseName": "xzy课题流量测试课009",
                "BeginTime": "2020-04-15T11:00:00+08:00",
                "EndTime": "2020-04-15T12:00:00+08:00",
                "Grade": 6001,
                "Subject": 6001,
                "ApplyNum": 0,
                "PreAmount": 0,
                "AfAmount": 0
            }
        },
        {
            "ID": 8881,
            "Csid": 272472,
            "Title": "第一节",
            "BeginTime": "2020-04-15T11:00:00+08:00",
            "EndTime": "2020-04-15T12:00:00+08:00",
            "CsType": 0,
            "ExamDuration": 0,
            "Grade": 6001,
            "Subject": 6001,
            "Teacher": {
                "ID": 8,
                "Tid": 882000201,
                "Name": "肖主讲",
                "Desc": "我是一名英语老师，跟我学习英语吧",
                "Pic": "http://pub.idqqimg.com/pc/misc/files/20191127/0cd1c693d1e84a38b0d9b5d4a822c326.jpg"
            },
            "CourseInfo": {
                "ID": 808,
                "Cid": 213164,
                "CourseName": "xzy课题流量测试课010",
                "BeginTime": "2020-04-15T11:00:00+08:00",
                "EndTime": "2020-04-15T12:00:00+08:00",
                "Grade": 6001,
                "Subject": 6001,
                "ApplyNum": 0,
                "PreAmount": 1,
                "AfAmount": 0
            }
        },
        {
            "ID": 8882,
            "Csid": 272498,
            "Title": "咏鹅",
            "BeginTime": "2020-04-15T16:00:00+08:00",
            "EndTime": "2020-04-15T18:00:00+08:00",
            "CsType": 0,
            "ExamDuration": 0,
            "Grade": 6001,
            "Subject": 6001,
            "Teacher": {
                "ID": 8,
                "Tid": 882000201,
                "Name": "肖主讲",
                "Desc": "我是一名英语老师，跟我学习英语吧",
                "Pic": "http://pub.idqqimg.com/pc/misc/files/20191127/0cd1c693d1e84a38b0d9b5d4a822c326.jpg"
            },
            "CourseInfo": {
                "ID": 809,
                "Cid": 213170,
                "CourseName": "xzy课题流量测试课011",
                "BeginTime": "2020-04-15T16:00:00+08:00",
                "EndTime": "2020-04-20T13:05:00+08:00",
                "Grade": 6001,
                "Subject": 6001,
                "ApplyNum": 0,
                "PreAmount": 1,
                "AfAmount": 0
            }
        },
        {
            "ID": 8918,
            "Csid": 265475,
            "Title": "第一节课",
            "BeginTime": "2020-04-15T18:02:59+08:00",
            "EndTime": "2020-04-15T20:03:00+08:00",
            "CsType": 0,
            "ExamDuration": 0,
            "Grade": 6001,
            "Subject": 6001,
            "Teacher": {
                "ID": 8,
                "Tid": 882000201,
                "Name": "肖主讲",
                "Desc": "我是一名英语老师，跟我学习英语吧",
                "Pic": "http://pub.idqqimg.com/pc/misc/files/20191127/0cd1c693d1e84a38b0d9b5d4a822c326.jpg"
            },
            "CourseInfo": {
                "ID": 817,
                "Cid": 212190,
                "CourseName": "xzy课题流量测试课003",
                "BeginTime": "2020-04-15T18:02:59+08:00",
                "EndTime": "2020-04-20T15:00:00+08:00",
                "Grade": 6001,
                "Subject": 6001,
                "ApplyNum": 0,
                "PreAmount": 0,
                "AfAmount": 0
            }
        },
        {
            "ID": 9218,
            "Csid": 274907,
            "Title": "数据直播",
            "BeginTime": "2020-04-15T14:45:00+08:00",
            "EndTime": "2020-04-15T15:45:00+08:00",
            "CsType": 0,
            "ExamDuration": 0,
            "Grade": 6001,
            "Subject": 6001,
            "Teacher": {
                "ID": 8,
                "Tid": 882000201,
                "Name": "肖主讲",
                "Desc": "我是一名英语老师，跟我学习英语吧",
                "Pic": "http://pub.idqqimg.com/pc/misc/files/20191127/0cd1c693d1e84a38b0d9b5d4a822c326.jpg"
            },
            "CourseInfo": {
                "ID": 850,
                "Cid": 194536,
                "CourseName": "xzy课中练习验证",
                "BeginTime": "2020-02-13T21:03:59+08:00",
                "EndTime": "2020-04-15T15:45:00+08:00",
                "Grade": 6001,
                "Subject": 6001,
                "ApplyNum": 0,
                "PreAmount": 0,
                "AfAmount": 0
            }
        }
    ]
}
```