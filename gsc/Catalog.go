// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gsc

import (
	"reflect"
	"sync"
)

// Catalog is a concurrency-safe catalog of objects.
type Catalog struct {
	*sync.RWMutex
	objects map[string]interface{}
	indices map[string]map[string]int
}

// NewCatalog returns a new instance of a Catalog.
func NewCatalog() *Catalog {

	catalog := &Catalog{
		RWMutex: &sync.RWMutex{},
		objects: map[string]interface{}{},
		indices: map[string]map[string]int{},
	}

	return catalog
}

// Get retrieves an object from the catalog by id and type.
// If found, returns the object and true.
// If not found, returns nil and false.
func (c *Catalog) Get(id string, t reflect.Type) (interface{}, bool) {
	typeName := ""
	if t.Kind() == reflect.Ptr {
		typeName = t.Elem().Name()
	} else {
		typeName = t.Name()
	}
	if index, ok := c.indices[typeName]; ok {
		if position, ok := index[id]; ok {
			if objects, ok := c.objects[typeName]; ok {
				return reflect.ValueOf(objects).Index(position).Interface(), true
			}
		}
	}
	return nil, false
}

// Add adds the given object to the Catalog by id.
// If the object already exists in the catalog, returns an ErrAlreadyExists error.
func (c *Catalog) Add(id string, obj interface{}) error {
	t := reflect.TypeOf(obj)
	typeName := ""
	if t.Kind() == reflect.Ptr {
		typeName = t.Elem().Name()
	} else {
		typeName = t.Name()
	}

	if _, ok := c.indices[typeName]; !ok {
		c.indices[typeName] = map[string]int{}
	} else {
		if _, ok := c.indices[typeName][id]; ok {
			return &ErrAlreadyExists{TypeName: typeName, Id: id}
		}
	}

	if _, ok := c.objects[typeName]; !ok {
		c.objects[typeName] = reflect.MakeSlice(reflect.SliceOf(t), 0, 0).Interface()
	}

	c.objects[typeName] = reflect.Append(reflect.ValueOf(c.objects[typeName]), reflect.ValueOf(obj)).Interface()
	c.indices[typeName][id] = reflect.ValueOf(c.objects[typeName]).Len() - 1

	return nil
}

// Update updates the object in the catalog by id.
// If the object does not exist in the catalog, then returns an ErrMissingObject error.
func (c *Catalog) Update(id string, obj interface{}) error {
	objectType := reflect.TypeOf(obj)
	typeName := objectType.Elem().Name()
	if index, ok := c.indices[typeName]; ok {
		if position, ok := index[id]; ok {
			if objects, ok := c.objects[typeName]; ok {
				reflect.ValueOf(objects).Index(position).Set(reflect.ValueOf(obj))
				return nil
			}
		}
	}
	return &ErrMissingObject{TypeName: typeName, Id: id}
}

// Delete deletes the object from the catalog by id and type.
// If the object does not exist in the catalog, then returns an ErrMissingObject error.
func (c *Catalog) Delete(id string, t reflect.Type) error {
	typeName := t.Name()
	if index, ok := c.indices[typeName]; ok {
		if position, ok := index[id]; ok {
			if list, ok := c.objects[typeName]; ok {
				listValue := reflect.ValueOf(list)
				c.objects[typeName] = reflect.AppendSlice(listValue.Slice(0, position), listValue.Slice(position+1, listValue.Len())).Interface()
			}
			delete(index, id)
			return nil
		}
	}
	return &ErrMissingObject{TypeName: typeName, Id: id}
}

// List returns a list of objects in the catalog by type.
// If no objects with that type are found, then returns an empty slice.
func (c *Catalog) List(t reflect.Type) interface{} {
	if t.Kind() == reflect.Ptr {
		if list, ok := c.objects[t.Elem().Name()]; ok {
			return list
		}
	} else {
		if list, ok := c.objects[t.Name()]; ok {
			return list
		}
	}
	return reflect.MakeSlice(reflect.SliceOf(t), 0, 0).Interface()
}

// Objects returns all the objects in the catalog grouped by type name.
func (c *Catalog) Objects() map[string]interface{} {
	return c.objects
}
