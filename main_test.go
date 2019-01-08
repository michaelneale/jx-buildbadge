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

func TestBranch(t *testing.T) {
	pipe, branch := SplitBranch("foo/bar")
	if pipe != "foo" {
		t.Errorf("pipeline name was %s", pipe)
	}

	if branch != "bar" {
		t.Errorf("branch was %s", branch)
	}
}

func TestDefaultBranch(t *testing.T) {
	pipe, branch := SplitBranch("foo")

	if pipe != "foo" {
		t.Errorf("pipeline name was %s", pipe)
	}

	if branch != "master" {
		t.Errorf("branch was %s but I expected it to default to master", branch)
	}

}
