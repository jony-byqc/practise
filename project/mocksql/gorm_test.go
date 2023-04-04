package main

import (
	_ "database/sql"
	_ "fmt"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // need to import postgres dialect
)

type User struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"not null"`
	Age  int    `gorm:"not null"`
}

// TestDB tests the database connection
func TestDB(t *testing.T) {
	// create mock db connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// create gorm.DB instance with mocked database
	gdb, err := gorm.Open("postgres", db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer gdb.Close()

	// create User instance
	user := &User{
		Name: "John",
		Age:  30,
	}

	// using mocked database to insert User instance to the User table
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "users" \("name","age"\) VALUES \(\$1\,\$2\) RETURNING "users"\."id"`).
		WithArgs(user.Name, user.Age).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	if err = gdb.Create(user).Error; err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	if user.ID != 1 {
		t.Fatalf("unexpected user ID: %v", user.ID)
	}

	// ensure all expectations from mock are met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}
