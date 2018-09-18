// Copyright 2018 The cil Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/zchee/cil/cmd/cil/app"
)

func main() {
	cmd := app.NewDefaultCilCommand(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
