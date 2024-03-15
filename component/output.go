// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package component

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v3"
)

const (
	colWidth    = 80
	indentation = `  `
)

// OutputWriter is an interface for something that can write output.
type OutputWriter interface {
	// SetKeys sets the values to use as the keys for the output values.
	SetKeys(headerKeys ...string)
	// MarkDynamicKeys sets the headers as dynamic and only shows the column
	// if at least one row is non empty for the specified header
	// Note: This is only supported with output render type TableOutputType
	MarkDynamicKeys(dynamicKeys ...string)
	// AddRow appends a new row to our table.
	AddRow(items ...interface{})
	// Render emits the generated table to the output once ready
	Render()
}

// OutputType defines the format of the output desired.
type OutputType string

const (
	// TableOutputType specifies output should be in table format.
	TableOutputType OutputType = "table"
	// YAMLOutputType specifies output should be in yaml format.
	YAMLOutputType OutputType = "yaml"
	// JSONOutputType specifies output should be in json format.
	JSONOutputType OutputType = "json"
	// ListTableOutputType specified output should be in a list table format.
	ListTableOutputType OutputType = "listtable"
)

// outputwriter is our internal implementation.
type outputwriter struct {
	out                 io.Writer
	keys                []string
	dynamicKeys         []string
	values              [][]interface{}
	outputFormat        OutputType
	autoStringifyFields bool
}

// OutputWriterOption is an option for outputwriter
type OutputWriterOption func(*outputwriter)

// WithAutoStringify configures the output writer to automatically convert
// row fields to their golang string representations. It is used to maintain
// backward compatibility with old rendering behavior, and should be _avoided_
// if that need does not apply.
func WithAutoStringify() OutputWriterOption {
	return func(ow *outputwriter) {
		ow.autoStringifyFields = true
	}
}

// NewOutputWriter gets a new instance of our output writer.
//
// Deprecated: NewOutputWriter is being deprecated in favor of NewOutputWriterWithOptions
// Until it is removed, it will retain the existing behavior of converting
// incoming row values to their golang string representation for backward
// compatibility reasons
func NewOutputWriter(output io.Writer, outputFormat string, headers ...string) OutputWriter {
	// for retaining old json/yaml rendering behavior
	opts := []OutputWriterOption{WithAutoStringify()}

	return NewOutputWriterWithOptions(output, outputFormat, opts, headers...)
}

// NewOutputWriterWithOptions gets a new instance of our output writer with some customization options.
func NewOutputWriterWithOptions(output io.Writer, outputFormat string, opts []OutputWriterOption, headers ...string) OutputWriter {
	// Initialize the output writer that we use under the covers
	ow := &outputwriter{}
	ow.out = output
	ow.outputFormat = OutputType(outputFormat)
	ow.keys = headers

	ow.applyOptions(opts)

	return ow
}

func (ow *outputwriter) applyOptions(opts []OutputWriterOption) {
	for i := range opts {
		opts[i](ow)
	}
}

// SetKeys sets the values to use as the keys for the output values.
func (ow *outputwriter) SetKeys(headerKeys ...string) {
	// Overwrite whatever was used in initialization
	ow.keys = headerKeys
}

// MarkDynamicKeys sets the headers as dynamic and only shows the column if at least one row is non empty
// This is only supported with output render type TableOutputType
func (ow *outputwriter) MarkDynamicKeys(dynamicKeys ...string) {
	ow.dynamicKeys = dynamicKeys
}

func stringify(items []interface{}) []interface{} {
	var results []interface{}
	for i := range items {
		results = append(results, fmt.Sprintf("%v", items[i]))
	}
	return results
}

// AddRow appends a new row to our table.
func (ow *outputwriter) AddRow(items ...interface{}) {
	row := []interface{}{}

	var rowValues []interface{}
	rowValues = items
	if ow.autoStringifyFields {
		rowValues = stringify(items)
	}

	row = append(row, rowValues...)
	ow.values = append(ow.values, row)
}

// Render emits the generated table to the output once ready
func (ow *outputwriter) Render() {
	switch ow.outputFormat {
	case JSONOutputType:
		renderJSON(ow.out, ow.dataStruct())
	case YAMLOutputType:
		renderYAML(ow.out, ow.dataStruct())
	case ListTableOutputType:
		renderListTable(ow)
	default:
		renderTable(ow)
	}
}

func (ow *outputwriter) dataStruct() []map[string]interface{} {
	data := []map[string]interface{}{}
	keys := ow.keys
	for i, k := range keys {
		keys[i] = strings.ToLower(strings.ReplaceAll(k, " ", "_"))
	}

	for _, itemValues := range ow.values {
		item := map[string]interface{}{}
		for i, value := range itemValues {
			if i == len(keys) {
				continue
			}
			item[keys[i]] = value
		}
		data = append(data, item)
	}

	return data
}

