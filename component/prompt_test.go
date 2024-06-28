// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package component

import (
	"errors"
	"testing"

	"github.com/AlecAivazis/survey/v2"

	"github.com/stretchr/testify/assert"
)

func Test_translatePromptConfig_Sensitive(t *testing.T) {
	assert := assert.New(t)

	promptConfig := PromptConfig{
		Message:   "Pick a card, any card",
		Options:   []string{"one", "two", "three"},
		Default:   "one",
		Sensitive: true,
		Help:      "Help will be given to those who need it",
	}

	prompt := buildPrompt(&promptConfig, false)
	assert.NotNil(prompt)

	// Secure should return a password prompt
	_, ok := prompt.(*survey.Password)
	assert.True(ok)
}

func Test_translatePromptConfig_OptionsSelect(t *testing.T) {
	assert := assert.New(t)

	promptConfig := PromptConfig{
		Message:   "Pick a card, any card",
		Options:   []string{"one", "two", "three"},
		Default:   "one",
		Sensitive: false,
		Help:      "Help will be given to those who need it",
	}

	// Prompt with options should return a Select or MultiSelect prompt
	// depending on whether multiselection is needed
	prompt := buildPrompt(&promptConfig, false)
	assert.NotNil(prompt)
	selectPrompt, ok := prompt.(*survey.Select)
	assert.True(ok)
	assert.Equal(len(promptConfig.Options), len(selectPrompt.Options))

	prompt = buildPrompt(&promptConfig, true)
	assert.NotNil(prompt)
	multiSelectPrompt, ok := prompt.(*survey.MultiSelect)
	assert.True(ok)
	assert.Equal(len(promptConfig.Options), len(multiSelectPrompt.Options))
}

func Test_translatePromptConfig_Input(t *testing.T) {
	assert := assert.New(t)

	promptConfig := PromptConfig{
		Message:   "Pick a card, any card",
		Default:   "one",
		Sensitive: false,
		Help:      "Help will be given to those who need it",
	}

	prompt := buildPrompt(&promptConfig, false)
	assert.NotNil(prompt)

	// Prompt without options should return an input prompt
	_, ok := prompt.(*survey.Input)
	assert.True(ok)
}

func Test_PromptOptions(t *testing.T) {
	assert := assert.New(t)

	options := defaultPromptOptions()
	opts := translatePromptOpts(options)

	assert.NotNil(options)
	assert.NotNil(opts)
	assert.Equal("?", options.Icons.Question.Text)
	assert.Equal("cyan+b", options.Icons.Question.Format)
	assert.Equal(2, len(opts))
}

func TestAskPrompt_Validation(t *testing.T) {
	// Setup mock prompts
	p := &mockPrompt{
		answers: []string{"", " ", "t", "very-very-long-name", "ALL_CAPS_NAME", "Test", "test"},
	}

	var res string

	// Setup PromptOpts validators
	var promptOpts []PromptOpt

	promptOpts = append(
		promptOpts,
		WithValidator(survey.Required),
		WithValidator(NoOnlySpaces),
		WithValidator(survey.MinLength(2)),
		WithValidator(survey.MaxLength(4)),
		WithValidator(NoUpperCase),
	)

	// Prepare the surveyOpts
	var options = &PromptOptions{}
	for _, opt := range promptOpts {
		err := opt(options)
		assert.Nil(t, err)
	}
	surveyOpts := translatePromptOpts(options)

	// Trigger the Prompt
	err := survey.Ask([]*survey.Question{
		{
			Prompt: p,
		},
	}, &res, surveyOpts...)

	if err != nil {
		t.Fatalf("Ask() = %v", err)
	}

	if res != "test" {
		t.Errorf("answer: %q, want %q", res, "test")
	}
	if p.cleanups != 1 {
		t.Errorf("cleanups: %d, want %d", p.cleanups, 1)
	}

	if err := p.printedErrors[0].Error(); err != "Value is required" {
		t.Errorf("printed error 1: %q, want %q", err, "Value is required")
	}
	if err := p.printedErrors[1].Error(); err != "value contains only spaces" {
		t.Errorf("printed error 2: %q, want %q", err, "value contains only spaces")
	}
	if err := p.printedErrors[2].Error(); err != "value is too short. Min length is 2" {
		t.Errorf("printed error 3: %q, want %q", err, "value is too short. Min length is 2")
	}
	const maxLen4 = "value is too long. Max length is 4"
	if err := p.printedErrors[3].Error(); err != maxLen4 {
		t.Errorf("printed error 2: %q, want %q", err, maxLen4)
	}
	if err := p.printedErrors[4].Error(); err != maxLen4 {
		t.Errorf("printed error 2: %q, want %q", err, maxLen4)
	}
	if err := p.printedErrors[5].Error(); err != "value contains uppercase characters" {
		t.Errorf("printed error 2: %q, want %q", err, "value contains uppercase characters")
	}
}

type mockPrompt struct {
	index         int
	answers       []string
	cleanups      int
	printedErrors []error
}

func (p *mockPrompt) Prompt(*survey.PromptConfig) (interface{}, error) {
	if p.index >= len(p.answers) {
		return nil, errors.New("no valid answers provided")
	}
	val := p.answers[p.index]
	p.index++
	return val, nil
}

func (p *mockPrompt) Cleanup(*survey.PromptConfig, interface{}) error {
	p.cleanups++
	return nil
}

func (p *mockPrompt) Error(_ *survey.PromptConfig, err error) error {
	p.printedErrors = append(p.printedErrors, err)
	return nil
}
