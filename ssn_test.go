package main

import (
	"testing"
)

func TestIsInvalid(t *testing.T) {

	satanIsOK := isValid("666-23-2344")
	if satanIsOK {
		t.Errorf("isValid was incorrect, got: %v, want: %v.", true, false)
	}

	startsWithNines := isValid("900-23-2344")
	if startsWithNines {
		t.Errorf("isValid was incorrect, got: %v, want: %v.", true, false)
	}

	startsWithNines2 := isValid("999-23-2344")
	if startsWithNines2 {
		t.Errorf("isValid was incorrect, got: %v, want: %v.", true, false)
	}

}
