package main

import (
	"testing"
)

func init() {
	panic("this never works")
}

func TestNothing(t *testing.T) {
	t.Fatal("this never gets run and goConvey shows a PASS")
}
