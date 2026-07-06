// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The condition-circleci Authors

package plugin

import (
	"strings"
	"testing"
)

func env(kv map[string]string) func(string) string {
	return func(key string) string { return kv[key] }
}

func TestCheck_HappyPath(t *testing.T) {
	t.Parallel()

	c := NewWithEnv(env(map[string]string{
		"CIRCLECI": "true",
	}))

	if err := c.Check(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCheck_NotInCircleCI(t *testing.T) {
	t.Parallel()

	c := NewWithEnv(env(map[string]string{}))
	err := c.Check()
	if err == nil || !strings.Contains(err.Error(), "CIRCLECI") {
		t.Fatalf("expected CIRCLECI error, got: %v", err)
	}
}

func TestCheck_BranchMatch(t *testing.T) {
	t.Parallel()

	c := NewWithEnv(env(map[string]string{
		"CIRCLECI":             "true",
		"SEMREL_PLUGIN_BRANCH": "main",
		"CIRCLE_BRANCH":        "main",
	}))

	if err := c.Check(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCheck_BranchMismatch(t *testing.T) {
	t.Parallel()

	c := NewWithEnv(env(map[string]string{
		"CIRCLECI":             "true",
		"SEMREL_PLUGIN_BRANCH": "main",
		"CIRCLE_BRANCH":        "develop",
	}))

	err := c.Check()
	if err == nil || !strings.Contains(err.Error(), "branch mismatch") {
		t.Fatalf("expected branch mismatch, got: %v", err)
	}
}

func TestCheck_MultipleErrors(t *testing.T) {
	t.Parallel()

	err := NewWithEnv(env(map[string]string{
		"SEMREL_PLUGIN_BRANCH": "main",
	})).Check()
	if err == nil {
		t.Fatal("expected error")
	}

	if !strings.Contains(err.Error(), "CIRCLECI") || !strings.Contains(err.Error(), `want "main" got ""`) {
		t.Fatalf("expected combined errors, got: %v", err)
	}
}
