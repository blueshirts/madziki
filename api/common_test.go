package api

import (
	"testing"
)


func TestInit__CreateAndCloseSession(t *testing.T) {
	session := getSession()
	defer session.Close()
	if session == nil {
		t.Error("Database session is null")
	}
}
