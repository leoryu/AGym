package q007

import (
	"container/list"
)

type Queue struct {
	appendStack *list.List
	deleteStack *list.List
}

func (q *Queue) AppendTail(value interface{}) {
	q.appendStack.PushBack(value)
}

func (q *Queue) DeleteHead() (value interface{}) {
	if q.deleteStack.Len() == 0 {
		for q.appendStack.Len() != 0 {
			q.deleteStack.PushBack(q.appendStack.Back())
			q.appendStack.Remove(q.appendStack.Back())
		}
	}
	if q.deleteStack.Len() != 0 {
		value = q.deleteStack.Back().Value
		q.deleteStack.Remove(q.deleteStack.Back())
	}
	return
}
