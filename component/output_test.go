// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package component

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var mapValue = map[string]string{
	"f": "foo",
	"b": "bar",
}

const (
	// leading newline, added for readability, should be trimmed during compare

	expectedYAMLStringified = `
- a: "1"
  b: map[b:bar f:foo]
  c: "2"
`

	expectedYAML = `
- a: "1"
  b:
    b: bar
    f: foo
  c: 2
`

	expectedJSONStringified = `
[
  {
    "a": "1",
    "b": "map[b:bar f:foo]",
    "c": "2"
  }
]`

	expectedJSON = `
[
  {
    "a": "1",
    "b": {
      "b": "bar",
      "f": "foo"
    },
    "c": 2
  }
]`
)

func TestNewOutputWriterTable(t *testing.T) {
	var b bytes.Buffer
	tab := NewOutputWriter(&b, string(TableOutputType), "a", "b", "c")
	require.NotNil(t, tab)
	tab.AddRow("1", "2", "3")
	tab.Render()

	validateTableOutput(t, b.String())
}

func TestNewOutputWriterNewHeaders(t *testing.T) {
	var b bytes.Buffer
	tab := NewOutputWriter(&b, string(TableOutputType), "a", "b", "c")
	require.NotNil(t, tab)
	tab.AddRow("1", "2", "3")
	tab.SetKeys("d", "e", "f")
	tab.Render()

	output := b.String()
	require.NotNil(t, output)

	lines := strings.Split(output, "\n")
	// Output should contain header row, data row, and a blank line
	require.Equal(t, 3, len(lines), "%v", lines)
	require.Contains(t, lines[0], "D")
	require.Contains(t, lines[0], "E")
	require.Contains(t, lines[0], "F")
}

func TestNewOutputWriterUnspecified(t *testing.T) {
	var b bytes.Buffer
	tab := NewOutputWriter(&b, "", "a", "b", "c")
	require.NotNil(t, tab)
	tab.AddRow("1", "2", "3")
	tab.Render()

	validateTableOutput(t, b.String())
}

// TestNewOutputWriterInvalid verifies an invalid output type string will fallback
// to rendering in the table format.
func TestNewOutputWriterInvalid(t *testing.T) {
	var b bytes.Buffer
	tab := NewOutputWriter(&b, "cowsay", "a", "b", "c")
	require.NotNil(t, tab)
	tab.AddRow("1", "2", "3")
	tab.Render()

	validateTableOutput(t, b.String())
}

func TestTableTooManyValues(t *testing.T) {
	var b bytes.Buffer
	tab := NewOutputWriter(&b, string(TableOutputType), "a", "b")
	require.NotNil(t, tab)
	tab.AddRow("1", "2", "3")
	tab.Render()

	output := b.String()
	require.NotNil(t, output)
	lines := strings.Split(output, "\n")

	// Output should contain header row, data row, and a blank line
	require.Equal(t, 3, len(lines), "%v", lines)
	require.Contains(t, lines[0], "A")
	require.Contains(t, lines[0], "B")
	require.Equal(t, lines[1], "  1  2  ")
}

func TestTableTooFewValues(t *testing.T) {
	var b bytes.Buffer
	tab := NewOutputWriter(&b, string(TableOutputType), "a", "b", "c")
	require.NotNil(t, tab)
	tab.AddRow("1", "2")
	tab.Render()

	output := b.String()
	require.NotNil(t, output)
	lines := strings.Split(output, "\n")

	// Output should contain header row, data row, and a blank line
	require.Equal(t, 3, len(lines), "%v", lines)
	require.Contains(t, lines[0], "A")
	require.Contains(t, lines[0], "B")
	require.Equal(t, lines[1], "  1  2  ")
}

func validateTableOutput(t *testing.T, output string) {
	require.NotNil(t, output)
	lines := strings.Split(output, "\n")

	// Output should contain header row, data row, and a blank line
	require.Equal(t, 3, len(lines), "%v", lines)
	require.Contains(t, lines[0], "A")
	require.Contains(t, lines[0], "B")
	require.Contains(t, lines[0], "C")
	require.Equal(t, lines[1], "  1  2  3  ")
}

