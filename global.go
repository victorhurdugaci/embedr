// Copyright (c) Victor Hurdugaci (https://victorhurdugaci.com). All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

package embedr

import "io"

var global = New()

// GlobalEmbedr returns the global (default) version of embedr
func GlobalEmbedr() *Embedr {
	return global
}

// Add adds a new file embeded file
func Add(path string, content string) {
	global.Add(path, []byte(content))
}

// Open is used to read the contents of an embeded file
func Open(path string) (io.Reader, error) {
	return global.Open(path)
}

// Walk traverses all embeded files.
// If the argument function returns an error, that error is returned by Walk
func Walk(walkFn WalkFunc) error {
	return global.Walk(walkFn)
}
