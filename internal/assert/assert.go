package assert

import (
	"testing"
	"strings"
)

func Equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got: %v \nwant: %v", got, want)
	}
}

func StringContains(t *testing.T, got, want string) {
	t.Helper()
	if !strings.Contains(got, want) {
		t.Errorf("got: %q \nwant: %q", got, want)
	}
}

func NilError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("got: %v \nwant: nil", got)
	}
}
