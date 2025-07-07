package adapterModels

import (
	"github.com/google/uuid"
	"github.com/isiyar/daily-energy/backend/internal/domain/models"
)

type Action struct {
	Id           uuid.UUID         `gorm:"column:id;primaryKey"`
	Utgid        int64             `gorm:"column:utgid"`
	Date         int64             `gorm:"column:date"`
	ActivityName string            `gorm:"column:activity_name"`
	Calories     int               `gorm:"column:calories"`
	Type         models.ActionType `gorm:"column:type"`
}
