package hcc_errors

import (
	"errors"
	"log"
	"strconv"
)

const Version string = "1.1.3"

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
