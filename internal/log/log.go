package log

import (
	"fmt"
	"os"
)

var Verbose bool

func LogVerbose(format string, a ...any) {
	if Verbose {
		fmt.Fprintf(os.Stderr, format+"\n", a...)
	}
}
