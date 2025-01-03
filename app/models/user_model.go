package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserAccount  string `json:"useraccount" grom:"comment:'用户账户'"`
	UserPassword string `json:"password" grom:"comment:'用户密码'"`
	UnionId      string `json:"unionid" gorm:"'微信开放平台id'"`
	MpOpenId     string `json:"openid" gorm:"'公众号openid'"`
	UserName     string `json:"username" gorm:"comment:'用户名'"`
	UserAvatar   string `json:"useravatar" gorm:"comment:'用户头像'"`
	UserProfile  string `json:"userprofile" gorm:"comment:'用户简介'"`
	UserRole     string `json:"userrole" gorm:"comment:'用户角色,user/admin/ban'"`
}
