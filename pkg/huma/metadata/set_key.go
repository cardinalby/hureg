package metadata

import "github.com/danielgtaylor/huma/v2"

// SetKey sets a key-value pair in the operation metadata, initializing the metadata map if it is nil.
func SetKey(o *huma.Operation, key string, value interface{}) {
	if o.Metadata == nil {
		o.Metadata = make(map[string]any)
	}
	o.Metadata[key] = value
}
