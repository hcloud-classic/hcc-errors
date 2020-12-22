package hcc_errors_test

import (
	herrors "hcc_errors"
	"testing"
)

func TestNewHccErrors(t *testing.T) {
	e := herrors.NewHccError(herrors.HccErrorTestCode, "Test Error Text")

	if e == nil {
		t.Error("Failed to create Error")
	}
}

func TestStack(t *testing.T) {
	s := herrors.NewHccErrorStack()

	e := s.Push(herrors.NewHccError(herrors.HccErrorTestCode, "Push 1"))
	if e != nil {
		t.Error("Push faied")
	}

	e = s.Push(herrors.NewHccError(0, ""))
	if e == nil {
		t.Error("Empty push error check failed")
	}

	err := s.Pop()
	if err == nil {
		t.Error("Pop failed")
	}

	err = s.Pop()
	if err != nil {
		t.Error("Stack bottom check failed")
	}
}
