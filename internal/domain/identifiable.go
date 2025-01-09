package domain

// Identifiable is a custom interface to ensure that entities have an ID
type Identifiable interface {
	GetID() interface{} // This method will return the entity's ID
}
