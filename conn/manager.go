package conn

import (
	"context"
	"sync"
)

type Manager[T any] interface {
	CreateConn(ctx context.Context, wg *sync.WaitGroup) Conn[T]
	GetConnection(id int) Conn[T]
}

func NewManager[T any](cap int) Manager[T] {
	return &manager[T]{
		id:   0,
		cap:  cap,
		pool: make(map[int]conn[T]),
	}
}

type manager[T any] struct {
	mu   sync.Mutex
	id   int
	cap  int
	pool map[int]conn[T]
}

func (m *manager[T]) CreateConn(ctx context.Context, wg *sync.WaitGroup) Conn[T] {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.pool) == m.cap {
		return nil
	}

	c := newConn[T](m.nextID())
	m.pool[c.ID()] = c

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-c.done:
				c.close()
				delete(m.pool, c.id)
				return
			case <-ctx.Done():
				return
			}
		}
	}()

	return c
}

func (m *manager[T]) GetConnection(id int) Conn[T] {
	return m.pool[id]
}

func (m *manager[T]) nextID() int {
	m.id++
	return m.id - 1
}
