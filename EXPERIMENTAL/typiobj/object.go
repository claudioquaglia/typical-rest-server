package typiobj

// Name of object. Return value from name field if available or return its type.
func Name(obj interface{}) string {
	// TODO:
	return "<name>"
}

// Description of Object. Return value from description field if available or return its type
func Description(obj interface{}) string {
	// TODO:
	return "<description>"
}

// IsCommandLiner return true if object implementation of CommandLiner
func IsCommandLiner(obj interface{}) (ok bool) {
	_, ok = obj.(CommandLiner)
	return
}

// IsConfigurer return true if object implementation of configurer
func IsConfigurer(obj interface{}) (ok bool) {
	_, ok = obj.(Configurer)
	return
}
