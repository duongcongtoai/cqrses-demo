package db

import (
	"testing"
)

func TestMysqlConnection(t *testing.T) {
	_, err := NewMysqlConnection()
	if err != nil {
		t.Error(err)
	}
}
