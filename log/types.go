// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package log

type LogType string

const (
	LogTypeINFO    LogType = "INFO"
	LogTypeWARN    LogType = "WARN"
	LogTypeERROR   LogType = "ERROR"
	LogTypeSUCCESS LogType = "SUCCESS"
	LogTypeOUTPUT  LogType = "OUTPUT"
)

func GetLogTypeIndicator(logType LogType) string {
	switch logType {
	case LogTypeINFO:
		return "[i] "
	case LogTypeWARN:
		return "[!] "
	case LogTypeERROR:
		return "[x] "
	case LogTypeSUCCESS:
		return "[ok] "
	case LogTypeOUTPUT:
	}
	return ""
}
