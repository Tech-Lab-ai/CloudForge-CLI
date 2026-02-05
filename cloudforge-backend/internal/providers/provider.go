package providers

// Provider is the interface that all cloud providers must implement.
type Provider interface {
	Apply()
	Destroy()
	Plan() (string, error)
	Deploy(path string)
}
