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

// ErrAlreadyExists is an error that represents that the object with a given id already exists in a catalog.
type ErrAlreadyExists struct {
	TypeName string // the name of the type
	Id       string // the id of the object given
}

// Error returns a string version of the error.
func (e *ErrAlreadyExists) Error() string {
	return fmt.Sprintf("%q with id %q already exists", e.TypeName, e.Id)
}
