package workers

// Worker interface defines the contract for all worker implementations.
type Worker interface {
	DoWork() string
}
