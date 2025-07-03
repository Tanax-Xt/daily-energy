package models

type ActionType string

const (
	Food     ActionType = "Food"
	Activity ActionType = "Activity"
)

type Action struct {
	Id           string
	Utgid        int64
	Date         int64
	ActivityName string
	Calories     int
	Type         ActionType
}
