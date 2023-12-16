package models

type Plan struct {
	Day        int    `json:"day"`
	TasksToday []Task `json:"tasks_today"`
}

type SuggestedPlan struct {
	Day        int    `json:"day"`
	PlansToday []Task `json:"plans_today"`
}

type DailyPlansResponse struct {
	Day        int    `json:"day"`
	PlansToday []Task `json:"plans_today"`
}
