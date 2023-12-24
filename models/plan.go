package models

import (
	"gin-okane-no-kyouiku/db"

	"gorm.io/gorm"
)

type Plan struct {
	Model
	Day    int  `json:"day"`
	TaskID uint `json:"task_id"`
	GoalID uint `json:"goal_id"`
}

type PlanResponse struct {
	Day        int            `json:"day"`
	TasksToday []TaskResponse `json:"tasks_today"`
}

type SuggestedPlan struct {
	Day        int            `json:"day"`
	PlansToday []TaskResponse `json:"plans_today"`
}

type DailyPlansResponse struct {
	Day        int            `json:"day"`
	PlansToday []TaskResponse `json:"plans_today"`
}

func GetSuggestedPlans() (*[]PlanResponse, error) {
	var planResponse []PlanResponse

	err := db.GetDB().Debug().Transaction(func(tx *gorm.DB) error {
		var goal Goal
		if err := tx.Model(&Goal{}).Order("created_at desc").First(&goal).Error; err != nil {
			return err
		}

		var tasks []Task
		if err := tx.Model(&Task{}).Where("goal_id = ?", goal.ID).Find(&tasks).Error; err != nil {
			return err
		}

		var plans []Plan = generatePlans(goal, tasks)
		if err := tx.Model(&Plan{}).Create(&plans).Error; err != nil {
			return err
		}

		for _, plan := range plans {
			planResponse = append(planResponse, ToPlanResponse(plan, tasks))
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &planResponse, nil
}

func generatePlans(goal Goal, tasks []Task) []Plan {
	// TODO: ここに高度でかっこいいアルゴリズムが入る、ランダム要素あり
	var plans = []Plan{
		{Day: 1, TaskID: tasks[0].ID, GoalID: goal.ID},
		{Day: 2, TaskID: tasks[0].ID, GoalID: goal.ID},
	}

	return plans
}

func AcceptSuggestedPlans() error {
	err := db.GetDB().Debug().Transaction(func(tx *gorm.DB) error {
		var plan Plan
		if err := tx.Model(&Plan{}).Order("created_at desc").First(&plan).Error; err != nil {
			return err
		}

		if err := tx.Model(&Goal{}).Where("id = ?", plan.GoalID).Update("status", 1).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func GetPlanByDay(day int) (*PlanResponse, error) {
	var planResponse PlanResponse
	var tasks []Task

	err := db.GetDB().Debug().Transaction(func(tx *gorm.DB) error {
		var goal Goal
		if err := tx.Model(&Goal{}).Where("status = ?", 1).First(&goal).Error; err != nil {
			return err
		}

		var plan Plan
		if err := tx.Model(&Plan{}).Where("goal_id = ? AND day = ?", goal.ID, day).First(&plan).Error; err != nil {
			return err
		}

		if err := tx.Model(&Task{}).Where("goal_id = ?", goal.ID).Find(&tasks).Error; err != nil {
			return err
		}

		planResponse = ToPlanResponse(plan, tasks)
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &planResponse, nil
}

func GetPlans() (*[]PlanResponse, error) {
	var planResponse []PlanResponse
	var tasks []Task

	err := db.GetDB().Debug().Transaction(func(tx *gorm.DB) error {
		var goal Goal
		if err := tx.Model(&Goal{}).Where("status = ?", 1).First(&goal).Error; err != nil {
			return err
		}

		var plans []Plan
		if err := tx.Model(&Plan{}).Where("goal_id = ?", goal.ID).Find(&plans).Error; err != nil {
			return err
		}

		if err := tx.Model(&Task{}).Where("goal_id = ?", goal.ID).Find(&tasks).Error; err != nil {
			return err
		}

		for _, plan := range plans {
			planResponse = append(planResponse, ToPlanResponse(plan, tasks))
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &planResponse, nil
}

func ToPlanResponse(plan Plan, tasks []Task) PlanResponse {
	taskResponse := ConvertTaskIDToTaskResponse(plan.TaskID, tasks)

	return PlanResponse{
		Day:        plan.Day,
		TasksToday: taskResponse,
	}
}
