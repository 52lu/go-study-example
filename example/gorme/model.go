/**
 * @Description 定义model
 **/
package gorme

import (
	"database/sql"
	"gorm.io/gorm"
)

// 用户模型
type User struct {
	gorm.Model
	NickName     string         `gorm:"type:varchar(20);not null;default:'';comment:昵称"`
	Age          uint8          `gorm:"size:4;comment:年龄"`
	Phone        sql.NullString `gorm:"type:char(11);not null;default:'';unique;comment:手机号"`
	Birthday     string         `gorm:"type:varchar(10);not null;default:'';comment:生日"`
	MemberNumber sql.NullString `gorm:"type:varchar(20);not null;default:'';comment:会员编号"`
	ActivatedAt  sql.NullTime   `gorm:"comment:激活时间"`
}
