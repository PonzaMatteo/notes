package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func VeryClumsyFunction() error {
	var out = os.Stdout
	if _, err := fmt.Fprint(out, "hello"); err != nil {
		return err
	}
	if _, err := fmt.Fprint(out, ","); err != nil {
		return err
	}
	if _, err := fmt.Fprint(out, "🌎"); err != nil {
		return err
	}
	if _, err := fmt.Fprint(out, "🌎!"); err != nil {
		return err
	}
	return nil
}

type errWriter struct {
	io.Writer
	Error error
}

func (ew *errWriter) WriteString(str string) {
	if ew.Error == nil {
		_, ew.Error = fmt.Fprint(ew, str)
	}
}

func SeemsABitBetter() error {
	var out = &errWriter{Writer: new(strings.Builder)}
	out.WriteString("hello")
	out.WriteString(",")
	out.WriteString("🌎")
	out.WriteString("🌎!")
	return out.Error // must remember this?!
}