func TestNewOutputWriterListTable(t *testing.T) {
	var b bytes.Buffer
	tab := NewOutputWriter(&b, string(ListTableOutputType), "a", "B", "c")
	require.NotNil(t, tab)
	tab.AddRow("1", "2", "3")
	tab.AddRow("4", "5", "6")
	tab.Render()

	output := b.String()
	require.NotNil(t, output)
	lines := strings.Split(output, "\n")

	// Output should contain row per header and a blank line
	require.Equal(t, 4, len(lines), "%v", lines)
	require.Contains(t, lines[0], "a:")
	require.Contains(t, lines[0], "1, 4")
	require.Contains(t, lines[1], "B:")
	require.Contains(t, lines[1], "2, 5")
	require.Contains(t, lines[2], "c:")
	require.Contains(t, lines[2], "3, 6")
}

func TestListTableTooManyValues(t *testing.T) {
	var b bytes.Buffer
	tab := NewOutputWriter(&b, string(ListTableOutputType), "a", "b")
	require.NotNil(t, tab)
	tab.AddRow("1", "2", "3")
	tab.Render()

	output := b.String()
	require.NotNil(t, output)
	lines := strings.Split(output, "\n")

	// Output should contain header row, data row, and a blank line
	require.Equal(t, 3, len(lines), "%v", lines)
	require.Contains(t, lines[0], "a:")
	require.Contains(t, lines[0], " 1")
	require.Contains(t, lines[1], "b:")
	require.Contains(t, lines[1], " 2")
}

func TestListTableTooFewValues(t *testing.T) {
	var b bytes.Buffer
	tab := NewOutputWriter(&b, string(ListTableOutputType), "a", "b", "c")
	require.NotNil(t, tab)
	tab.AddRow("1", "2")
	tab.Render()

	output := b.String()
	require.NotNil(t, output)
	lines := strings.Split(output, "\n")

	// Output should contain header row, data row, and a blank line
	require.Equal(t, 4, len(lines), "%v", lines)
	require.Contains(t, lines[0], "a:")
	require.Contains(t, lines[0], " 1")
	require.Contains(t, lines[1], "b:")
	require.Contains(t, lines[1], " 2")
	require.Contains(t, lines[2], "c:")
}

func TestNewOutputWriterYAML(t *testing.T) {
	var b bytes.Buffer
	tab := NewOutputWriter(&b, string(YAMLOutputType), "a", "b", "c")
	require.NotNil(t, tab)
	tab.AddRow("1", "2", "3")
	tab.AddRow("4", "5", "6")
	tab.Render()

	output := b.String()
	require.NotNil(t, output)

	lines := strings.Split(output, "\n")
	// Output should contain our two objects of three values and a blank line
	require.Equal(t, 7, len(lines), "%v", lines)
	require.Contains(t, lines[0], "- a: \"1\"")
	require.Contains(t, lines[1], "  b: \"2\"")
	require.Contains(t, lines[2], "  c: \"3\"")
	require.Contains(t, lines[3], "- a: \"4\"")
	require.Contains(t, lines[4], "  b: \"5\"")
	require.Contains(t, lines[5], "  c: \"6\"")
}

func TestYAMLWriterTooManyValues(t *testing.T) {
	var b bytes.Buffer
	tab := NewOutputWriter(&b, string(YAMLOutputType), "a", "b")
	require.NotNil(t, tab)
	tab.AddRow("1", "2", "3")
	tab.AddRow("4", "5", "6")
	tab.Render()

	output := b.String()
	require.NotNil(t, output)

	lines := strings.Split(output, "\n")
	// Output should contain our two objects of three values and a blank line
	require.Equal(t, 5, len(lines), "%v", lines)
	require.Contains(t, lines[0], "- a: \"1\"")
	require.Contains(t, lines[1], "  b: \"2\"")
	require.Contains(t, lines[2], "- a: \"4\"")
	require.Contains(t, lines[3], "  b: \"5\"")
}

