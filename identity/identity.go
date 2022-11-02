package identity

type Identity interface {
	// GetID returns a globally unique string identifier.
	GetID() string

	// GetDisplayName returns a name that humans would like to read
	GetDisplayName() string
}
