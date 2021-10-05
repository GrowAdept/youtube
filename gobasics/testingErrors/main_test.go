package main

import (
	"errors"
	"io/fs"
	"testing"
)

func TestCheckNameLen(t *testing.T) {
	okNames := []string{"Samantha", "1234567890123456"}
	for _, n := range okNames {
		actual := checkNameLen(n)
		if actual != nil {
			t.Errorf("actual: %v expected: nil\n", actual)
		}
	}
	notOKnames := []string{"Blake", "Too-Long-Username"}
	for _, n := range notOKnames {
		actual := checkNameLen(n)
		if actual != errName {
			t.Errorf("got: %v expected: username length not between 5 and 17 characters\n", actual)
		}
	}
}

func TestOpenFile(t *testing.T) {
	actual := openFile("non-existing-file")
	if !errors.Is(actual, fs.ErrNotExist) {
		t.Errorf("actual: %v expected: fs.ErrNotExist\n", actual)
	}
	actual = openFile("myFile.txt")
	if !errors.Is(actual, nil) {
		t.Errorf("actual: %v expected: nil\n", actual)
	}
}
