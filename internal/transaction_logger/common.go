package transaction_logger

import (
	"fmt"
	"goStorage/internal/storage"
)

const (
	_                     = iota
	EventDelete EventType = iota
	EventPut
)

type EventType byte

type Event struct {
	Sequence  uint64
	EventType EventType
	Key       string
	Value     string
}

var Logger TransactionLogger

func InitializeTransactionLog() error {
	var err error

	Logger, err = NewFileTransactionLogger("transaction.log")
	if err != nil {
		return fmt.Errorf("failed to create event logger: %w", err)
	}

	events, errors := Logger.ReadEvents()
	e, ok := Event{}, true

	for ok && err == nil {
		select {
		case err, ok = <-errors:
		case e, ok = <-events:
			switch e.EventType {
			case EventDelete:
				err = storage.Delete(e.Key)
			case EventPut:
				err = storage.Put(e.Key, e.Value)
			}
		}
	}

	Logger.Run()

	return err
}
