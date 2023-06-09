package queue

import (
	"fmt"
)

const (
	MAXLEN = 1024 * 64
)

type SupportType interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | string
}

type Queue[T SupportType] struct {
	elems  []T
	length int
}

func (q *Queue[T]) Len() int { return q.length }

func (q *Queue[T]) Push(elems ...T) {
	l := len(elems)
	if l > MAXLEN-q.length {
		l = MAXLEN - q.length
	}
	q.elems = append(q.elems, elems[:l]...)
	q.length += l
}

func (q *Queue[T]) Pop() (T, error) {
	var front T
	if q.length == 0 {
		return front, fmt.Errorf("queue is empty")
	}
	front = q.elems[0]
	q.elems = q.elems[1:]
	q.length--
	return front, nil
}

func (q Queue[T]) Display() string {
	return fmt.Sprintln(q.elems)
}
