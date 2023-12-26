package models_test

import (
	"regexp"
	"testing"

	"gin-okane-no-kyouiku/models"
	"gin-okane-no-kyouiku/testutils"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetGoal(t *testing.T) {
	mockDB, mock, err := testutils.NewMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `goals` ORDER BY created_at desc,`goals`.`id` LIMIT 1")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "point", "status"}).AddRow(1, "test", 10, 0),
		)

	_, err = models.GetGoal(mockDB)
	if err != nil {
		t.Errorf("error was not expected while getting goal: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestInsertGoalAndTasks(t *testing.T) {
	mockDB, mock, err := testutils.NewMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `goals` (`created_at`,`name`,`point`,`status`) VALUES (?,?,?,?)")).
		WithArgs(sqlmock.AnyArg(), "test", 10, 0).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `tasks` (`created_at`,`name`,`point`,`goal_id`) VALUES (?,?,?,?)")).
		WithArgs(sqlmock.AnyArg(), "test", 10, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	goal := models.Goal{Name: "test", Point: 10, Status: 0}
	tasks := []models.Task{{Name: "test", Point: 10}}
	err = models.InsertGoalAndTasks(mockDB, &goal, tasks)
	if err != nil {
		t.Errorf("error was not expected while inserting goal and tasks: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
