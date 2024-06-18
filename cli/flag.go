package cli

import (
	"errors"
	"flag"
	"strconv"
)

// FlagSet is a wrapper around flag.FlagSet
// It is used to add more functionality to the flag.FlagSet such as int32 support
type FlagSet struct {
	flag.FlagSet
}

// errParse is returned by Set if a flag's value fails to parse, such as with an invalid integer for Int.
// It then gets wrapped through failf to provide more information.
var errParse = errors.New("parse error")

// errRange is returned by Set if a flag's value is out of range.
// It then gets wrapped through failf to provide more information.
var errRange = errors.New("value out of range")

func numError(err error) error {
	ne, ok := err.(*strconv.NumError)
	if !ok {
		return err
	}
	if ne.Err == strconv.ErrSyntax {
		return errParse
	}
	if ne.Err == strconv.ErrRange {
		return errRange
	}
	return err
}

// -- int32 Value
type int32Value int32

func newInt32Value(val int32, p *int32) *int32Value {
	*p = val
	return (*int32Value)(p)
}

func (i *int32Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		err = numError(err)
	}
	*i = int32Value(v)
	return err
}

func (i *int32Value) Get() any { return int32(*i) }

func (i *int32Value) String() string { return strconv.FormatInt(int64(*i), 10) }

// Int32 defines a int32 flag with specified name, default value, and usage string.
// The return value is the address of an int32 variable that stores the value of the flag.
func (f *FlagSet) Int32Var(p *int32, name string, value int32, usage string) {
	f.Var(newInt32Value(value, p), name, usage)
}

// NOTE: CommandLine is a global variable, so we can't use it in the package
// // IntVar defines an int flag with specified name, default value, and usage string.
// // The argument p points to an int32 variable in which to store the value of the flag.
// func Int32Var(p *int32, name string, value int32, usage string) {
// 	flag.CommandLine.Var(newInt32Value(value, p), name, usage)
// }

// Int defines an int flag with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
func (f *FlagSet) Int32(name string, value int32, usage string) *int32 {
	p := new(int32)
	f.Int32Var(p, name, value, usage)
	return p
}

// NOTE: CommandLine is a global variable, so we can't use it in the package
// // Int defines an int flag with specified name, default value, and usage string.
// // The return value is the address of an int variable that stores the value of the flag.
// func Int32(name string, value int32, usage string) *int32 {
// 	return flag.CommandLine.Int(name, value, usage)
// }
