package test

import (
	 "52lu/go-study-example/example/collyDemo"
	"testing"
)

func TestCollyDemo(t *testing.T) {
	err := collyDemo.RunDemo()
	if err != nil {
		t.Error(err)
	}
}