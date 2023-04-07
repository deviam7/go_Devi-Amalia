package main

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestConnectDB(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer dbMock.Close()

	mock.ExpectPing()

	db, err := ConnectDB(dbMock)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, db)

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCloseDB(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer dbMock.Close()

	mock.ExpectClose()

	db, err := ConnectDB(dbMock)
	if err != nil {
		t.Fatal(err)
	}

	err = CloseDB(db)
	if err != nil {
		t.Fatal(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
