package message

type Headers map[string]interface{}

// Get returns the metadata value for the given key. If the key is not found, an empty string is returned.
func (h Headers) Get(key string) interface{} {
	if v, ok := h[key]; ok {
		return v
	}

	return nil
}

// Set sets the metadata key to value.
func (h Headers) Set(key string, value interface{}) {
	h[key] = value
}
