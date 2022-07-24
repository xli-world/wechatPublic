package responsibilitychain

type department interface {
	execute(*patient)
	setNext(department)
}
