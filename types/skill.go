package types

import "time"

type Skill struct {
	Id         int64      `json:"id"                    gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Skill      string     `json:"skill,omitempty"       gorm:"type:varchar(200)"`
	Users      []User     `json:"users,omitempty"       gorm:"many2many:statements;"`
	Claim      int        `json:"claim,omitempty"`
	Submit     int        `json:"submit,omitempty"`
	Confirm    int        `json:"confirm,omitempty"`
	UpdateTime *time.Time `json:"update_time,omitempty" gorm:"column:updatetime;type:DATETIME"`
}
