// Copyright (c) Victor Hurdugaci (https://victorhurdugaci.com). All rights reserved.
// Licensed under the MIT License. See LICENSE in the project root for license information.

package main

import "sort"

type sortedStringSet map[string]struct{}

func (s sortedStringSet) Add(values ...string) {
	for _, value := range values {
		s[value] = struct{}{}
	}
}

func (s sortedStringSet) All() []string {
	all := make([]string, len(s))
	i := 0
	for k := range s {
		all[i] = k
		i++
	}

	sort.Strings(all)
	return all
}
