package db

type CounterPersistence interface {
	GetCurrentCount() int
	UpdateCurrentCount() int
}
