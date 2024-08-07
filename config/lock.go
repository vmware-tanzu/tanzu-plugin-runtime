// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/alexflint/go-filemutex"
)

const (
	LocalTanzuFileLock = ".tanzu.lock"
	// DefaultLockTimeout is the default time waiting on the filelock
	DefaultLockTimeout = 10 * time.Minute
)

var tanzuConfigLockFile string

// testDefaultTimeout used for unit test only
var testDefaultTimeout time.Duration

// tanzuConfigLock used as a static lock variable that stores fslock
// This is used for interprocess locking of the config file
var tanzuConfigLock *filemutex.FileMutex

// mutex is used to handle the locking behavior between concurrent calls
// within the existing process trying to acquire the lock
var mutex sync.Mutex

// AcquireTanzuConfigLock tries to acquire lock to update tanzu config file with timeout
func AcquireTanzuConfigLock() {
	var err error

	if tanzuConfigLockFile == "" {
		path, err := ClientConfigPath()
		if err != nil {
			panic(fmt.Sprintf("cannot get config path while acquiring lock on tanzu config file, reason: %v", err))
		}
		tanzuConfigLockFile = filepath.Join(filepath.Dir(path), LocalTanzuFileLock)
	}

	// using fslock to handle interprocess locking
	var timeout time.Duration
	timeout = DefaultLockTimeout
	// To enable testing use testDefaultTimeout, if set use it else use the default
	if testDefaultTimeout.Seconds() != 0 {
		timeout = testDefaultTimeout
	}
	lock, err := getFileLockWithTimeOut(tanzuConfigLockFile, timeout)
	if err != nil {
		panic(fmt.Sprintf("cannot acquire lock for tanzu config file, reason: %v", err))
	}

	// Lock the mutex to prevent concurrent calls to acquire and configure the tanzuConfigLock
	mutex.Lock()
	tanzuConfigLock = lock

	// Get lock on config-ng.yaml
	AcquireTanzuConfigNextGenLock()
}

// ReleaseTanzuConfigLock releases the lock if the tanzuConfigLock was acquired
func ReleaseTanzuConfigLock() {
	if tanzuConfigLock == nil {
		return
	}
	if errUnlock := tanzuConfigLock.Close(); errUnlock != nil {
		panic(fmt.Sprintf("cannot release lock for tanzu config file, reason: %v", errUnlock))
	}

	tanzuConfigLock = nil
	// Unlock the mutex to allow other concurrent calls to acquire and configure the tanzuConfigLock
	mutex.Unlock()

	// Release lock on config-ng.yaml
	ReleaseTanzuConfigNextGenLock()
}

// getFileLockWithTimeOut returns a file lock with timeout
func getFileLockWithTimeOut(lockPath string, lockDuration time.Duration) (*filemutex.FileMutex, error) {
	dir := filepath.Dir(lockPath)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0o700); err != nil {
			return nil, err
		}
	}

	flock, err := filemutex.New(lockPath)
	if err != nil {
		return nil, err
	}

	result := make(chan error)
	cancel := make(chan struct{})
	go func() {
		err := flock.Lock()
		select {
		case <-cancel:
			// Timed out, cleanup if necessary.
			_ = flock.Close()
		case result <- err:
		}
	}()

	select {
	case err := <-result:
		return flock, err
	case <-time.After(lockDuration):
		close(cancel)
		return flock, fmt.Errorf("timeout waiting for lock")
	}
}
