package main

import (
	"reflect"
	"testing"
)

func TestFun(t *testing.T) {
	r := []rune("0001?")
	ans := make([]string, 0)
	out := process(r, 0, ans)
	want := []string{"00010", "00011"}
	if reflect.DeepEqual(out, want) != true {
		t.Errorf("got %v, want %v", out, want)
	}

}
