package models

type Task struct {
	Model
	Name   string `json:"name"`
	Point  int    `json:"point"`
	GoalID uint   `json:"goal_id"`
}

type TaskResponse struct {
	Name  string `json:"name"`
	Point int    `json:"point"`
}
