package models

import "github.com/lib/pq"

type Course struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	CourseCode    string         `json:"courseCode" gorm:"unique"`
	CourseTitle   string         `json:"courseTitle"`
	Prerequisites string         `json:"prerequisites"`
	Hubs          pq.StringArray `json:"hubs" gorm:"type:text[]"`
}

