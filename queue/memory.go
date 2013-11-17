package queue

import "container/list"

type Memory struct {
	list.List
}

func (m *Memory) Put(item interface{}) {
	m.PushBack(item)
}

func (m *Memory) Get() *Task {
	var tx Txid
	item := m.Front()
	if item == nil {
		return nil
	}
	return &Task{
		Item:  m.Remove(item),
		tx:    tx,
		queue: m,
	}
}

func (m *Memory) Commit(tx Txid) error {
	return nil
}
