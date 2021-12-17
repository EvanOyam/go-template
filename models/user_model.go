package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    sql.NullString `gorm:"size:32;unique;" json:"username" form:"username"` // 用户名
	Password    string         `gorm:"size:512" json:"password" form:"password"`        // 密码
	Nickname    string         `gorm:"size:16;" json:"nickname" form:"nickname"`        // 昵称
	Description string         `gorm:"type:text" json:"description" form:"description"` // 个人描述
}
