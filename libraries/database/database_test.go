package database

import "testing"

func TestMYSQLConnection(t *testing.T) {
	if _, err := Databases().MYSQLConnection(); err != nil {
		t.Error("MYSQL Connection FAILED. check the error bellow! \n")
		t.Error(err)
	} else {
		t.Log("MYSQL Connection ESTABLISHED!")
	}
}
