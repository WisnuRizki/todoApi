package modules

import (
	"time"
)


type Activity struct {
	ID        	uint      		`json:"id"`
	Title 		string    		`gorm:"not null" json:"title"`
	Email     	string    		`json:"email"`
	Todos 		[]Todo			`gorm:"foreignKey:ActivityId;;constraint:OnDelete:CASCADE" json:"todo,omitempty"`
	CreatedAt 	time.Time 		`json:"created_at"`
	UpdatedAt 	time.Time 		`json:"updated_at"`
	DeletedAt 	*time.Time 		`gorm:"index" json:"deleted_at"`
}

type Todo struct {
	ID        	uint      		`json:"id"`
	Title 		string    		`json:"title"`
	IsActive  	bool    		`gorm:"default:true;" json:"is_active"`
	Priority    PriorityEnum    `gorm:"type:enum('low', 'medium', 'high', 'very-low','very-high');default:'very-high';check:priority IN ('low', 'medium', 'high','very-low','very-high')" json:"priority"`
	ActivityId  uint			`json:"activity_group_id"`
	CreatedAt 	time.Time 		`json:"created_at"`
	UpdatedAt 	time.Time 		`json:"updated_at"`
	DeletedAt 	*time.Time 		`gorm:"index" json:"deleted_at"`
}

type PriorityEnum string

const (
	VeryLowPriority    	PriorityEnum = "very-low"
    LowPriority    		PriorityEnum = "low"
    MediumPriority 		PriorityEnum = "medium"
    HighPriority   		PriorityEnum = "high"
	VeryHighPriority    PriorityEnum = "very-high"
)








