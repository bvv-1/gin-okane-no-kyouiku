package models

import (
	"gin-okane-no-kyouiku/db"

	"gorm.io/gorm"
)

type Progress struct {
	Model
	Day    int    `json:"day"`
	GoalID uint   `json:"goal_id"`
	PlanID uint   `json:"plan_id"`
	TaskID uint   `json:"task_id"`
	Name   string `json:"name"`
	Point  int    `json:"point"`
	IsDone bool   `json:"is_done"`
}

type TaskAndStatus struct {
	Task   Task `json:"task"`
	IsDone bool `json:"is_done"`
}

func InsertProgress(day int, taskProgress []TaskAndStatus) error {
	err := db.GetDB().Debug().Transaction(func(tx *gorm.DB) error {
		var progresses []Progress
		for _, taskAndStatus := range taskProgress {
			var goal Goal
			if err := tx.Model(&Goal{}).Where("status = ?", 1).First(&goal).Error; err != nil {
				return err
			}

			var plan Plan
			if err := tx.Model(&Plan{}).Where("goal_id = ? AND day = ?", goal.ID, day).First(&plan).Error; err != nil {
				return err
			}

			var task Task
			if err := tx.Model(&Task{}).Where("goal_id = ? AND name = ?", goal.ID, taskAndStatus.Task.Name).First(&task).Error; err != nil {
				return err
			}

			progresses = append(progresses, Progress{
				Day:    day,
				GoalID: goal.ID,
				PlanID: plan.ID,
				TaskID: task.ID,
				Name:   task.Name,
				Point:  task.Point,
				IsDone: taskAndStatus.IsDone,
			})
		}

		if err := tx.Model(&Progress{}).Create(&progresses).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
