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
	QuestionID  int    `json:"question_id"`
	Language    string `json:"language"`
	Code        string `json:"code"`
	Status      string `json:"status"`
	ExitStatus  int    `json:"exit_status"`
	Time        int64  `json:"time"`
	Memory      int64  `json:"memory"`
	RunTime     int64  `json:"run_time"`
	Stdout      string `json:"stdout"`
	Stderr      string `json:"stderr"`
}
