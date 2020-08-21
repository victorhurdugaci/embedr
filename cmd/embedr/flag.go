// Copyright (c) Victor Hurdugaci (https://victorhurdugaci.com). All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

package main

type arrayFlags []string

func (arrFlags *arrayFlags) String() string {
	return ""
}

func (arrFlags *arrayFlags) Set(value string) error {
	*arrFlags = append(*arrFlags, value)
	return nil
}
