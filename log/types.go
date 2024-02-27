// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package log

import (
	"fmt"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/log/color"
)

type LogType string

const (
	LogTypeINFO    LogType = "INFO"
	LogTypeWARN    LogType = "WARN"
	LogTypeERROR   LogType = "ERROR"
	LogTypeSUCCESS LogType = "SUCCESS"
	LogTypeOUTPUT  LogType = "OUTPUT"
)

func GetLogBasedOnLogType(msg []byte, logType string) []byte {
	msgString := string(msg)

	if color.IsTTYEnabled() {
		switch LogType(logType) {
		case LogTypeWARN:
			msgString = color.Warnf(msgString)
		case LogTypeERROR:
			msgString = color.Errorf(msgString)
		case LogTypeSUCCESS:
			msgString = color.Successf(msgString)
		case LogTypeINFO:
		case LogTypeOUTPUT:
		}
	} else {
		switch LogType(logType) {
		case LogTypeINFO:
			msgString = fmt.Sprintf("[i] %s", msgString)
		case LogTypeWARN:
			msgString = fmt.Sprintf("[!] %s", msgString)
		case LogTypeERROR:
			msgString = fmt.Sprintf("[x] %s", msgString)
		case LogTypeSUCCESS:
			msgString = fmt.Sprintf("[ok] %s", msgString)
		case LogTypeOUTPUT:
		}
	}
	return []byte(msgString)
}
