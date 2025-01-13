package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

// 题目表单
type Question struct {
	gorm.Model
	// ID          uint `gorm:"primarykey"`
	// CreatedAt   time.Time
	// UpdatedAt   time.Time
	// DeletedAt   time.Time `gorm:"index"`
	QuestionID  uint            `json:"question_id" grom:"index;comment:'题目id'"`
	Title       string          `json:"title" grom:"comment:'问题标题'"`
	Content     string          `json:"content" grom:"comment:'问题内容'"`
	Tags        json.RawMessage `json:"tags" grom:"comment:'问题标签'"`
	Answer      string          `json:"answer" grom:"comment:'问题答案'"`
	SubmitNum   int             `json:"submitnum" grom:"comment:'提交次数'"`
	AcceptedNum int             `json:"acceptednum" grom:"comment:'题目通过次数'"`
	JudgeCase   string          `json:"judge_case" gorm:"comment:'判题用例'"`
	JudgeConfig string          `json:"judge_config" grom:"comment:'判题配置'"`
	UserID      uint            `json:"user_id" grom:"index;comment:'用户id'"`
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
