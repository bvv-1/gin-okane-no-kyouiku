package testutils

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	mockDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true, // Ref: https://zenn.dev/tatane616/scraps/27d701e8c6658e
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	return mockDB, mock, nil
}
