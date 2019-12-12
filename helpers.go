package usage

import (
	"fmt"
	"io"
)

func CheckErrorWithPanic(logger io.Writer, err error) {
	checkError(logger, err, true)
}

func CheckErrorWithOnlyLogging(logger io.Writer, err error) {
	checkError(logger, err, false)
}

func checkError(logger io.Writer, err error, _panic bool) {
	if err != nil {
		if _panic {
			panic("error: " + err.Error())
		}

		_, err := fmt.Fprintf(logger, "error: %s\n", err.Error())
		checkError(logger, err, false)
	}
}
