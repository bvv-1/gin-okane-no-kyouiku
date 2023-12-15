package models

type SuggestedPlan struct {
	Day        int    `json:"day"`
	PlansToday []Task `json:"plans_today"`
}
