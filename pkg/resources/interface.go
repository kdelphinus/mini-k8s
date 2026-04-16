package resources

type Resource interface {
	GetName() string
	GetKind() string
}
