// Copyright 2016 gouvinb. All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

package shared

import (
	"flag"
	"os"
	"strconv"
)

var (
	flagEnableCors = flag.Bool("cors-enable", false, "enable cors")
)

// IsCorsEnable return true if use cors
func IsCorsEnable(c Cors) bool {
	if *flagEnableCors != false {
		return *flagEnableCors
	} else if value, err := strconv.ParseBool(os.Getenv("CORS_ENABLE")); err == nil {
		return value
	} else if c.EnableCors != true {
		return c.EnableCors
	}
	return true
}
