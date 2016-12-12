package main

import (
	"testing"
	"io"
	"reflect"
	"os"
)

type testWriter struct {
	text []byte
}

func (t *testWriter) Write(p []byte) (n int, err error) {
	t.text = p
	return len(p), nil
}

func TestWriter(t *testing.T) {
	w := &testWriter{}
	sayHelloWorld(w)
	if !reflect.DeepEqual(w.text, []byte("Hello World!")) {
		t.Error("Text was wrong")
	}
}

func sayHelloWorld(w io.Writer) {
	w.Write([]byte("Hello World!"))
}

func main1() {
	sayHelloWorld(os.Stdout)
}