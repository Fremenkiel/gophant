package interfaces

type ErrorReporter interface {
	Report(err error)
}

