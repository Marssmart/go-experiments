package tests

import "testing"
import "go-experiments/dataStructures"

func TestCreation(t *testing.T) {
	var list = dataStructures.NewLinkedList[int]()
	if list.Size() != 0 {
		t.Fatalf("Should have been 0")
	}
}

func TestAdding(t *testing.T) {
	var list = dataStructures.NewLinkedList[int]()
	if list.Size() != 0 {
		t.Fatalf("Should have been 0")
	}

	list.Add(7)
	if list.Size() != 1 || !list.Contains(7) {
		t.Fatalf("Not added correctly")
	}

	list.Add(5)
	if list.Size() != 2 || !list.Contains(5) {
		t.Fatalf("Not added correctly")
	}

	//add duplicate
	list.Add(7)
	if list.Size() != 3 || !list.Contains(7) {
		t.Fatalf("Not added correctly")
	}
}

func TestRemovingAll(t *testing.T) {
	var list = dataStructures.NewLinkedList[int]()
	list.Add(7)
	list.Add(5)
	list.Add(54)
	if list.Size() != 3 || !list.ContainsAll(5, 7, 54) {
		t.Fatalf("Not added properly")
	}

	list.RemoveAll(5)
	if list.Size() != 2 && !list.ContainsAll(7, 54) || list.Contains(5) {
		t.Fatalf("Not removed properly")
	}
}

func TestRemovingFirst(t *testing.T) {
	var list = dataStructures.NewLinkedList[int]()
	list.Add(7)
	list.Add(5)
	list.Add(54)
	list.Add(5)
	list.Add(54)
	list.Add(54)

	if list.Size() != 6 || !list.ContainsAll(5, 7, 54) {
		t.Fatalf("Not added properly")
	}

	list.RemoveFirst(5)
	list.RemoveFirst(54)
	list.RemoveFirst(54)
	if list.Size() != 3 || !list.ContainsAll(5, 7, 54) {
		t.Fatalf("Not removed properly")
	}
}

func TestLargeCollectionAdd(t *testing.T) {
	list := CreateBackLoadedListOfSize(100_000)

	if list.Size() != 100_000 {
		t.Fatalf("Should have been %v", 100_000)
	}
}

func CreateBackLoadedListOfSize(size int) dataStructures.LinkedList[int] {
	var list = dataStructures.NewLinkedList[int]()
	nums := make([]int, size)
	for i := range nums {
		list.Add(i)
	}
	return list
}

func TestLargeCollectionAddFirst(t *testing.T) {
	list := createFrontLoadedListOfSize(100_000)

	if list.Size() != 100_000 {
		t.Fatalf("Should have been %v", 100_000)
	}
}

func TestLargeCollectionContains(t *testing.T) {
	var list = createFrontLoadedListOfSize(100_000)

	if !list.Contains(99_999) {
		t.Fatalf("Should have contained %v", 99_999)
	}
}

func TestLargeCollectionContainsAll(t *testing.T) {
	var list = createFrontLoadedListOfSize(100_000)

	if !list.ContainsAll(9_999, 99_999, 99_999) {
		t.Fatalf("Should have contained %v", 99_999)
	}
}

func TestLargeCollectionGet(t *testing.T) {
	list := CreateBackLoadedListOfSize(100_000)

	safeGet(t, list, 1_000)
	safeGet(t, list, 10_000)
	safeGet(t, list, 54_322)
}

func safeGet(t *testing.T, list dataStructures.LinkedList[int], index int) {
	value, _, err := list.Get(index)
	if err != nil || value != index {
		t.Fatalf("Should have had %v", index)
	}
}

func createFrontLoadedListOfSize(size int) dataStructures.LinkedList[int] {
	var list = dataStructures.NewLinkedList[int]()
	nums := make([]int, size)
	for i := range nums {
		list.AddFirst(i)
	}
	return list
}
