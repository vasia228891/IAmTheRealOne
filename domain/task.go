package domain

import "time"

type Task struct {
	Id          uint64
	User_id     uint64
	Title       string
	Description *string
	Status      TaskStatus
	Date        *time.Time
	Name        string
	Password    string
}

type TaskStatus string

const (
	Draft     TaskStatus = "Draft"
	Assigned  TaskStatus = "ASSIGNED"
	Completed TaskStatus = "COMPLETED"
)
