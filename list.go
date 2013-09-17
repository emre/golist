package golist

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type List struct {
	data   []interface{}
	locker sync.Mutex
}

var (
	EmptyListError  = "%v operation from empty list"
	OutOfRangeError = "%s index out of range"
)

// Returns a new List
func New(items ...interface{}) List {

	list := List{data: []interface{}{}}

	if len(items) > 0 {
		list.Append(items...)
	}

	return list
}

// Adds an item to the end of the list data
func (l *List) Append(items ...interface{}) {

	for _, value := range items {
		l.data = append(l.data, value)
	}

}

// Extend the list by appending all the items in te given list.
func (l *List) Extend(target_list List) {

	for _, value := range target_list.data {
		l.data = append(l.data, value)
	}

}

// Returns the length of the list
func (l *List) Len() int {
	return len(l.data)
}

// Insert an item at a given position.
// The first argument is the index of the element before which to insert
func (l *List) Insert(index int, value interface{}) {

	// Resize list to size(list) + 1 to get free space for new element.
	size := l.Len()

	l.Append(value)

	if size+1 >= index {

		for i := size - 1; i >= 0; i-- {

			if index == i {
				l.data[i+1] = l.data[i]
				l.data[index] = value
				break
			} else {
				l.data[i+1] = l.data[i]
			}

		}
	}
}

// Remove the first item from the list whose value is x.
// Returns an error if there is no such item exists.
func (l *List) Remove(value interface{}) error {

	error_text := fmt.Sprintf("'%v' is not in list", value)

	for index, data_value := range l.data {
		if data_value == value {
			l.data = append(l.data[:index], l.data[index+1:]...)
			return nil
		}
	}

	return errors.New(error_text)

}

// Remove the index at the given position in the list, and return it.
// If no index is specified, removes and returns the last item in the list.
func (l *List) Pop(index ...interface{}) (interface{}, error) {

	list_size := l.Len()

	if list_size == 0 {
		return nil, errors.New(fmt.Sprintf(EmptyListError, "Pop"))
	}

	var value interface{}
	var delete_index int

	if len(index) == 0 {

		value = l.data[list_size-1]

		delete_index = list_size - 1

	} else {

		_index := reflect.ValueOf(index[0]).Int()

		if int(_index) > list_size-1 {
			return nil, errors.New(fmt.Sprintf(OutOfRangeError, "Pop"))
		}

		value = l.data[_index]

		delete_index = int(_index)

	}

	error := l.Delete(delete_index)

	if error != nil {
		return nil, error
	}

	return value, nil

}

// Delete the item at the given position in the list.
func (l *List) Delete(index int) error {

	if l.Len() == 0 {
		return errors.New(fmt.Sprintf(EmptyListError, "Delete"))
	}

	l.data = append(l.data[:index], l.data[index+1:]...)

	return nil

}

// Returns the index in the list of the first item whose value is x. It is an error if there is no such item.
func (l *List) Index(value interface{}) (int, error) {

	list_size := l.Len()

	if list_size == 0 {
		return 0, errors.New(fmt.Sprintf(EmptyListError, "Index"))
	}

	for index, data_value := range l.data {
		if data_value == value {
			return index, nil
		}
	}

	error_text := fmt.Sprintf("'%v' is not in list", value)

	return 0, errors.New(error_text)

}

// Return the number of times x appears in the list.
func (l *List) Count(value interface{}) int {
	total_count := 0

	for _, data_value := range l.data {
		if data_value == value {
			total_count++
		}
	}

	return total_count

}

// Reverse the elements of the list, in place.
func (l *List) Reverse() {

	list_size := l.Len()

	if list_size > 0 {
		top_index := list_size - 1
		for index := 0; index < (top_index/2)+1; index++ {
			l.data[index], l.data[top_index-index] = l.data[top_index-index], l.data[index]
		}
	}
}
