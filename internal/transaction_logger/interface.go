package transaction_logger

type TransactionLogger interface {
	ReadEvents() (<-chan Event, <-chan error)
	WritePut(key, value string)
	WriteDelete(key string)
	Err() <-chan error
	Run()
}
