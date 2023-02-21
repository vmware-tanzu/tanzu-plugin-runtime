// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package core contains compatibility testing framework core types and functions
package core

type RuntimeAPIName string

const (
	SetContextAPIName RuntimeAPIName = "SetContext"
	GetContextAPIName RuntimeAPIName = "GetContext"
	AddServerAPIName  RuntimeAPIName = "AddServer"
	GetServerAPIName  RuntimeAPIName = "GetServer"
)
