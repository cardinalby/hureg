package op_handler

import (
	"github.com/danielgtaylor/huma/v2"
)

// SetHidden sets the Hidden field of the operation to the given value.
// If `override` is false and the operation already has Hidden set to true, it will not be changed.
func SetHidden(isHidden bool, override bool) func(o *huma.Operation) {
	return func(o *huma.Operation) {
		if override || !o.Hidden {
			o.Hidden = isHidden
		}
	}
}