// objectwriter is our internal implementation.
type objectwriter struct {
	out          io.Writer
	data         interface{}
	outputFormat OutputType
}

// NewObjectWriter gets a new instance of our output writer.
func NewObjectWriter(output io.Writer, outputFormat string, data interface{}) OutputWriter {
	// Initialize the output writer that we use under the covers
	obw := &objectwriter{}
	obw.out = output
	obw.data = data
	obw.outputFormat = OutputType(outputFormat)

	return obw
}

// SetKeys sets the values to use as the keys for the output values.
func (obw *objectwriter) SetKeys(_ ...string) {
	// Object writer does not have the concept of keys
	fmt.Fprintln(obw.out, "Programming error, attempt to add headers to object output")
}

// MarkDynamicKeys marks the headers as dynamic
func (obw *objectwriter) MarkDynamicKeys(dynamicKeys ...string) {
	// Object writer does not have the concept of dynamic keys
	fmt.Fprintln(obw.out, "Programming error, attempt to mark dynamic headers to object output")
}

// AddRow appends a new row to our table.
func (obw *objectwriter) AddRow(_ ...interface{}) {
	// Object writer does not have the concept of keys
	fmt.Fprintln(obw.out, "Programming error, attempt to add rows to object output")
}

// Render emits the generated table to the output once ready
func (obw *objectwriter) Render() {
	switch obw.outputFormat {
	case JSONOutputType:
		renderJSON(obw.out, obw.data)
	case YAMLOutputType:
		renderYAML(obw.out, obw.data)
	default:
		fmt.Fprintf(obw.out, "Invalid output format: %v\n", obw.outputFormat)
	}
}

// renderJSON prints output as json
func renderJSON(out io.Writer, data interface{}) {
	bytesJSON, err := json.MarshalIndent(data, "", indentation)
	if err != nil {
		fmt.Fprint(out, err)
		return
	}

	fmt.Fprintf(out, "%v", string(bytesJSON))
}

// renderYAML prints output as yaml
func renderYAML(out io.Writer, data interface{}) {
	yamlInBytes, err := yaml.Marshal(data)
	if err != nil {
		fmt.Fprint(out, err)
		return
	}

	fmt.Fprintf(out, "%s", yamlInBytes)
}

// renderListTable prints output as a list table.
func renderListTable(ow *outputwriter) {
	headerLength := 10
	for _, header := range ow.keys {
		length := len(header) + 2
		if length > headerLength {
			headerLength = length
		}
	}

	for i, header := range ow.keys {
		row := []string{}
		for _, data := range ow.values {
			if i >= len(data) {
				// There are more headers than values, leave it blank
				continue
			}
			row = append(row, fmt.Sprintf("%v", data[i]))
		}
		headerLabel := strings.ToUpper(header) + ":"
		values := strings.Join(row, ", ")
		fmt.Fprintf(ow.out, "%-"+strconv.Itoa(headerLength)+"s   %s\n", headerLabel, values)
	}
}

// renderTable prints output as a table
func renderTable(ow *outputwriter) {
	// Filter keys and values based on dynamic keys
	ow.keys, ow.values = filterDynamicColumns(ow.keys, ow.dynamicKeys, ow.values)

	table := tablewriter.NewWriter(ow.out)
	table.SetBorder(false)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)
	table.SetColWidth(colWidth)
	table.SetTablePadding("\t\t")
	table.SetHeader(ow.keys)
	table.AppendBulk(convertInterfaceToString(ow.values))
	table.Render()
}

func convertInterfaceToString(values [][]interface{}) [][]string {
	result := [][]string{}
	for _, valuesRow := range values {
		row := []string{}
		for _, field := range valuesRow {
			row = append(row, fmt.Sprintf("%v", field))
		}

		result = append(result, row)
	}
	return result
}

func filterDynamicColumns(headers, dynamicHeaders []string, values [][]interface{}) ([]string, [][]interface{}) {
	resultHeaders := make([]string, 0)
	resultValues := make([][]interface{}, len(values))

	// Create a map for faster lookup of dynamic headers
	dynamicHeaderMap := make(map[string]bool)
	for _, header := range dynamicHeaders {
		dynamicHeaderMap[header] = true
	}

	// Iterate over the headers and values
	for i, header := range headers {
		// Check if the header is dynamic and all values for that column are empty
		if _, isDynamic := dynamicHeaderMap[header]; !isDynamic || !allEmptyColumn(values, i) {
			resultHeaders = append(resultHeaders, header)

			// Copy non-empty values for the current column
			for j, row := range values {
				if j < len(resultValues) && i < len(row) {
					resultValues[j] = append(resultValues[j], row[i])
				}
			}
		}
	}

	return resultHeaders, resultValues
}

func allEmptyColumn(values [][]interface{}, columnIndex int) bool {
	for _, row := range values {
		if row[columnIndex] != "" {
			return false
		}
	}
	return true
}
