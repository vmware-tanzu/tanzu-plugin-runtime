// Copyright 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetDatabaseLastUpdateTimestamp(t *testing.T) {
	// Setup config test data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	currTime := time.Now()
	currTimePlusOne := currTime.Add(1)
	currISOTime := currTime.Format(time.RFC3339)
	currentISOTimePlusOne := currTime.Add(1).Format(time.RFC3339)

	tests := []struct {
		name        string
		value       time.Time
		expectedVal string
	}{
		{
			name:        "should persist DatabaseLastUpdateTimestamp value when empty client config",
			value:       currTime,
			expectedVal: currISOTime,
		},
		{
			name:        "should update and persist DatabaseLastUpdateTimestamp value",
			value:       currTimePlusOne,
			expectedVal: currentISOTimePlusOne,
		},
		{
			name:        "should not persist same value false",
			value:       currTimePlusOne,
			expectedVal: currentISOTimePlusOne,
		},
	}

	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := SetDatabaseLastUpdateTimestamp(spec.value)
			assert.NoError(t, err)
			lastUpdateTimestamp, err := GetDatabaseLastUpdateTimestamp()
			assert.Equal(t, spec.expectedVal, lastUpdateTimestamp.Format(time.RFC3339))
			assert.NoError(t, err)
		})
	}
}

func TestSetGetDatabaseRefreshTime(t *testing.T) {
	// Setup config test data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	tests := []struct {
		name  string
		value int
	}{
		{
			name:  "should persist refreshTime value when empty client config",
			value: 24,
		},
		{
			name:  "should update and persist refreshTime value",
			value: 10,
		},
	}

	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			err := SetDatabaseRefreshTime(spec.value)
			assert.NoError(t, err)
			refreshTime, err := GetDatabaseRefreshTime()
			assert.Equal(t, spec.value, refreshTime)
			assert.NoError(t, err)
		})
	}
}

