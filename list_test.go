package golist

import (
	//"errors"
	//"reflect"
	"testing"
)

func TestListInit(t *testing.T) {
	my_list := New()

	list_size := my_list.Len()
	if list_size != 0 {
		t.Error("new: wrong length")
	}

	another_list := New("string", 99, 10.45)

	list_size = another_list.Len()
	if list_size != 3 {
		t.Error("new(2): wrong length")
	}

}

func TestListAppend(t *testing.T) {
	my_list := New(1, 2, 3, 4, 5)

	my_list.Append(6)

	if my_list.data[5] != 6 {
		t.Error("wrong value after append operation.")
	}

}

func TestListExtend(t *testing.T) {
	my_list := New("Turkiye", "Guney kore")
	other_list := New("Ukrayna", "Tayland")

	my_list.Extend(other_list)

	if my_list.Len() != 4 {
		t.Error("extend: wrong size for the list")
	}

	if my_list.data[2] != "Ukrayna" || my_list.data[3] != "Tayland" {
		t.Error("extend: wrong values for the extended list")
	}
}

func TestListIndex(t *testing.T) {
	my_list := New("That's", "what", "she", "said")

	index, error := my_list.Index("she")
	if error != nil {
		t.Error(error)
	}

	if index != 2 {
		t.Error("index: wrong index")
	}

	other_list := New("thor", "is", "here")

	_, _error := other_list.Index("nothere")

	if _error == nil {
		t.Error("index: there should be an error")
	}

}

func TestListPop(t *testing.T) {
	my_list := New(1, 2, 3, 4, 5)

	value, error := my_list.Pop()

	if error != nil {
		t.Error(error)
	}

	if value != 5 {
		t.Error("pop: wrong value")
	}

	my_list = New()

	_, error = my_list.Pop()

	if error == nil {
		t.Error("pop: there should be an error if list size is zero.")
	}

}

func TestListPopItem(t *testing.T) {
	my_list := New(1, 2, 3, 4, 5)

	value, error := my_list.Pop(0)

	if error != nil {
		t.Error(error)
	}

	if value != 1 {
		t.Error("pop_item: wrong value")
	}

	my_list = New()

	_, error = my_list.Pop(0)

	if error == nil {
		t.Error("pop_item: there should be an error if list size is zero.")
	}
}

func TestListCount(t *testing.T) {
	my_list := New(1, 10, 10, 2, 10)

	total_ten := my_list.Count(10)

	if total_ten != 3 {
		t.Error("count: item 10's count should be three.")
	}

	total_three := my_list.Count(3)

	if total_three != 0 {
		t.Error("count: item 3's count should be zero.")
	}

}

func TestRemoveOperations(t *testing.T) {
	my_list := New(1, 10, 20, 30)
	my_list.Delete(0)

	if len(my_list.data) != 3 {
		t.Error("delete: list size should be three after delete.")
	}

	for _, value := range my_list.data {
		if value == 1 {
			t.Error("remove: item 1 should be removed already.")
		}
	}

	my_list = New(5, 10, 15)
	my_list.Remove(5)

	if len(my_list.data) != 2 {
		t.Error("delete: list size should be three after delete.")
	}

	for _, value := range my_list.data {
		if value == 5 {
			t.Error("remove: item 5 should be removed already.")
		}
	}
}

func TestReverse(t *testing.T) {
	my_list := New(1, 2)
	my_list.Reverse()

	if my_list.data[0] != 2 {
		t.Error("reverse: invalid reverse operation")
	}
}