func TestYAMLTooFewValues(t *testing.T) {
	var b bytes.Buffer
	tab := NewOutputWriter(&b, string(YAMLOutputType), "a", "b", "c")
	require.NotNil(t, tab)
	tab.AddRow("1", "2")
	tab.AddRow("4", "5")
	tab.Render()

	output := b.String()
	require.NotNil(t, output)

	lines := strings.Split(output, "\n")
	// Output should contain our two objects of three values and a blank line
	require.Equal(t, 5, len(lines), "%v", lines)
	require.Contains(t, lines[0], "- a: \"1\"")
	require.Contains(t, lines[1], "  b: \"2\"")
	require.Contains(t, lines[2], "- a: \"4\"")
	require.Contains(t, lines[3], "  b: \"5\"")
}

func TestNewOutputWriterJSON(t *testing.T) {
	var b bytes.Buffer
	tab := NewOutputWriter(&b, string(JSONOutputType), "a", "b", "c")
	require.NotNil(t, tab)
	tab.AddRow("1", "2", "3")
	tab.AddRow("4", "5", "6")
	tab.Render()

	output := b.String()
	require.NotNil(t, output)

	lines := strings.Split(output, "\n")
	// Output should contain an array of two objects with three values each
	require.Equal(t, 12, len(lines), "%v", lines)
	require.Contains(t, lines[0], "[")
	require.Contains(t, lines[1], "{")
	require.Contains(t, lines[2], "\"a\": \"1\",")
	require.Contains(t, lines[3], "\"b\": \"2\",")
	require.Contains(t, lines[4], "\"c\": \"3\"")
	require.Contains(t, lines[5], "},")
	require.Contains(t, lines[6], "{")
	require.Contains(t, lines[7], "\"a\": \"4\",")
	require.Contains(t, lines[8], "\"b\": \"5\",")
	require.Contains(t, lines[9], "\"c\": \"6\"")
	require.Contains(t, lines[10], "}")
	require.Contains(t, lines[11], "]")
}

func TestNewOutputWriterNonStrings(t *testing.T) {
	var b bytes.Buffer
	tab := NewOutputWriter(&b, string(TableOutputType), "a", "b", "c")
	require.NotNil(t, tab)
	tab.AddRow("1", mapValue, 2)
	tab.Render()

	output := b.String()
	require.NotNil(t, output)

	// leading newline, added for readability, should be trimmed during compare
	// note: the trailing spaces in this string are intentional
	expected := `
  A  B                 C  
  1  map[b:bar f:foo]  2  
`

	require.Equal(t, expected[1:], output)
}

func TestNewOutputWriterYAMLNonStrings(t *testing.T) {
	var b bytes.Buffer
	tab := NewOutputWriter(&b, string(YAMLOutputType), "a", "b", "c")
	require.NotNil(t, tab)
	tab.AddRow("1", mapValue, 2)
	tab.Render()

	output := b.String()
	require.NotNil(t, output)

	require.Equal(t, expectedYAMLStringified[1:], output)
}

func TestNewOutputWriterWithOptionsYAML(t *testing.T) {
	var b bytes.Buffer

	tab := NewOutputWriterWithOptions(&b, string(YAMLOutputType), []OutputWriterOption{}, "a", "b", "c")
	require.NotNil(t, tab)
	tab.AddRow("1", mapValue, 2)
	tab.Render()

	output := b.String()
	require.NotNil(t, output)

	require.Equal(t, expectedYAML[1:], output)
}

func TestNewOutputWriterJSONNonStrings(t *testing.T) {
	var b bytes.Buffer
	tab := NewOutputWriter(&b, string(JSONOutputType), "a", "b", "c")
	require.NotNil(t, tab)
	tab.AddRow("1", mapValue, 2)
	tab.Render()

	output := b.String()
	require.NotNil(t, output)

	require.Equal(t, expectedJSONStringified[1:], output)
}

func TestNewOutputWriterWithOptionsJSON(t *testing.T) {
	var b bytes.Buffer

	tab := NewOutputWriterWithOptions(&b, string(JSONOutputType), []OutputWriterOption{}, "a", "b", "c")
	require.NotNil(t, tab)
	tab.AddRow("1", mapValue, 2)
	tab.Render()

	output := b.String()
	require.NotNil(t, output)

	require.Equal(t, expectedJSON[1:], output)
}

func TestNewOutputWriterWithOptionsJSONAutoStringify(t *testing.T) {
	var b bytes.Buffer

	tab := NewOutputWriterWithOptions(&b, string(JSONOutputType), []OutputWriterOption{WithAutoStringify()}, "a", "b", "c")
	require.NotNil(t, tab)
	tab.AddRow("1", mapValue, 2)
	tab.Render()

	output := b.String()
	require.NotNil(t, output)

	require.Equal(t, expectedJSONStringified[1:], output)
}

