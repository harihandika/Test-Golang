package models

import "time"

type Todo struct {
	ID              int       `json:"id"`
	ActivityGroupId string    `json:"activity_group_id" form:"activity_group_id"`
	Title           string    `json:"title" form:"title"`
	Priority        string    `json:"priority" form:"priority"`
	IsActive        bool      `json:"is_active"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
