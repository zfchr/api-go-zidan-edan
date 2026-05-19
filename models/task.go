package models

import "time"

type TaskStatus string

const (
	StatusTodo     TaskStatus = "todo"
	StatusProgress TaskStatus = "progress"
	StatusDone     TaskStatus = "done"
)

type Task struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status" gorm:"type:enum('todo','progress','done');default:'todo'"`
	DueDate     time.Time  `json:"due_date"`
}
