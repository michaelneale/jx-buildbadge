package main

import (
	"testing"
)

func TestState(t *testing.T) {
	icon := StateIcon("show-fail")
	if icon != "fail.png" {
		t.Errorf("got %s", icon)
	}
}

func testBranch(t *testing.T) {
	pipe, branch := SplitBranch("foo/bar")
	if pipe != "foo" {
		t.Errorf("pipeline name was %s", pipe)
	}

	if branch != "bar" {
		t.Errorf("branch was %s", branch)
	}

}
