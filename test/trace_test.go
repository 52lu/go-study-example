package test

import (
	"testing"
	"time"
)

func TestUseTrace(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := i
		selectDB()
		writeFile()
		sum(a)
	}
}

func selectDB() {
	time.Sleep(time.Millisecond * 50)
	_ = make(map[string]string, 10000)
}

func writeFile() {
	time.Sleep(time.Millisecond * 100)
	_ = make(map[string]string, 3000)
}

var total int

func sum(i int) {
	time.Sleep(time.Millisecond * 40)
	total = total + i
}
