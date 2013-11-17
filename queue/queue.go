package queue

type Txid int

type Queue interface {
	Get() *Task
	Put(interface{})
	Commit(Txid) error
}

type Task struct {
	Item  interface{}
	tx    Txid
	queue Queue
}

func (t *Task) Commit(err error) {
	err = t.queue.Commit(t.tx)
	return
}
