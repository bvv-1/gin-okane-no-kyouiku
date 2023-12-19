package models

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
