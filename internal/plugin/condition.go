// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The condition-circleci Authors

package plugin

import (
	"fmt"
	"os"
	"strings"
)

type Condition struct {
	env func(string) string
}

func New() *Condition { return &Condition{env: os.Getenv} }

func NewWithEnv(env func(string) string) *Condition { return &Condition{env: env} }

func (c *Condition) Check() error {
	var errs []string

	if c.env("CIRCLECI") != "true" {
		errs = append(errs, "CIRCLECI is not set to \"true\"; this plugin requires a CircleCI environment")
	}

	if branch := c.env("SEMREL_PLUGIN_BRANCH"); branch != "" {
		gotBranch := c.env("CIRCLE_BRANCH")
		if gotBranch != branch {
			errs = append(errs, fmt.Sprintf("branch mismatch: want %q got %q", branch, gotBranch))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("%s", strings.Join(errs, "; "))
	}
	return nil
}
