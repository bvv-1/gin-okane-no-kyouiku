package models

import (
	"gin-okane-no-kyouiku/db"

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
	if err := db.GetDB().Debug().Model(&Goal{}).Order("created_at desc").First(&goal).Error; err != nil {
		return nil, err
	}
	return &goal, nil
}

func InsertGoalAndTasks(goal *Goal, tasks []Task) error {
	db.GetDB().Debug().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Goal{}).Create(&goal).Error; err != nil {
			return err
		}

		for i := range tasks {
			tasks[i].GoalID = goal.ID
		}

		if err := tx.Model(&Task{}).Create(&tasks).Error; err != nil {
			return err
		}
		return nil
	})
	return nil
}

// func GetInProgressGoalID() (uint, error) {
// 	var goal Goal
// 	if err := db.GetDB().Debug().Where("status = ?", 1).First(&goal).Error; err != nil {
// 		return 0, err
// 	}
// 	return goal.ID, nil
// }
