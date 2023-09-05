package logs

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func getFunctionInfo() (fileName string, functionName string, lineNumber int) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Error retrieving function info")
		return
	}

	functionName = runtime.FuncForPC(pc).Name()
	fileName = filepath.Base(file) // Get the short file name from the full path
	lineNumber = line

	return
}
