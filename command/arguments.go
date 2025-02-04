// Copyright 2025 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ErrIgnoreArg = fmt.Errorf("ignore argument")

type Arg struct {
	Name     string
	Arity    int
	Optional bool
	Set      func(cmd *cobra.Command, args []string, offset int) error
}

func Args(cmd *cobra.Command, argDefs ...Arg) {
	cmd.Args = func(cmd *cobra.Command, args []string) error {
		offset := 0

		for _, argDef := range argDefs {
			arity := argDef.Arity
			if arity == -1 {
				// consume all remaining args
				arity = len(args) - offset
			}
			if len(args)-offset < arity {
				if argDef.Optional {
					continue
				}
				// TODO create a better message saying what is missing
				return fmt.Errorf("missing required argument(s)")
			}

			if err := argDef.Set(cmd, args, offset); err != nil {
				if err == ErrIgnoreArg {
					continue
				}
				return err
			}

			offset += arity
		}

		// no additional args
		return cobra.NoArgs(cmd, args[offset:])
	}

	addArgsToUseString(cmd, argDefs)
}

func Argument(name string, val *string) Arg {
	return Arg{
		Name:  name,
		Arity: 1,
		Set: func(_ *cobra.Command, args []string, offset int) error {
			*val = args[offset]
			return nil
		},
	}
}

func OptionalArgument(name string, val *string) Arg {
	arg := Argument(name, val)
	arg.Optional = true
	return arg
}

func RemainingArguments(name string, values *[]string) Arg {
	return Arg{
		Name:  name,
		Arity: -1,
		Set: func(_ *cobra.Command, args []string, offset int) error {
			*values = args[offset:]
			return nil
		},
	}
}

func OptionalRemainingArguments(name string, values *[]string) Arg {
	arg := RemainingArguments(name, values)
	arg.Optional = true
	return arg
}

func BareDoubleDashArgs(values *[]string) Arg {
	return Arg{
		Arity: -1,
		Set: func(cmd *cobra.Command, args []string, _ int) error {
			if cmd.ArgsLenAtDash() == -1 {
				return nil
			}
			*values = args[cmd.ArgsLenAtDash():]
			return nil
		},
	}
}

// addArgsToUseString automatically adds the argument names to the Use field of the command
func addArgsToUseString(cmd *cobra.Command, argDefs []Arg) {
	for i := range argDefs {
		name := argDefs[i].Name
		if name == "" {
			continue
		}

		if argDefs[i].Optional {
			name = fmt.Sprintf("[%s]", name)
		} else {
			name = fmt.Sprintf("<%s>", name)
		}

		cmd.Use += " " + name
	}
}

// Name argument specific helpers

const (
	NameArgumentName  = "name"
	NamesArgumentName = "name(s)"
)

func NameArg(val *string) Arg {
	return Argument(NameArgumentName, val)
}

func OptionalNameArg(val *string) Arg {
	arg := NameArg(val)
	arg.Optional = true
	return arg
}

func NamesArg(vals *[]string) Arg {
	return RemainingArguments(NamesArgumentName, vals)
}
