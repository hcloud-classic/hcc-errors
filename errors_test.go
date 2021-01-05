package hcc_errors

import (
	"testing"
)

func TestSingleError(t *testing.T) {
	e := NewHccError(HccErrorTestCode, "Test Error Text")

	if e == nil {
		t.Error("Failed to create Error")
	}
}

func TestStack(t *testing.T) {
	s := NewHccErrorStack()

	e := s.Push(NewHccError(HccErrorTestCode, "Push 1"))
	if e != nil {
		t.Error("Push faied")
	}

	e = s.Push(NewHccError(0, ""))
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

func TestDumpAndConvert(t *testing.T) {
	s := NewHccErrorStack()

	s.Push(NewHccError(HccErrorTestCode, "Push 1"))
	s.Push(NewHccError(HccErrorTestCode, "Push 2"))
	s.Push(NewHccError(HccErrorTestCode, "Push 3"))

	e := s.Dump()
	if e != nil {
		t.Error("Dump Failed")
	}

	newS := s.ConvertReportForm()
	if newS != nil {
		t.Error("Convert Bottom Check Failed")
	}

	s = NewHccErrorStack()

	s.Push(NewHccError(HccErrorTestCode, "Push 1"))
	s.Push(NewHccError(HccErrorTestCode, "Push 2"))
	s.Push(NewHccError(HccErrorTestCode, "Push 3"))

	newS = s.ConvertReportForm()
	if newS == nil {
		t.Error("Convert Failed")
	}

	e = newS.Dump()
	if e == nil {
		t.Error("Convert Check Failed")
	}
}
