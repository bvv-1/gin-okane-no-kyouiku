package models_test

import (
	"regexp"
	"testing"

	"gin-okane-no-kyouiku/models"
	"gin-okane-no-kyouiku/testutils"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestInsertProgress(t *testing.T) {
	mockDB, mock, err := testutils.NewMockDB()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	testTaskProgress := []models.TaskAndStatus{
		{Task: models.Task{Name: "test1", Point: 1}, IsDone: true},
		{Task: models.Task{Name: "test2", Point: 5}, IsDone: false},
	}

	mock.ExpectBegin()

	for id, taskStatus := range testTaskProgress {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `goals` WHERE status = ? ORDER BY `goals`.`id` LIMIT 1")).
			WithArgs(1).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "point", "status"}).AddRow(1, "test", 10, 1),
			)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `plans` WHERE goal_id = ? AND day = ?")).
			WithArgs(1, 1).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "day", "goal_id", "task_id"}).AddRow(1, 1, 1, 1),
			)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tasks` WHERE goal_id = ? AND name = ?")).
			WithArgs(1, taskStatus.Task.Name).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "point", "goal_id"}).AddRow(id+1, taskStatus.Task.Name, taskStatus.Task.Point, 1),
			)
	}
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `progresses` (`created_at`,`day`,`goal_id`,`plan_id`,`task_id`,`name`,`point`,`is_done`) VALUES (?,?,?,?,?,?,?,?),(?,?,?,?,?,?,?,?)")).
		WithArgs(sqlmock.AnyArg(), 1, 1, 1, 1, testTaskProgress[0].Task.Name, testTaskProgress[0].Task.Point, testTaskProgress[0].IsDone, sqlmock.AnyArg(), 1, 1, 1, 2, testTaskProgress[1].Task.Name, testTaskProgress[1].Task.Point, testTaskProgress[1].IsDone).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	if err = models.InsertProgress(mockDB, 1, testTaskProgress); err != nil {
		t.Errorf("error was not expected while inserting progress: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCheckProgress(t *testing.T) {
	// FIXME: write unit test
	// mockDB, mock, err := testutils.NewMockDB()
	// if err != nil {
	// 	t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	// }

	// mock.ExpectBegin()

	// mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `goals` WHERE status = ? ORDER BY `goals`.`id` LIMIT 1")).
	// 	WithArgs(1).
	// 	WillReturnRows(
	// 		sqlmock.NewRows([]string{"id", "name", "point", "status"}).AddRow(1, "test", 10, 1),
	// 	)
	// mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `progresses` WHERE (day, plan_id, task_id) IN (SELECT day, plan_id, task_id FROM progresses WHERE goal_id = ? GROUP BY plan_id, task_id HAVING MAX(created_at) = created_at) AND is_done = true")).
	// 	WithArgs(1).
	// 	WillReturnRows(
	// 		sqlmock.NewRows([]string{"id", "day", "goal_id", "plan_id", "task_id", "name", "point", "is_done"}).AddRow(1, 1, 1, 1, 1, "test1", 1, true),
	// 	)

	// mock.ExpectCommit()

	// if _, err = models.CheckProgress(mockDB); err != nil {
	// 	t.Errorf("error was not expected while checking progress: %s", err)
	// }

	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("there were unfulfilled expectations: %s", err)
	// }
}
