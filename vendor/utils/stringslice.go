// Copyright 2016 gouvinb. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

// Package utils contains some methods/struct for all projet.
package utils

import "fmt"

// StringSlice Is a string slice for flag array string
type StringSlice []string

// String return slice to string type
func (i *StringSlice) String() string {
	return fmt.Sprintf("%s", *i)
}

// Set append a value into StringSlice
func (i *StringSlice) Set(value string) error {
	*i = append(*i, value)
	return nil
}
