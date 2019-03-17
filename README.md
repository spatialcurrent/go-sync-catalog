[![CircleCI](https://circleci.com/gh/spatialcurrent/go-sync-catalog/tree/master.svg?style=svg)](https://circleci.com/gh/spatialcurrent/go-sync-catalog/tree/master) [![Go Report Card](https://goreportcard.com/badge/spatialcurrent/go-sync-catalog)](https://goreportcard.com/report/spatialcurrent/go-sync-catalog)  [![GoDoc](https://godoc.org/github.com/spatialcurrent/go-sync-catalog?status.svg)](https://godoc.org/github.com/spatialcurrent/go-sync-catalog) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/spatialcurrent/go-sync-catalog/blob/master/LICENSE.md)

# go-sync-catalog

# Description

**go-sync-catalog** (aka GSC) is a simple library providing a concurrency-safe catalog of objects.  GSC uses a `*sync.RWMutex` to block writes during read operations.

# Usage

**Go**

You can import **go-sync-catalog** as a library with:

```
import (
  "github.com/spatialcurrent/go-sync-catalog/gsc"
)
```

See [gsc](https://godoc.org/github.com/spatialcurrent/go-sync-catalog/gsc) in GoDoc for information on how to use Go API.

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-sync-catalog/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
