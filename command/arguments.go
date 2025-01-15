// Copyright 2025 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	NameArgumentName  = "name"
	NamesArgumentName = "name(s)"
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

func NameArg(name *string) Arg {
	return Arg{
		Name:  NameArgumentName,
		Arity: 1,
		Set: func(_ *cobra.Command, args []string, offset int) error {
			*name = args[offset]
			return nil
		},
	}
}

func OptionalNameArg(name *string) Arg {
	arg := NameArg(name)
	arg.Optional = true
	return arg
}

func NamesArg(names *[]string) Arg {
	return Arg{
		Name:  NamesArgumentName,
		Arity: -1,
		Set: func(_ *cobra.Command, args []string, offset int) error {
			*names = args[offset:]
			return nil
		},
	}
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
