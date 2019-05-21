package types

import (
	"math/big"
	"time"
)

type User struct {
	Id             int64          `json:"id"                    gorm:"AUTO_INCREMENT;PRIMARY_KEY"          form:"id"`
	Address        string         `json:"address,omitempty"     gorm:"column:addr;type:varchar(43);unique" form:"address" validate:"required"`
	Email          string         `json:"email,omitempty"       gorm:"type:varchar(80)"                    form:"email"   validate:"email"`
	Skills         []Skill        `json:"skills,omitempty"      gorm:"many2many:statements;"                              validate:"-"`
	Missions       []Mission      `json:"missions,omitempty"                                                              validate:"-"`
	MissionSummary MissionSummary `json:"mission_summary,omitempty"`
	UpdateTime     *time.Time     `json:"update_time,omitempty" gorm:"column:updatetime;type:DATETIME"                    validate:"-"`
}

type MissionSummary struct {
	Publish         int64      `json:"publish,omitempty"`
	PaidRewardDET   *big.Float `json:"paid_reward_det,omitempty"`
	Submit          int64      `json:"submit,omitempty"`
	Accept          int64      `json:"accept,omitempty"`
	EarnedRewardDET *big.Float `json:"earned_reward_det,omitempty"`
	LastActive      *time.Time `json:"last_active,omitempty"`
}
