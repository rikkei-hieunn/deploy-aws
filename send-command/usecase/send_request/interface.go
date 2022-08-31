package sendrequest

// ISender method send operator command
type ISender interface {
	HandleRequest() error
}
