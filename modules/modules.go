package modules

import (
	"time"

)


type Activity struct {
	ID        	uint      		`gorm:"column:activity_id" json:"id"`
	Title 		string    		`gorm:"not null;" json:"title"`
	Email     	string    		`json:"email"`
	Todos 		[]Todo			`gorm:"foreignKey:ActivityId;;constraint:OnDelete:CASCADE" json:"todo,omitempty"`
	CreatedAt 	time.Time 		`json:"created_at"`
	UpdatedAt 	time.Time 		`json:"updated_at"`
	DeletedAt 	*time.Time 		`json:"deleted_at"`
}

type Todo struct {
	ID        	uint      		`gorm:"column:todo_id" json:"id"`
	ActivityId  uint			`json:"activity_group_id"`
	Title 		string    		`gorm:"not null;" json:"title"`
	Priority    PriorityEnum    `gorm:"column:priority;type:enum('low', 'medium', 'high', 'very-low','very-high');default:'very-high';check:priority IN ('low', 'medium', 'high','very-low','very-high')" json:"priority"`
	CreatedAt 	time.Time 		`json:"created_at"`
	IsActive  	bool    		`gorm:"default:true;" json:"is_active"`
	UpdatedAt 	time.Time 		`json:"updated_at"`
	DeletedAt 	*time.Time 		`json:"deleted_at,omitempty"`
}

type PriorityEnum string

const (
	VeryLowPriority    	PriorityEnum = "very-low"
    LowPriority    		PriorityEnum = "low"
    MediumPriority 		PriorityEnum = "medium"
    HighPriority   		PriorityEnum = "high"
	VeryHighPriority    PriorityEnum = "very-high"
)








