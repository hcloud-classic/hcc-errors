package hcc_errors

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

const Version string = "1.1.0"

// const variable for test
const HccErrorTestCode = modulename + category + operation

var errlogger = log.New(log.Writer(), "[hcc_error] ", log.LstdFlags)

func SetErrLogger(l *log.Logger) {
	errlogger = l
}

/*    HCCERROR    */

type HccError struct {
	errCode uint64 `json:"errcode"` // decimal error code
	errText string `json:"errtext"` // error string
}

func NewHccError(errorCode uint64, errorText string) *HccError {
	err := new(HccError)
	err.errText = errorText
	err.errCode = errorCode

	return err
}

func (e HccError) ToError() error {
	return errors.New(e.ToString())
}

func (e HccError) Error() string {
	return e.ToString()
}

func (e HccError) Code() uint64 {
	return e.errCode
}

func (e HccError) Text() string {
	return e.errText
}

func (e HccError) ToString() string {
	m := e.Code() / 10000
	f := e.Code() % 10000 / 1000
	a := e.Code() % 1000

	return "[" + middleWareList[m] + "] Code :" + strconv.FormatUint(e.Code(), 10) + " (" + functionList[f] + ") " + actionList[a] + " " + e.Text()
}

func (e HccError) Println() {
	errlogger.Println(e.ToString())
}

func (e HccError) Fatal() {
	errlogger.Fatal(e.ToString())
}

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
