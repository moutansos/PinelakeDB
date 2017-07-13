package storage_engine

import (
	"testing"
)

func TestAddField(t *testing.T) {
	var table CsvStorageEngineTable
	table.FieldNames = []string{"Test1", "Test2"}
	err := table.AddField("Test3")
	if err != nil {
		t.Error(err)
	} else if len(table.FieldNames) != 3 {
		t.Error("Invalid number of fields")
	} else if table.FieldNames[2] != "Test3" {
		t.Error("The new field is not equal to the expected value")
	}
}

func TestGetOrdinal(t *testing.T) {
	var table CsvStorageEngineTable
	table.FieldNames = []string{"Test1", "Test2", "Test3"}
	i, err := table.GetOrdinal("Test3")
	if err != nil {
		t.Error(err)
	} else if *i != 2 {
		t.Error("Invalid ordinal returned")
	}
}

func TestRemoveField(t *testing.T) {
	var table CsvStorageEngineTable
	table.FieldNames = []string{"Test1", "Test2", "Test3"}
	err := table.RemoveField("Test2")
	if err != nil {
		t.Error(err)
	} else if len(table.FieldNames) != 2 {
		t.Error("Invalid number of fields")
	} else if table.FieldNames[1] != "Test3" {
		t.Error("The fields are not in the expected order")
	}
}
