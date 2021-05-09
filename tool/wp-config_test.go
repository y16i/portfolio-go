package tool

import (
	"testing"
)

func TestFindKeys(t *testing.T) {
	var result string = ""
	result = getValue("DB_NAME", "DB_NAME', 'databaseName'")
	if result != "databaseName" {
		t.Error(result)
	}
	result = getValue("DB_USER", "DB_USER', 'userName'")
	if result != "userName" {
		t.Error(result)
	}
	result = getValue("DB_PASSWORD", "DB_PASSWORD', 'strongPassword'")
	if result != "strongPassword" {
		t.Error(result)
	}
	result = getValue("DB_HOST", "DB_HOST', 'localhost'")
	if result != "localhost" {
		t.Error(result)
	}
}
