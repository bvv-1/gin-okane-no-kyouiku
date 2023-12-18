package models

import (
	"gin-okane-no-kyouiku/db"

	"golang.org/x/xerrors"
)

type Goal struct {
	Name  string `json:"name"`
	Point int    `json:"point"`
	// Status    int       `json:"status"`
}

func GetGoal() (*Goal, error) {
	db := db.GetDB()

	var goal Goal
	if err := db.Debug().First(&goal).Error; err != nil {
		return nil, xerrors.Errorf("failed to get goal: %w", err)
	}
	return &goal, nil
}
