// Copyright 2018 The cil Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package app

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/k0kubun/pp"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	circleci "github.com/zchee/cil/provider/circleci"
)

type cmdCircleCI struct {
	out    io.Writer
	client *circleci.DefaultApiService
	token  string
}

func NewCmdCircleCI(out io.Writer) *cobra.Command {
	ctx := context.Background()

	c := &cmdCircleCI{
		out:    out,
		client: circleci.NewAPIClient(circleci.NewConfiguration()).DefaultApi,
	}
	cmd := &cobra.Command{
		Use:   "circleci",
		Short: "manage the Circle CI",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return c.init(cmd, args)
		},
	}

	flags := cmd.PersistentFlags()
	flags.StringVarP(&c.token, "token", "t", "", "CircleCI API token. ($CIL_CIRCLECI_TOKEN)")

	cmd.AddCommand(&cobra.Command{
		Use:   "user",
		Short: "show user information",
		RunE: func(*cobra.Command, []string) error {
			return c.user(ctx)
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "projects",
		Short: "show projects",
		RunE: func(*cobra.Command, []string) error {
			return c.projects(ctx)
		},
	})

	return cmd
}

func (c *cmdCircleCI) init(cmd *cobra.Command, args []string) error {
	tok, err := cmd.Flags().GetString("token")
	if tok = os.Getenv("CIL_CIRCLECI_TOKEN"); tok == "" || err != nil {
		return errors.New("-token flag or $CIL_CIRCLECI_TOKEN env are must be not empty")
	}
	c.token = tok

	return nil
}

func (c *cmdCircleCI) user(ctx context.Context) error {
	ctx = context.WithValue(ctx, circleci.ContextAPIKey, circleci.APIKey{Key: c.token, Prefix: ""})
	user, resp, err := c.client.Me(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get user")
	}

	var _ = resp
	// fmt.Fprintf(c.out, "resp: %v\n", pp.Sprintln(resp))
	fmt.Fprintf(c.out, "user: %v\n", pp.Sprint(user))

	return nil
}

func (c *cmdCircleCI) projects(ctx context.Context) error {
	ctx = context.WithValue(ctx, circleci.ContextAPIKey, circleci.APIKey{Key: c.token, Prefix: ""})
	projects, resp, err := c.client.Projects(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get projects")
	}

	var _ = resp
	// fmt.Fprintf(c.out, "resp: %v\n", pp.Sprintln(resp))
	fmt.Fprintf(c.out, "projects: %v\n", pp.Sprint(projects))

	return nil
}
