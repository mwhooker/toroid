package queue

type Txid int

type Queue interface {
	Get() (interface{}, Txid)
	Put(interface{})
	Commit(Txid)
}
