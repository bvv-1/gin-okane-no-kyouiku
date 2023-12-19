package models

import (
	"gin-okane-no-kyouiku/db"

	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

type Goal struct {
	Model
	Name   string `json:"name"`
	Point  int    `json:"point"`
	Status int    `json:"status"`
	// UserID uint   `json:"user_id"`
}

func GetGoal() (*Goal, error) {
	var goal Goal
	if err := db.GetDB().Debug().First(&goal).Error; err != nil {
		return nil, xerrors.Errorf("failed to get goal: %w", err)
	}
	return &goal, nil
}

func InsertGoalAndTasks(goal *Goal, tasks []Task) error {
	db.GetDB().Debug().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&goal).Error; err != nil {
			return err
		}
		for _, task := range tasks {
			task.GoalID = goal.ID
			if err := tx.Create(&task).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return nil
}
