package conn

type Conn[T any] interface {
	ID() int
	C() chan T
	Done()
}

type conn[T any] struct {
	id   int
	c    chan T
	done chan struct{}
}

func (c conn[T]) ID() int {
	return c.id
}

func (c conn[T]) C() chan T {
	return c.c
}

func (c conn[T]) Done() {
	c.done <- struct{}{}
}

func (c conn[T]) close() {
	close(c.c)
	close(c.done)
}

func newConn[T any](id int) conn[T] {
	return conn[T]{
		id:   id,
		c:    make(chan T),
		done: make(chan struct{}, 1),
	}
}
