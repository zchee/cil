// Copyright 2018 The cil Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package app

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

func NewDefaultCilCommand(args []string) *cobra.Command {
	return NewCilCommand(os.Stdin, os.Stdout, os.Stderr, args)
}

func NewCilCommand(in io.Reader, out, err io.Writer, args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "cil",
		Short:        "cil is a manage the multiple Continuous Integration services",
		SilenceUsage: true,
	}

	flags := cmd.PersistentFlags()
	flags.Parse(args)

	cmd.AddCommand(NewCmdCircleCI(out))

	return cmd
}
