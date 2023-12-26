package models_test

import (
	"math/rand"
	"reflect"
	"regexp"
	"testing"
	"time"

	"gin-okane-no-kyouiku/models"
	"gin-okane-no-kyouiku/testutils"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetSuggestedPlans(t *testing.T) {
	mockDB, mock, err := testutils.NewMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `goals` ORDER BY created_at desc,`goals`.`id` LIMIT 1")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "point", "status"}).AddRow(1, "test", 10, 0),
		)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tasks` WHERE goal_id = ?")).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "point", "goal_id"}).AddRow(1, "test", 10, 1),
		)
	// NOTE: タスクが1つだけなので結果が一意に定まる
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `plans` (`created_at`,`day`,`task_id`,`goal_id`) VALUES (?,?,?,?),(?,?,?,?)")).
		WithArgs(sqlmock.AnyArg(), 1, 1, 1, sqlmock.AnyArg(), 2, 1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	_, err = models.GetSuggestedPlans(mockDB)
	if err != nil {
		t.Errorf("error was not expected while getting suggested plans: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGeneratePlans(t *testing.T) {
	// シードを設定
	rand.Seed(time.Now().UnixNano())

	testGoal := models.Goal{Name: "test", Point: 10, Status: 0}
	testTasks := []models.Task{
		{Name: "test1", Point: 10, GoalID: 1},
		{Name: "test2", Point: 20, GoalID: 1},
		{Name: "test3", Point: 30, GoalID: 1},
	}
	testDays := 5

	result := models.GeneratePlans(testGoal, testTasks, testDays)

	expectedResult := make([]models.Plan, testDays)
	for day := 1; day <= testDays; day++ {
		taskID := testTasks[rand.Intn(len(testTasks))].ID
		expectedResult[day-1] = models.Plan{Day: day, TaskID: taskID, GoalID: testGoal.ID}
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("GeneratePlans() returned %+v, expected %+v", result, expectedResult)
	}
}

func TestAcceptSuggestedPlans(t *testing.T) {
	mockDB, mock, err := testutils.NewMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `plans` ORDER BY created_at desc,`plans`.`id` LIMIT 1")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "day", "task_id", "goal_id"}).AddRow(1, 1, 1, 1),
		)
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `goals` SET `status`=? WHERE id = ?")).
		WithArgs(1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = models.AcceptSuggestedPlans(mockDB)
	if err != nil {
		t.Errorf("error was not expected while accepting suggested plans: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetPlanByDay(t *testing.T) {
	mockDB, mock, err := testutils.NewMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `goals` WHERE status = ? ORDER BY `goals`.`id` LIMIT 1")).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "point", "status"}).AddRow(1, "test", 10, 1),
		)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `plans` WHERE goal_id = ? AND day = ?")).
		WithArgs(1, 1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "day", "task_id", "goal_id"}).AddRow(1, 1, 1, 1),
		)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tasks` WHERE goal_id = ?")).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "point", "goal_id"}).AddRow(1, "test", 10, 1),
		)
	mock.ExpectCommit()

	_, err = models.GetPlanByDay(mockDB, 1)
	if err != nil {
		t.Errorf("error was not expected while getting plan by day: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetPlans(t *testing.T) {
	mockDB, mock, err := testutils.NewMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `goals` WHERE status = ? ORDER BY `goals`.`id` LIMIT 1")).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "point", "status"}).AddRow(1, "test", 10, 1),
		)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `plans` WHERE goal_id = ?")).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "day", "task_id", "goal_id"}).AddRow(1, 1, 1, 1),
		)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tasks` WHERE goal_id = ?")).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "point", "goal_id"}).AddRow(1, "test", 10, 1),
		)
	mock.ExpectCommit()

	_, err = models.GetPlans(mockDB)
	if err != nil {
		t.Errorf("error was not expected while getting plans: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
