// Copyright 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAcquireAndReleaseTanzuMetricDBLock(t *testing.T) {
	// Test acquiring the lock
	AcquireTanzuConfigLock()

	// Verify the lock is held
	assert.NotNil(t, tanzuConfigLock, "Expected the lock to be acquired")

	// Test releasing the lock
	assert.NotPanics(t, func() {
		ReleaseTanzuConfigLock()
	}, "Expected no panic while releasing the lock")

	// Verify the lock is released
	assert.Nil(t, tanzuConfigLock, "Expected the lock to be released")
}

func TestLockTimeout(t *testing.T) {
	// Acquire the lock for the first time
	AcquireTanzuConfigLock()

	testDefaultTimeout = 10 * time.Second
	// Try acquiring the lock again, should time out
	assert.Panics(t, func() {
		AcquireTanzuConfigLock()
	}, "Expected a panic while trying to acquire the lock again")

	// Release the initial lock
	ReleaseTanzuConfigLock()
}

func TestMultipleAcquireAndRelease(t *testing.T) {
	// Acquire and release the lock multiple times
	for i := 0; i < 3; i++ {
		assert.NotPanics(t, func() {
			AcquireTanzuConfigLock()
		}, "Expected no panics while acquiring the lock")

		assert.NotNil(t, tanzuConfigLock, "Expected the lock to be acquired")

		assert.NotPanics(t, func() {
			ReleaseTanzuConfigLock()
		}, "Expected no panic while releasing the lock")

		assert.Nil(t, tanzuConfigLock, "Expected the lock to be released")
	}
}

func TestParallelLockingAndUnlocking(t *testing.T) {
	const goroutines = 10
	var wg sync.WaitGroup
	wg.Add(goroutines)
	successCount := int32(0)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			AcquireTanzuConfigLock()
			// The first goroutine that successfully acquire the lock successfully
			// would sleep and hold lock for less than timeout period so that
			// all other go routines could acquire and release lock successfully
			time.Sleep(1000 * time.Millisecond)
			defer ReleaseTanzuConfigLock()
			successCount++
		}()
	}

	wg.Wait()
	assert.Equal(t, int32(10), successCount, "Expected all the goroutine to successfully acquire the lock")
}
