package hcc_errors

import (
	"errors"
	"strconv"
	"strings"
)

/*    HCCERRORSTACK    */

type HccErrorStack struct {
	version  string     `json:"version"`
	errStack []HccError `json:"errstack"`
	IsMixed  bool       `json:"ismixed"`
}

func NewHccErrorStack(errList ...*HccError) *HccErrorStack {
	es := new(HccErrorStack)

	es.version = Version
	es.IsMixed = false

	es.errStack = append(es.errStack, HccError{errCode: 0, errText: ""})

	for _, err := range errList {
		es.Push(err)
	}
	return es
}

func (es *HccErrorStack) setMixed() {
	es.IsMixed = true
}

func (es *HccErrorStack) Version() string {
	return es.version
}

func (es *HccErrorStack) Stack() *[]HccError {
	return &es.errStack
}

func (es *HccErrorStack) Len() int {
	return es.length() - 1
}

func (es *HccErrorStack) length() int {
	return len(*(es.Stack()))
}

func (es *HccErrorStack) Pop() *HccError {
	if es.IsMixed {
		errlogger.Println("Pop(): This HccErrorStack has mixed version of errors. Stack infomations may not correct.")
	}

	l := es.length()
	if l > 1 {
		err := (*(es.Stack()))[l-1]
		*(es.Stack()) = (*(es.Stack()))[:l-1]
		return &err
	}
	return nil
}

func (es *HccErrorStack) Push(err *HccError) error {
	if err.Code() == 0 || err.Text() == "" {
		errlogger.Print("Do not push **Empty** HccError into HccErrorStack\n")
		return errors.New("Empty Error Push")
	}

	if strings.Compare(Version, es.Version()) != 0 {
		errlogger.Println("Push(): hcc_errors version not matched. Stack infomations can be corrupt.")
	}

	es.errStack = append(es.errStack, *err)
	return nil
}

func (es *HccErrorStack) Merge(other *HccErrorStack) {
	var tmp *HccErrorStack = NewHccErrorStack()

	for err := other.Pop(); err != nil; err = other.Pop() {
		tmp.Push(err)
	}

	for err := tmp.Pop(); err != nil; err = tmp.Pop() {
		es.Push(err)
	}

	if other.IsMixed {
		es.IsMixed = true
	}
}

var verWarning string = "WARNING: Mixed version errors. Stack infomations may not correct.\n"

// Dump() will clean stack
func (es *HccErrorStack) Dump() error {

	stack := es.Stack()

	if (*stack)[0].Code() != 0 {
		errlogger.Print("Error Stack already converted to report form. Cannot dump.\n")
		return errors.New("Stack is already converted report form")
	}

	dumpStr := "\n------ [Dump Error Stack] ------\n"
	dumpStr += "Stack Size : " + strconv.Itoa(es.Len()) + "\n"

	if es.IsMixed {
		dumpStr += verWarning
	}

	for err := es.Pop(); err != nil; err = es.Pop() {
		dumpStr += err.ToString() + "\n"
	}
	dumpStr += "--------- [ End Here ] ---------"
	errlogger.Println(dumpStr)
	return nil
}

func (es *HccErrorStack) Print() {
	stack := es.Stack()

	logStr := "\n------ [Print Error Stack] ------\n"
	if es.IsMixed {
		logStr += verWarning
	}
	for rIdx := es.Len(); rIdx >= 0; rIdx-- {
		logStr += (*stack)[rIdx].ToString() + "\n"
	}
	logStr += "--------- [ End Here ] ---------"
	errlogger.Println(logStr)
}

func (es *HccErrorStack) ConvertReportForm() *HccErrorStack {
	newES := NewHccErrorStack()

	if es.Len() > 0 {
		newES.errStack = newES.errStack[1:]
		for idx := 0; idx < es.length(); idx++ {
			reportText := "#" + strconv.Itoa(idx) + " " + es.errStack[idx].ToString()
			newES.Push(NewHccError(es.errStack[idx].Code(), reportText))
		}
	} else {
		return nil
	}
	return newES
}
