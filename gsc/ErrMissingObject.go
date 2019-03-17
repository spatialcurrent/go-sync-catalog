// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gsc

import (
	"fmt"
)

// ErrMissingObject is an error that represents that the object with a given id does not exist in a catalog.
type ErrMissingObject struct {
	TypeName string // the name of the type
	Id       string // the id of the object given
}

// Error returns a string version of the error.
func (e *ErrMissingObject) Error() string {
	return fmt.Sprintf("%q with id %q does not exist or otherwise missing", e.TypeName, e.Id)
}
