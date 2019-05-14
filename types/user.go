package types

import "time"

type User struct {
	Id         int64      `json:"id"                    gorm:"AUTO_INCREMENT;PRIMARY_KEY"          form:"id"`
	Address    string     `json:"address,omitempty"     gorm:"column:addr;type:varchar(43);unique" form:"address" validate:"required"`
	Email      string     `json:"email,omitempty"       gorm:"type:varchar(80)"                    form:"email"   validate:"email"`
	Skills     []Skill    `json:"skills,omitempty"      gorm:"many2many:statements;"                              validate:"-"`
	Missions   []Mission  `json:"missions,omitempty"`
	UpdateTime *time.Time `json:"update_time,omitempty" gorm:"column:updatetime;type:DATETIME"`
}
