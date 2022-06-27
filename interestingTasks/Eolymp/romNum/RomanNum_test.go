package main

import (
	"reflect"
	"testing"
)

func TestRomNumberInLatNumb(t *testing.T) {
	list := []string{"III", "IV"}

	got := RomNumberInLatNumb(list)
	want := []int{3, 4}
	check := reflect.DeepEqual(got, want)
	if check != true {
		t.Error("didn't match expected value")
	}
}
