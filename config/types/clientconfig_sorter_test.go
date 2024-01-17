// Copyright 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package types specifies the types for clientconfig
package types

import (
	"sort"
	"testing"
)

func TestContextSorter_Sort(t *testing.T) {
	// Create a list of Context instances to be sorted
	contexts := []*Context{
		{Name: "Context1", Target: "tmc"},
		{Name: "Context2", Target: "k8s"},
		{Name: "Context1", Target: "k8s"},
	}

	// Sort the contexts using the ContextSorter
	sort.Sort(ContextSorter(contexts))

	// Verify the sorted order, which should be by name and then by target
	expectedOrder := []*Context{
		{Name: "Context1", Target: "k8s"},
		{Name: "Context1", Target: "tmc"},
		{Name: "Context2", Target: "k8s"},
	}

	for i, context := range contexts {
		if context.Name != expectedOrder[i].Name || context.Target != expectedOrder[i].Target {
			t.Errorf("Expected %v, but got %v", expectedOrder[i], context)
		}
	}
}
