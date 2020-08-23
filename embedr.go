// Copyright (c) Victor Hurdugaci (https://victorhurdugaci.com). All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

package embedr

import (
	"bytes"
	"encoding/base64"
	"io"
	"os"
)

// Embedr represents a collection of embeded files
type Embedr struct {
	files map[string][]byte
}

// New creates a new instance of Embedr
func New() *Embedr {
	return &Embedr{}
}

// Add adds a new file embeded file
func (e *Embedr) Add(path string, content []byte) error {
	if e.files == nil {
		e.files = make(map[string][]byte)
	}

	e.files[path] = content
	return nil
}

// Open is used to read the contents of an embeded file
func (e *Embedr) Open(path string) (io.Reader, error) {
	content, exists := e.files[path]

	if !exists {
		return nil, os.ErrNotExist
	}

	return base64.NewDecoder(base64.StdEncoding, bytes.NewReader(content)), nil
}

// WalkFunc is the argument passed to Walk
type WalkFunc func(path string) error

// Walk traverses all embeded files.
// If the argument function returns an error, that error is returned by Walk
func (e *Embedr) Walk(walkFn WalkFunc) error {
	for filePath := range e.files {
		err := walkFn(filePath)
		if err != nil {
			return err
		}
	}

	return nil
}
