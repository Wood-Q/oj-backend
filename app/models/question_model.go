package models

import (
	"github.com/lib/pq"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Question 结构体
type Question struct {
	gorm.Model

	QuestionID  uint           `json:"question_id" gorm:"index;comment:'题目id'"`
	Title       string         `json:"title" gorm:"comment:'问题标题'"`
	Content     string         `json:"content" gorm:"comment:'问题内容'"`
	Tags        pq.StringArray `json:"tags" gorm:"type:text[];comment:'问题标签'"`
	Answer      string         `json:"answer" gorm:"comment:'问题答案'"`
	SubmitNum   int            `json:"submitnum" gorm:"comment:'提交次数'"`
	AcceptedNum int            `json:"acceptednum" gorm:"comment:'题目通过次数'"`
	JudgeCase   datatypes.JSON `json:"judge_case" gorm:"type:json;comment:'判题用例'"`
	JudgeConfig datatypes.JSON `json:"judge_config" gorm:"type:json;comment:'判题配置'"`
	UserID      uint           `json:"user_id" gorm:"index;comment:'用户id'"`
}

// 题目提交表
type QuestionSubmit struct {
	gorm.Model
	// ID         uint `gorm:"primarykey"`
	// CreatedAt  time.Time
	// UpdatedAt  time.Time
	// DeletedAt  time.Time `gorm:"index"`
	QuestionID uint   `json:"question_id" grom:"index;comment:'题目id'"`
	Content    string `json:"content" grom:"comment:'提交内容'"`
	Language   string `json:"language" grom:"comment:'提交语言'"`
	UserID     uint   `json:"user_id" grom:"index;comment:'提交用户id'"`
	Status     string `json:"status" grom:"comment:'判题状态'"`
	JudgeInfo  string `json:"judgeInfo" grom:"comment:'判题信息'"`
	Code       string `json:"code" grom:"comment:'提交代码'"`
}
