package tests

import (
	"go-experiments/dataStructures"
	"testing"
)

func TestAddingToHashTable(t *testing.T) {
	table := dataStructures.NewHashTable[int]()

	err := table.Add("One", 1)
	if err != nil {
		t.Fatalf("Unable to add : %v", err.Error())
	}
	err = table.Add("Two", 2)
	if err != nil {
		t.Fatalf("Unable to add : %v", err.Error())
	}
	err = table.Add("Three", 3)
	if err != nil {
		t.Fatalf("Unable to add : %v", err.Error())
	}

	if table.Size() != 3 {
		t.Fatalf("Size should be 3")
	}

	data, present, err := table.Get("One")
	if err != nil || !present || data != 1 {
		t.Fatalf("Unable to get One: %v / %v / %v", err.Error(), data, present)
	}

	data, present, err = table.Get("Two")
	if err != nil || !present || data != 2 {
		t.Fatalf("Unable to get Two : %v / %v / %v", err.Error(), data, present)
	}

	data, present, err = table.Get("Three")
	if err != nil || !present || data != 3 {
		t.Fatalf("Unable to get Three : %v / %v / %v", err.Error(), data, present)
	}

	data, present, err = table.Get("Five")
	if present {
		t.Fatalf("Should not find Five : %v / %v / %v", err.Error(), data, present)
	}
}
