// Copyright 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"strconv"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// GetDatabaseLastUpdateTimestamp retrieves the last update timestamp of the database image and refreshes the cache.
// It fetches the configuration details and returns the last update timestamp.
// Returns the timestamp of the last update timestamp and an error, if any.
func GetDatabaseLastUpdateTimestamp() (time.Time, error) {
	// Retrieve client config
	cfg, err := GetClientConfig()
	if err != nil {
		return time.Time{}, err
	}
	return cfg.GetDatabaseLastUpdateTimestamp()
}

// SetDatabaseLastUpdateTimestamp sets the timestamp of the last database update.
// It acquires a lock to ensure thread safety, updates the timestamp in the configuration,
// and then releases the lock. The function will persist the updated configuration if needed.
// Takes a time.Time value to set as the new timestamp.
// Returns an error if the operation fails.
func SetDatabaseLastUpdateTimestamp(val time.Time) error {
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}

	// Add or Update last update timestamp in the yaml node
	persist := setDatabaseLastUpdateTimestamp(node, KeyDatabaseLastUpdateTimestamp, val)

	// Persist the config node to the file
	if persist {
		return persistConfig(node)
	}
	return err
}

// setDatabaseLastUpdateTimestamp updates the specified yaml node with the new timestamp.
// This is a helper function to modify the configuration node with the provided timestamp.
// It formats the timestamp to RFC3339 format before updating.
// Returns a boolean indicating whether the configuration should be persisted.
func setDatabaseLastUpdateTimestamp(node *yaml.Node, key string, val time.Time) (persist bool) {
	databaseLastUpdateTimestampNode := getNGCLIOptionsChildNode(key, node)
	if databaseLastUpdateTimestampNode != nil && databaseLastUpdateTimestampNode.Value != val.Format(time.RFC3339) {
		databaseLastUpdateTimestampNode.Value = val.Format(time.RFC3339)
		persist = true
	}
	return persist
}

// DatabaseRefreshOptions struct holds options for database update operations.
type DatabaseRefreshOptions struct {
	Duration time.Duration // Duration specifies the time interval in hours.
}

// DatabaseRefreshOpts defines a function type for modifying DatabaseRefreshOptions.
type DatabaseRefreshOpts func(o *DatabaseRefreshOptions)

// WithDuration is an option setter function to set the duration in DatabaseRefreshOptions.
// Takes an integer representing the duration in hours.
// Returns a function of type DatabaseRefreshOpts.
func WithDuration(duration int) DatabaseRefreshOpts {
	return func(d *DatabaseRefreshOptions) {
		d.Duration = time.Duration(duration)
	}
}

// IsDatabaseLastUpdateTimestampPassed determines whether the database needs to be refreshed.
// It calculates the time elapsed since the last update and compares it with the provided duration.
// The function can accept optional parameters to override the default update interval.
// Returns a boolean indicating whether the database should be refreshed and an error, if any.
func IsDatabaseLastUpdateTimestampPassed(opts ...DatabaseRefreshOpts) (bool, error) {
	lastUpdate, err := GetDatabaseLastUpdateTimestamp()
	if err != nil {
		return false, nil
	}

	// Initialize database refresh options
	options := &DatabaseRefreshOptions{
		Duration: 24, // Default time to perform update is 24 hours
	}

	// Fetch refresh time from config if specified
	refreshTime, _ := GetDatabaseRefreshTime()

	//  If refreshTime is specified in config override the default refresh time
	if refreshTime >= 0 {
		options.Duration = time.Duration(refreshTime)
	}

	// If Options are passed to the function it overrides the refresh time
	for _, opt := range opts {
		opt(options)
	}

	// Verify if duration has passed from last update timestamp
	if time.Since(lastUpdate) > options.Duration*time.Hour {
		lastUpdate = time.Now()
		// Reset the timestamp
		err = SetDatabaseLastUpdateTimestamp(lastUpdate)
		if err != nil {
			return false, errors.Wrap(err, "unable to set database last update timestamp")
		}
		return true, nil
	}

	return false, nil
}

// GetDatabaseRefreshTime retrieves the refreshTime of the database image
// It fetches the configuration details and returns the databaseRefreshTime
// Returns the databaseRefreshTime and an error, if any.
func GetDatabaseRefreshTime() (int, error) {
	// Retrieve client config
	cfg, err := GetClientConfig()
	if err != nil {
		return -1, err
	}
	return cfg.GetDatabaseRefreshTime()
}

// SetDatabaseRefreshTime sets the databaseRefreshTime of the database image.
// It acquires a lock to ensure thread safety, updates the databaseRefreshTime in the configuration,
// and then releases the lock. The function will persist the updated configuration if needed.
// Takes a int value to set as the new databaseRefreshTime.
// Returns an error if the operation fails.
func SetDatabaseRefreshTime(val int) error {
	// Retrieve client config node
	AcquireTanzuConfigLock()
	defer ReleaseTanzuConfigLock()
	node, err := getClientConfigNodeNoLock()
	if err != nil {
		return err
	}

	// Add or Update cliId in the yaml node
	persist := setDatabaseRefreshTime(node, KeyDatabaseRefreshTime, val)

	// Persist the config node to the file
	if persist {
		return persistConfig(node)
	}
	return err
}

// setDatabaseRefreshTime updates the specified yaml node with the new databaseRefreshTime.
// This is a helper function to modify the configuration node with the provided databaseRefreshTime.
// Returns a boolean indicating whether the configuration should be persisted.
func setDatabaseRefreshTime(node *yaml.Node, key string, val int) (persist bool) {
	refreshTimeNode := getNGCLIOptionsChildNode(key, node)
	if refreshTimeNode != nil && refreshTimeNode.Value != strconv.Itoa(val) {
		refreshTimeNode.Value = strconv.Itoa(val)
		persist = true
	}
	return persist
}
