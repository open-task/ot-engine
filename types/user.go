package types

import "time"

type User struct {
	Id         int64      `json:"id"                    gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Address    string     `json:"address,omitempty"     gorm:"column:addr;type:varchar(43);unique"`
	Email      string     `json:"email,omitempty"       gorm:"type:varchar(80)"`
	Skills     []Skill    `json:"skills,omitempty"      gorm:"many2many:statements;"`
	UpdateTime *time.Time `json:"update_time,omitempty" gorm:"column:updatetime;type:DATETIME"`
}
