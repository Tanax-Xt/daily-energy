package adapterModels

import (
	"github.com/google/uuid"
)

type UserWeightHistory struct {
	Id         uuid.UUID `gorm:"column:id;primaryKey"`
	Utgid      int64       `gorm:"column:utgid"`
	Date       int64       `gorm:"column:date"`
	UserWeight int         `gorm:"column:user_weight"`
	UserHeight int         `gorm:"column:user_height"`
}
