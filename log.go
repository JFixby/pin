package pin

import (
	"reflect"
	"strconv"
	"fmt"
	"os"
	"runtime"
	"strings"
	"bytes"

	"github.com/davecgh/go-spew/spew"
)

const (
	ERROR = "com.jfixby.scarabei.log.error"
	DEBUG = "com.jfixby.scarabei.log.debug"
	SPEW  = "com.jfixby.scarabei.log.spew"
)

var indent = strconv.Itoa(24)

var LogPrinter LogPrinterComponent = &DefaultLogPrinter{}

type LogPrinterComponent interface {
	Debug(msg string)
	Error(msg string)
}

type DefaultLogPrinter struct {
}

func (*DefaultLogPrinter) Debug(msg string) {
	fmt.Fprintln(os.Stdout, msg)
}

func (*DefaultLogPrinter) Error(msg string) {
	fmt.Fprintln(os.Stderr, msg)
}

func D(tag string, message ...interface{}) {
	log(DEBUG, tag, message...)
}

func S(tag string, message ...interface{}) {
	log(SPEW, tag, message...)
}

func E(tag string, message ...interface{}) {
	log(ERROR, tag, message...)
}

func Error(tag string, message ...interface{}) {
	log(ERROR, tag, message...)
}

func Debug(tag string, message ...interface{}) {
	log(DEBUG, tag, message...)
}

func Spew(tag string, message ...interface{}) {
	log(SPEW, tag, message...)
}

func log(mode string, tag string, message ...interface{}) {
	var a interface{} = nil
	if len(message) > 0 {
		a = message[0]
	}

	if a == nil {
		msg := decorate(tag)
		if mode == DEBUG {
			LogPrinter.Debug(msg)
		} else if mode == ERROR {
			LogPrinter.Error(msg)
		} else {
			LogPrinter.Debug(msg)
		}
		return
	}

	// a!=nil
	array := reflect.ValueOf(a)
	if array.Kind() == reflect.Array || array.Kind() == reflect.Slice {
		if mode == DEBUG {
			msg := ArrayToString(tag, a)
			msg = decorate(msg)
			LogPrinter.Debug(msg)
		} else if mode == ERROR {
			msg := ArrayToString(tag, a)
			msg = decorate(msg)
			LogPrinter.Error(msg)
		} else {
			msg := "<" + tag + ">:\n" + spew.Sdump(a)
			msg = decorate(msg)
			LogPrinter.Debug(msg)
		}
	} else {
		if mode == DEBUG {
			msg := fmt.Sprintf("%v > %v", tag, a)
			msg = decorate(msg)

			LogPrinter.Debug(msg)
		} else if mode == ERROR {
			msg := fmt.Sprintf("%v > %v", tag, a)
			msg = decorate(msg)

			LogPrinter.Error(msg)
		} else { //spew
			msg := fmt.Sprintf("%v > %v", tag, spew.Sdump(a))
			msg = decorate(msg)

			LogPrinter.Debug(msg)
		}
	}
}

func decorate(msg string) string {
	_, file, line, ok := runtime.Caller(3)
	if ok {
		// Truncate file name at last file name separator.
		if index := strings.LastIndex(file, "/"); index >= 0 {
			file = file[index+1:]
		} else if index = strings.LastIndex(file, "\\"); index >= 0 {
			file = file[index+1:]
		}
	} else {
		file = "???"
		line = 1
	}
	buf := new(bytes.Buffer)
	prefixFile := fmt.Sprintf("%v:%d ", file, line)
	{
		padded := fmt.Sprintf("%"+indent+"v", prefixFile)
		buf.WriteString(padded)
	}
	lines := strings.Split(msg, "\n")
	if l := len(lines); l > 1 && lines[l-1] == "" {
		lines = lines[:l-1]
	}
	for i, line := range lines {
		if i > 0 {
			// Second and subsequent lines are indented an extra tab.
			buf.WriteString("\n")
			padded := fmt.Sprintf("%"+indent+"v", "")
			buf.WriteString(padded)
		}
		buf.WriteString(line)
	}
	return buf.String()
}

// ArrayToString converts any array into pretty-print string
//
// Format:
/*
---(%tag)[%arraySize]---
      (0) %valueAtIndex0
      (1) %valueAtIndex1
      (2) %valueAtIndex2
      (3) %valueAtIndex3
      ...

Example call: ArrayToString("array",[5]int{14234, 42, -1, 1000, 5})
Output:
		---(array)[5]---
		       (0) 14234
		       (1) 42
		       (2) -1
		       (3) 1000
		       (4) 5
*/
func ArrayToString(tag string, iface interface{}) string {
	array := reflect.ValueOf(iface)
	//if array.Kind() != reflect.Array {
	//	panic("This is not array: " + fmt.Sprint(iface))
	//}
	n := array.Len()
	prefix := "---(" + tag + ")"
	head := prefix + "-[" + strconv.Itoa(n) + "]---"
	prefixLen := len(prefix) + 1
	result := head + "\n"
	for i := 0; i < n; i++ {
		e := array.Index(i)
		val := fmt.Sprintf("%v", e)
		line := fmt.Sprintf("%"+strconv.Itoa(prefixLen)+"s",
			"("+strconv.Itoa(i)+") ") + val
		result = result + line + "\n"
	}
	return result
}
