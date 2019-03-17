// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gsc

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type test struct {
	Id string
}

func TestCatalog(t *testing.T) {

	c := NewCatalog()

	foo := "foo"

	err := c.Add("foo", foo)
	assert.Nil(t, err)

	bar, ok := c.Get("foo", reflect.TypeOf(foo))
	assert.True(t, ok)
	assert.NotNil(t, bar)
	assert.Equal(t, foo, bar)

	x := test{Id: "foo"}
	err = c.Add(x.Id, x)
	assert.Nil(t, err)

	y, ok := c.Get("foo", reflect.TypeOf(test{}))
	assert.True(t, ok)
	assert.NotNil(t, y)
	assert.Equal(t, x, y)

	y, ok = c.Get("foo", reflect.TypeOf(&test{}))
	assert.True(t, ok)
	assert.NotNil(t, y)
	assert.Equal(t, x, y)

}
