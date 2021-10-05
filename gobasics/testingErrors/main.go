package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

type errorOther struct{}

func (e errorOther) Error() string {
	return "this is a different error"
}

var errOther errorOther

// func New(text string) error
var errName = errors.New("username length not between 5 and 17 characters")

func main() {
	err := openFile("non-existing-file")
	fmt.Printf("err value:%v err type: %T\n", err, err)
	// func Is(err, target error) bool
	// Is reports whether any error in err's chain matches target.
	if errors.Is(err, fs.ErrNotExist) {
		fmt.Println("errors.Is(err, fs.ErrNotExist) equates to true (err is fs.ErrNotExist)")
	} else {
		fmt.Println("errors.Is(err, fs.ErrNotExist) equates to false (err is something other than fs.ErrNotExist")
	}
	if errors.Is(err, fs.ErrPermission) {
		fmt.Println("errors.Is(err, fs.ErrPermission) equates to true (err is fs.ErrPermission)")
	} else {
		fmt.Println("errors.Is(err, fs.ErrPermission) equates to false (err is something other than fs.ErrPermission)")
	}
	fmt.Println("********************************************************************")
	// func New(text string) error
	// Each call to New returns a distinct error value even if the text is identical.
	// These do not point to the same value in memory or are not nil (conditions to be equal)
	err1 := errors.New("username length not between 5 and 17 characters")
	err2 := errors.New("username length not between 5 and 17 characters")
	fmt.Printf("err1 value:%v err1 type: %T\n", err1, err1)
	fmt.Printf("err2 value:%v err2 type: %T\n", err2, err2)
	if err1 == err2 {
		fmt.Println("err1 == err2 equates to true")
	} else {
		fmt.Println("err1 == err2 equates to false")
	}
	// this does not work either
	if errors.Is(err1, err2) {
		fmt.Println("errors.Is(err1, err2) equates to true")
	} else {
		fmt.Println("errors.Is(err1, err2) equates to false")
	}
	fmt.Println("********************************************************************")
	// name := "bob"
	name := "Too-Long-Username"
	err3 := checkNameLen(name)
	if err3 == errName {
		fmt.Println("err3 == errName equates to true")
	} else {
		fmt.Println("err3 == errName equates to false")
	}
	if errors.Is(err3, errName) {
		fmt.Println("errors.Is(err3, errName) equates to true")
	} else {
		fmt.Println("errors.Is(err3, errName) equates to false")
	}
	if err3 == errOther {
		fmt.Println("err3 == errOther equates to true")
	} else {
		fmt.Println("err3 == errOther equates to false")
	}
	if errors.Is(err3, errOther) {
		fmt.Println("errors.Is(err3, errOther) equates to true")
	} else {
		fmt.Println("errors.Is(err3, errOther) equates to false")
	}
}

// checkNameLen checks that a username is between 5 and 17 characters long (non-inclusive)
// checkNameLen returns nil error or *errors.errorString
func checkNameLen(name string) (err error) {
	// if 5 <= len(name) && len(name) <= 17 { // incorrect code for test fail demonstration
	if 5 < len(name) && len(name) < 17 { // correct
		return err
	}
	// return errors.New("username length not between 5 and 17 characters") // don't use for equality check
	return errName
}

// openFile return a nil error or *fs.PathError
func openFile(s string) error {
	// func OpenFile(name string, flag int, perm FileMode) (*File, error)
	_, err := os.Open(s)
	// normally would do something with returned file here
	return err
}
