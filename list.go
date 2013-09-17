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
	l.locker.Lock()
	l.data = append(l.data, items...)
	l.locker.Unlock()

}

// Extend the list by appending all the items in te given list.
func (l *List) Extend(target_list List) {
	l.locker.Lock()
	for _, value := range target_list.data {
		l.data = append(l.data, value)
	}
	l.locker.Unlock()

}

// Returns the length of the list
func (l *List) Len() int {
	return len(l.data)
}

// Insert an item at a given position.
// The first argument is the index of the element before which to insert
func (l *List) Insert(index int, value interface{}) {

	// Resize list to size(list) + 1 to get free space for new element.
	l.locker.Lock()
	size := l.Len()
	l.locker.Unlock()
	l.Append(value)

	if size+1 >= index {

		for i := size - 1; i >= 0; i-- {

			if index == i {
				l.locker.Lock()
				l.data[i+1] = l.data[i]
				l.data[index] = value
				l.locker.Unlock()
				break
			} else {
				l.locker.Lock()
				l.data[i+1] = l.data[i]
				l.locker.Unlock()
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
			l.locker.Lock()
			l.data = append(l.data[:index], l.data[index+1:]...)
			l.locker.Unlock()
			return nil
		}
	}

	return errors.New(error_text)

}

// Remove the index at the given position in the list, and return it.
// If no index is specified, removes and returns the last item in the list.
func (l *List) Pop(index ...interface{}) (interface{}, error) {

	l.locker.Lock()
	list_size := l.Len()
	l.locker.Unlock()

	if list_size == 0 {
		return nil, errors.New(fmt.Sprintf(EmptyListError, "Pop"))
	}

	var value interface{}
	var delete_index int

	if len(index) == 0 {

		l.locker.Lock()
		value = l.data[list_size-1]

		delete_index = list_size - 1
		l.locker.Unlock()

	} else {
		l.locker.Lock()
		_index := reflect.ValueOf(index[0]).Int()

		if int(_index) > list_size-1 {
			return nil, errors.New(fmt.Sprintf(OutOfRangeError, "Pop"))
		}

		value = l.data[_index]

		delete_index = int(_index)
		l.locker.Unlock()

	}

	error := l.Delete(delete_index)

	if error != nil {
		return nil, error
	}

	return value, nil

}

// Delete the item at the given position in the list.
func (l *List) Delete(index int) error {

	list_size := l.Len()

	if list_size == 0 {
		return errors.New(fmt.Sprintf(EmptyListError, "Delete"))
	}

	if index > list_size-1 {
		return errors.New(fmt.Sprintf(OutOfRangeError, "Delete"))
	}

	l.locker.Lock()
	l.data = append(l.data[:index], l.data[index+1:]...)
	l.locker.Unlock()
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

	l.locker.Lock()

	for _, data_value := range l.data {
		if data_value == value {
			total_count++
		}
	}

	l.locker.Unlock()

	return total_count

}

// Reverse the elements of the list, in place.
func (l *List) Reverse() {

	l.locker.Lock()
	list_size := l.Len()
	l.locker.Unlock()

	if list_size > 0 {
		top_index := list_size - 1
		l.locker.Lock()
		for index := 0; index < (top_index/2)+1; index++ {
			l.locker.Lock()
			l.data[index], l.data[top_index-index] = l.data[top_index-index], l.data[index]
			l.locker.Unlock()
		}
		l.locker.Unlock()

	}

}