func TestNewOutputWriterWithOptionsYAMLAutoStringify(t *testing.T) {
	var b bytes.Buffer

	tab := NewOutputWriterWithOptions(&b, string(YAMLOutputType), []OutputWriterOption{WithAutoStringify()}, "a", "b", "c")
	require.NotNil(t, tab)
	tab.AddRow("1", mapValue, 2)
	tab.Render()

	output := b.String()
	require.NotNil(t, output)

	require.Equal(t, expectedYAMLStringified[1:], output)
}

func TestNewOutputWriterTableListNonStrings(t *testing.T) {
	var b bytes.Buffer
	tab := NewOutputWriter(&b, string(ListTableOutputType), "a", "b", "c")
	require.NotNil(t, tab)
	tab.AddRow("1", mapValue, 2)
	tab.AddRow("3", "4", "5")
	tab.Render()

	output := b.String()
	require.NotNil(t, output)

	// leading newline, added for readability, should be trimmed during compare
	expected := `
  a:           1, 3
  b:           map[b:bar f:foo], 4
  c:           2, 5
`
	require.Equal(t, expected[1:], output)
}

func TestNewOutputWriterSpinnerJSONNonStrings(t *testing.T) {
	var b bytes.Buffer
	tab, err := NewOutputWriterWithSpinner(&b, string(JSONOutputType), "unused", true, "a", "b", "c")
	require.Nil(t, err)
	require.NotNil(t, tab)
	tab.AddRow("1", mapValue, 2)
	tab.RenderWithSpinner()

	output := b.String()
	require.NotNil(t, output)

	require.Equal(t, expectedJSONStringified[1:], output)
}

func TestNewOutputWriterSpinnerWithOptionsJSON(t *testing.T) {
	var b bytes.Buffer

	tab, err := NewOutputWriterSpinnerWithOptions(&b, string(JSONOutputType), "unused", true, []OutputWriterOption{}, "a", "b", "c")
	require.Nil(t, err)
	require.NotNil(t, tab)
	tab.AddRow("1", mapValue, 2)
	tab.RenderWithSpinner()

	output := b.String()
	require.NotNil(t, output)

	require.Equal(t, expectedJSON[1:], output)
}

func TestNewOutputWriterSpinnerWithOptionsJSONAutoStringify(t *testing.T) {
	var b bytes.Buffer

	tab, err := NewOutputWriterSpinnerWithOptions(&b, string(JSONOutputType), "unused", true, []OutputWriterOption{WithAutoStringify()}, "a", "b", "c")
	require.Nil(t, err)
	require.NotNil(t, tab)
	tab.AddRow("1", mapValue, 2)
	tab.RenderWithSpinner()

	output := b.String()
	require.NotNil(t, output)

	require.Equal(t, expectedJSONStringified[1:], output)
}

func TestNewOutputWriterSpinnerWithOptionsYAMLAutoStringify(t *testing.T) {
	var b bytes.Buffer

	tab, err := NewOutputWriterSpinnerWithOptions(&b, string(YAMLOutputType), "unused", true, []OutputWriterOption{WithAutoStringify()}, "a", "b", "c")
	require.Nil(t, err)
	require.NotNil(t, tab)
	tab.AddRow("1", mapValue, 2)
	tab.RenderWithSpinner()

	output := b.String()
	require.NotNil(t, output)

	require.Equal(t, expectedYAMLStringified[1:], output)
}

func TestOutputWriterCharactersInKeys(t *testing.T) {
	var b bytes.Buffer
	tab := NewOutputWriter(&b, string(JSONOutputType), "a key with spaces", "one/two", "zip:zap")
	require.NotNil(t, tab)
	tab.AddRow("1", "2", "3")
	tab.Render()

	output := b.String()
	require.NotNil(t, output)

	lines := strings.Split(output, "\n")
	// Output should contain an array of an object with three values, keys adjusted
	require.Equal(t, 7, len(lines), "%v", lines)
	require.Contains(t, lines[0], "[")
	require.Contains(t, lines[1], "{")
	require.Contains(t, lines[2], "\"a_key_with_spaces\": \"1\",")
	require.Contains(t, lines[3], "\"one/two\": \"2\",")
	require.Contains(t, lines[4], "\"zip:zap\": \"3\"")
	require.Contains(t, lines[5], "}")
	require.Contains(t, lines[6], "]")
}

