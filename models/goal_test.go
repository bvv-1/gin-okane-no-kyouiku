package models

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetGoal(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mockDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true, // Ref: https://zenn.dev/tatane616/scraps/27d701e8c6658e
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer sqlDB.Close()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `goals` ORDER BY created_at desc,`goals`.`id` LIMIT 1")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "point", "status"}).AddRow(1, "test", 10, 0),
		)

	_, err = GetGoal(mockDB)
	if err != nil {
		t.Errorf("error was not expected while getting goal: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestInsertGoalAndTasks(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mockDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true, // Ref: https://zenn.dev/tatane616/scraps/27d701e8c6658e
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer sqlDB.Close()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `goals` (`created_at`,`name`,`point`,`status`) VALUES (?,?,?,?)")).
		WithArgs(sqlmock.AnyArg(), "test", 10, 0).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `tasks` (`created_at`,`name`,`point`,`goal_id`) VALUES (?,?,?,?)")).
		WithArgs(sqlmock.AnyArg(), "test", 10, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	goal := Goal{Name: "test", Point: 10, Status: 0}
	tasks := []Task{{Name: "test", Point: 10}}
	err = InsertGoalAndTasks(mockDB, &goal, tasks)
	if err != nil {
		t.Errorf("error was not expected while inserting goal and tasks: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
