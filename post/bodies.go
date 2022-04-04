package post

import "sync"

type Bodies struct {
	Store map[int]string
	m *sync.RWMutex
}

func (b *Bodies) SetValue(id int, val string) {
	b.m.Lock()
	defer b.m.Unlock()
	b.Store[id] = val
}

func (b *Bodies) GetValue(id int) string {
	b.m.RLock()
	defer b.m.RUnlock()
	return b.Store[id]
}

func (b *Bodies) GetIds() []int {
	b.m.RLock()
	defer b.m.RUnlock()
	ids := make([]int, 0, len(b.Store))
	for id := range b.Store {
		ids = append(ids, id)
	}
	return ids
}
