package queue

import "container/list"

type Memory struct {
	list.List
}

func (m *Memory) Put(value interface{}) {
	m.PushBack(value)
}

func (m *Memory) Get() (value interface{}, tx Txid) {
	return m.Front().Value, -1
}

func (m *Memory) Commit(tx Txid) {
}