func TestObjectWriterJSON(t *testing.T) {
	var b bytes.Buffer
	out := NewObjectWriter(&b, string(JSONOutputType), &testStruct{Name: "hal", Namespace: "Jupiter"})
	out.Render()

	output := b.String()
	require.NotNil(t, output)

	lines := strings.Split(output, "\n")
	require.Equal(t, 4, len(lines), "%v", lines)
	require.Contains(t, lines[0], "{")
	require.Contains(t, lines[1], "\"name\": \"hal\",")
	require.Contains(t, lines[2], "\"spacename\": \"Jupiter\"")
	require.Contains(t, lines[3], "}")
}

func TestObjectWriterYAML(t *testing.T) {
	var b bytes.Buffer
	out := NewObjectWriter(&b, string(YAMLOutputType), &testStruct{Name: "hal", Namespace: "Jupiter"})
	out.Render()

	output := b.String()
	require.NotNil(t, output)

	lines := strings.Split(output, "\n")
	require.Equal(t, 3, len(lines), "%v", lines)
	require.Contains(t, lines[0], "name: hal")
	require.Contains(t, lines[1], "spacename: Jupiter")
}

type testStruct struct {
	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace string `json:"spacename,omitempty" yaml:"spacename,omitempty"`
}

func TestFilterEmptyColumns(t *testing.T) {
	testCases := []struct {
		name            string
		headers         []string
		dynamicHeaders  []string
		values          [][]interface{}
		expectedHeaders []string
		expectedValues  [][]interface{}
	}{
		{
			name:           "Basic test case",
			headers:        []string{"name", "namespace", "active", "status"},
			dynamicHeaders: []string{"active"},
			values: [][]interface{}{
				{"a1", "n1", "", ""},
				{"a2", "n1", "", ""},
				{"a3", "n1", "", ""},
				{"a4", "n1", "", ""},
			},
			expectedHeaders: []string{"name", "namespace", "status"},
			expectedValues: [][]interface{}{
				{"a1", "n1", ""},
				{"a2", "n1", ""},
				{"a3", "n1", ""},
				{"a4", "n1", ""},
			},
		},
		{
			name:           "No dynamic headers",
			headers:        []string{"name", "namespace", "active", "status"},
			dynamicHeaders: []string{},
			values: [][]interface{}{
				{"a1", "n1", "", ""},
				{"a2", "n1", "", ""},
				{"a3", "n1", "", ""},
				{"a4", "n1", "", ""},
			},
			expectedHeaders: []string{"name", "namespace", "active", "status"},
			expectedValues: [][]interface{}{
				{"a1", "n1", "", ""},
				{"a2", "n1", "", ""},
				{"a3", "n1", "", ""},
				{"a4", "n1", "", ""},
			},
		},
		{
			name:           "All empty values in dynamic column",
			headers:        []string{"name", "namespace", "active", "status"},
			dynamicHeaders: []string{"active"},
			values: [][]interface{}{
				{"", "", "", ""},
				{"", "", "", ""},
				{"", "", "", ""},
				{"", "", "", ""},
			},
			expectedHeaders: []string{"name", "namespace", "status"},
			expectedValues: [][]interface{}{
				{"", "", ""},
				{"", "", ""},
				{"", "", ""},
				{"", "", ""},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			resultHeaders, resultValues := filterDynamicColumns(testCase.headers, testCase.dynamicHeaders, testCase.values)

			// Test if result headers are as expected
			if !reflect.DeepEqual(resultHeaders, testCase.expectedHeaders) {
				t.Errorf("Expected headers: %v, got: %v", testCase.expectedHeaders, resultHeaders)
			}

			// Test if result values are as expected
			if !reflect.DeepEqual(resultValues, testCase.expectedValues) {
				t.Errorf("Expected values: %v, got: %v", testCase.expectedValues, resultValues)
			}
		})
	}
}
