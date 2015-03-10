package list

import "sync"

type List struct {
	Head   *Item
	Tail   *Item
	Len    int
	Locker sync.RWMutex
}

type Item struct {
	Val  interface{}
	Next *Item
	Prev *Item
	List *List
}

func New() *List {
	list := &List{}
	list.Len = 0
	return list
}

func Insert(value interface{}, list *List) *List {
	newItem := &Item{value, list.Head, list.Tail, list}
	list.Locker.Lock()
	defer list.Locker.Unlock()

	if list.Head == nil {
		list.Head = newItem
		list.Tail = newItem
	} else {
		list.Head = newItem
		list.Head.Prev = newItem
		list.Tail.Next = newItem
	}

	list.Len++

	return list
}

func Has(value interface{}, list *List) bool {
	if list.Head == nil {
		return false
	}
	first := list.Head

	for {
		if first.Val == value {
			return true
		} else {
			if first.Next != nil {
				first = first.Next
			} else {
				return false
			}
		}
	}

	return false
}

func Remove(value interface{}, list *List) *List {
	list.Locker.RLock()

	if list.Head == nil {
		return list
	}

	list.Locker.RUnlock()

	list.Locker.RLock()
	first := list.Head
	last := list.Tail
	list.Locker.RUnlock()
	list.Locker.Lock()
	defer list.Locker.Unlock()

	for {
		if last.Next == nil {
			return list
		}

		if first.Val == value {
			first.Prev.Next = first.Next
			first.Next.Prev = first.Prev
			first.Prev = nil
			first.Next = nil
			first.Val = nil
			first.List = nil
			list.Len--
			return list
		} else {
			first = first.Next
		}
	}
}

func Length(list *List) int {
	return list.Len
}
