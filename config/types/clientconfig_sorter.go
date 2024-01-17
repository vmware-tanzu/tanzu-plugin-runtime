// Copyright 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package types specifies the types for clientconfig
package types

// ContextSorter is a type that implements the sort interface for Contexts to sort by name and target.
type ContextSorter []*Context

// Len returns the length of the ContextSorter.
func (c ContextSorter) Len() int {
	return len(c)
}

// Swap swaps the elements at the given indices.
func (c ContextSorter) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// Less compares the Contexts (by name and then by target) at the given indices.
func (c ContextSorter) Less(i, j int) bool {
	if c[i].Name == c[j].Name {
		return c[i].Target < c[j].Target
	}
	return c[i].Name < c[j].Name
}
