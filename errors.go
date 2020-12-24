package hcc_errors

import (
	"errors"
	"log"
	"strconv"
)

const version = "1.0.3"

// const variable for test
const HccErrorTestCode = modulename + category + operation

var errlogger = log.New(log.Writer(), "hcc_error", log.LstdFlags)

func SetErrLogger(l *log.Logger) {
	errlogger = l
}

/*    HCCERROR    */

type HccError struct {
	errCode uint64 `json:"errcode"` // decimal error code
	errText string `json:"errtext"` // error string
}

func NewHccError(errorCode uint64, errorText string) *HccError {
	return &HccError{
		errText: errorText,
		errCode: errorCode,
	}
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

type HccErrorStack []HccError

func NewHccErrorStack(errList ...*HccError) *HccErrorStack {
	es := HccErrorStack{HccError{errCode: 0, errText: ""}}

	for _, err := range errList {
		es.Push(err)
	}
	return &es
}

func (es *HccErrorStack) Len() int {
	return es.len() - 1
}

func (es *HccErrorStack) len() int {
	return len(*es)
}

func (es *HccErrorStack) Pop() *HccError {
	l := es.len()
	if l > 1 {
		err := (*es)[l-1]
		*es = (*es)[:l-1]
		return &err
	}
	return nil
}

func (es *HccErrorStack) Push(err *HccError) error {
	if err.Code() == 0 || err.Text() == "" {
		errlogger.Print("Do not push **Empty** HccError into HccErrorStack\n")
		return errors.New("Empty Error Push")
	}
	*es = append(*es, *err)
	return nil
}

// Dump() will clean stack
func (es *HccErrorStack) Dump() error {

	if (*es)[0].Code() == 0 {
		errlogger.Print("Error Stack is already converted to report form. Cannot dump.\n")
		return errors.New("Stack is already converted report form")
	}

	errlogger.Printf("------ [Dump Error Stack] ------\n")
	errlogger.Printf("Stack Size : %v\n", es.Len())
	for err := es.Pop(); err != nil; err = es.Pop() {
		err.Println()
	}
	errlogger.Println("--------- [ End Here ] ---------")
	return nil
}

func (es *HccErrorStack) Print() {
	var stack []HccError = *es
	logStr := "\n------ [Start Error Stack] ------\n"
	for rIdx := es.Len(); rIdx >= 0; rIdx-- {
		logStr += stack[rIdx].Text() + "\n"
	}
	logStr += "--------- [ End Here ] ---------"
	errlogger.Println(logStr)
}

func (es *HccErrorStack) ConvertReportForm() *HccErrorStack {
	var newES HccErrorStack

	if es.Len() > 0 {
		*es = (*es)[1:]
		for idx := 0; idx < es.len(); idx++ {
			reportText := "#" + strconv.Itoa(idx) + " " + (*es)[idx].ToString()
			newES.Push(NewHccError((*es)[idx].Code(), reportText))
		}
	} else {
		return nil
	}
	return &newES
}
