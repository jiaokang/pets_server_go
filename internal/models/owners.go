package models

import (
	"time"
)

type Owner struct {
	ID         int        `gorm:"primaryKey;autoIncrement;comment:主人唯一标识"`
	Name       string     `gorm:"size:50;not null;comment:用户名"`
	NickName   string     `gorm:"size:64;comment:昵称"`
	Phone      string     `gorm:"size:20;uniqueIndex;comment:主人手机号"`
	Email      string     `gorm:"size:100;uniqueIndex;comment:主人邮箱"`
	Address    string     `gorm:"type:text;comment:主人地址"`
	CreateTime time.Time  `gorm:"default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime *time.Time `gorm:"comment:更新时间"`
	Pwd        string     `gorm:"size:64;comment:密码"`
}

// 自定义表名
func (Owner) TableName() string {
	return "business_owners"
}
