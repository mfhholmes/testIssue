package main

import (
	"sync"
	"testing"
)

var failOnce sync.Once

func init() {
	failOnce.Do(openFail)
}

func openFail() {
	// fail at this...
	panic("this never works")
}

func TestNothing(t *testing.T) {
	t.Fatal("this never gets run and goConvey shows a PASS")
}
