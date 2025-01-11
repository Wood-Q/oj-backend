package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Title       string `json:"title" grom:"comment:'问题标题'"`
	Content     string `json:"content" grom:"comment:'问题内容'"`
	Tags        string `json:"tags" grom:"comment:'问题标签'"`
	Answer      string `json:"answer" grom:"comment:'问题答案'"`
	SubmitNum   int    `json:"submitnum" grom:"comment:'提交次数'"`
	AcceptedNum int    `json:"acceptednum" grom:"comment:'题目通过次数'"`
}
