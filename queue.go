package trex

type queue[T any] struct {
	items []T
}

func (q *queue[T]) Schedule(item T) {
	q.items = append(q.items, item)
}

func (q *queue[T]) Next() (item T) {
	if q.isEmpty() {
		return
	}
	return q.pop()
}

func (q *queue[T]) isEmpty() bool {
	return len(q.items) == 0
}

func (q *queue[T]) pop() T {
	item := q.head()
	q.drop()
	return item
}

func (q *queue[T]) head() T {
	return q.items[0]
}

func (q *queue[T]) drop() {
	q.items = q.items[1:]
}
