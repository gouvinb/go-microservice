package utils

import "fmt"

// StringSlice Is a string slice for flag array string
type StringSlice []string

func (i *StringSlice) String() string {
	return fmt.Sprintf("%s", *i)
}

// Set append a value into StringSlice
func (i *StringSlice) Set(value string) error {
	*i = append(*i, value)
	return nil
}
