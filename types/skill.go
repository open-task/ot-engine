package types

import "time"

type Skill struct {
	Id         int64      `json:"id"                    gorm:"AUTO_INCREMENT;PRIMARY_KEY"      form:"id"`
	Tag        string     `json:"tag,omitempty"         gorm:"type:varchar(200)"               form:"tag"`
	Users      []User     `json:"users,omitempty"       gorm:"many2many:statements;"                           validate:"-"`
	Claim      int        `json:"claim,omitempty"                                              form:"claim"`
	Submit     int        `json:"submit,omitempty"                                             form:"submit"`
	Confirm    int        `json:"confirm,omitempty"                                            form:"confirm"`
	UpdateTime *time.Time `json:"update_time,omitempty" gorm:"column:updatetime;type:DATETIME"                 validate:"-"`
}