func TestIsDatabaseLastUpdateTimestampPassed(t *testing.T) {
	// Setup config test data
	_, cleanUp := setupTestConfig(t, &CfgTestData{})

	defer func() {
		cleanUp()
	}()

	currTime := time.Now()

	currTimePlusTwo := currTime.Add(2 * time.Hour)
	currTimeMinusTwo := currTime.Add(-2 * time.Hour)

	currTimePlus5 := currTime.Add(5 * time.Hour)
	currTimeMinus5 := currTime.Add(-5 * time.Hour)

	currTimeMinus24 := currTime.Add(-24 * time.Hour)
	currTimePlus24 := currTime.Add(24 * time.Hour)

	tests := []struct {
		name            string
		value           time.Time
		durationOptions *DatabaseRefreshOptions
		durationPassed  bool
		refreshTime     int
	}{

		// No Duration options and no Refresh Time
		{
			name:           "should return false for IsDatabaseLastUpdateTimestampPassed",
			value:          currTimeMinusTwo,
			durationPassed: false},
		{
			name:           "should return false for IsDatabaseLastUpdateTimestampPassed",
			value:          currTimePlusTwo,
			durationPassed: false,
		},
		{
			name:           "should return true for IsDatabaseLastUpdateTimestampPassed",
			value:          currTimeMinus24,
			durationPassed: true,
		},
		{
			name:           "should return false for IsDatabaseLastUpdateTimestampPassed",
			value:          currTimePlus24,
			durationPassed: false,
		},

		// Only Duration options and no refresh time
		{
			name:  "should return false for IsDatabaseLastUpdateTimestampPassed",
			value: currTime,
			durationOptions: &DatabaseRefreshOptions{
				Duration: 1,
			},
			durationPassed: false,
		},
		{
			name:  "should return true for IsDatabaseLastUpdateTimestampPassed",
			value: currTimeMinusTwo,
			durationOptions: &DatabaseRefreshOptions{
				Duration: 1,
			},
			durationPassed: true,
		},
		{
			name:  "should return false for IsDatabaseLastUpdateTimestampPassed",
			value: currTimePlusTwo,
			durationOptions: &DatabaseRefreshOptions{
				Duration: 1,
			},
			durationPassed: false,
		},
		{
			name:  "should return true for IsDatabaseLastUpdateTimestampPassed",
			value: currTime,
			durationOptions: &DatabaseRefreshOptions{
				Duration: 0,
			},
			durationPassed: true,
		},
		{
			name:  "should return true for IsDatabaseLastUpdateTimestampPassed",
			value: currTimeMinusTwo,
			durationOptions: &DatabaseRefreshOptions{
				Duration: 0,
			},
			durationPassed: true,
		},
		{
			name:  "should return false for IsDatabaseLastUpdateTimestampPassed",
			value: currTimePlusTwo,
			durationOptions: &DatabaseRefreshOptions{
				Duration: 0,
			},
			durationPassed: false,
		},
		{
			name:  "should return true for IsDatabaseLastUpdateTimestampPassed",
			value: currTime,
			durationOptions: &DatabaseRefreshOptions{
				Duration: -1,
			},
			durationPassed: true,
		},
		{
			name:  "should return true for IsDatabaseLastUpdateTimestampPassed",
			value: currTimeMinusTwo,
			durationOptions: &DatabaseRefreshOptions{
				Duration: -1,
			},
			durationPassed: true,
		},
		{
			name:  "should return false for IsDatabaseLastUpdateTimestampPassed",
			value: currTimePlusTwo,
			durationOptions: &DatabaseRefreshOptions{
				Duration: -1,
			},
			durationPassed: false,
		},

		// Only RefreshTime and No Duration options
		{
			name:           "should return true for IsDatabaseLastUpdateTimestampPassed",
			value:          currTimeMinusTwo,
			durationPassed: false,
			refreshTime:    5,
		},
		{
			name:           "should return false for IsDatabaseLastUpdateTimestampPassed",
			value:          currTimePlusTwo,
			durationPassed: false,
			refreshTime:    5,
		},
		{
			name:           "should return false for IsDatabaseLastUpdateTimestampPassed",
			value:          currTimePlus5,
			durationPassed: false,
			refreshTime:    5,
		},
		{
			name:           "should return true for IsDatabaseLastUpdateTimestampPassed",
			value:          currTimeMinus5,
			durationPassed: true,
			refreshTime:    5,
		},

		// RefreshTime and Duration options
		{
			name:  "should return true for IsDatabaseLastUpdateTimestampPassed",
			value: currTimeMinusTwo,
			durationOptions: &DatabaseRefreshOptions{
				Duration: 1,
			},
			durationPassed: true,
			refreshTime:    5,
		},
		{
			name:  "should return false for IsDatabaseLastUpdateTimestampPassed",
			value: currTimePlusTwo,
			durationOptions: &DatabaseRefreshOptions{
				Duration: 1,
			},
			durationPassed: false,
			refreshTime:    5,
		},
		{
			name:  "should return false for IsDatabaseLastUpdateTimestampPassed",
			value: currTimePlusTwo,
			durationOptions: &DatabaseRefreshOptions{
				Duration: 0,
			},
			durationPassed: false,
			refreshTime:    5,
		},
		{
			name:  "should return false for IsDatabaseLastUpdateTimestampPassed",
			value: currTimePlusTwo,
			durationOptions: &DatabaseRefreshOptions{
				Duration: -1,
			},
			durationPassed: false,
			refreshTime:    5,
		},
	}

	for _, spec := range tests {
		t.Run(spec.name, func(t *testing.T) {
			// Verify RefreshTime

			if spec.refreshTime > 0 {
				err := SetDatabaseRefreshTime(spec.refreshTime)
				assert.NoError(t, err)
				refreshTime, err := GetDatabaseRefreshTime()
				assert.NoError(t, err)
				assert.Equal(t, spec.refreshTime, refreshTime)
			}

			// Verify DatabaseLastUpdateTimestamp
			err := SetDatabaseLastUpdateTimestamp(spec.value)
			assert.NoError(t, err)

			lastUpdateTimestamp, err := GetDatabaseLastUpdateTimestamp()
			assert.Equal(t, spec.value.Format(time.RFC3339), lastUpdateTimestamp.Format(time.RFC3339))
			assert.NoError(t, err)

			var timePassed bool

			if spec.durationOptions != nil {
				timePassed, err = IsDatabaseLastUpdateTimestampPassed(WithDuration(int(spec.durationOptions.Duration)))
				assert.NoError(t, err)
			} else {
				timePassed, err = IsDatabaseLastUpdateTimestampPassed()
				assert.NoError(t, err)
			}

			assert.Equal(t, spec.durationPassed, timePassed)
		})
	}
}
