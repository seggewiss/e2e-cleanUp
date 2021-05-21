package output

import (
	"fmt"
	"time"
)

const (
	InfoColor  = "\033[1;34m%s\033[0m"
	ErrorColor = "\033[1;31m%s\033[0m"
	DebugColor = "\033[0;36m%s\033[0m"
)

// always prints the command output and fatally logs errors
func HandleCommandOutput(out []byte, e error) {
	fmt.Printf(DebugColor, string(out))

	if e != nil {
		LogError(e.Error())
	}
}

// logs "current-time: Handling url"
func LogInfo(url string) {
	fmt.Printf(InfoColor, time.Now().UTC().Format(time.UnixDate)+": Handling "+url+"\n")
}

func LogError(errorMsg string) {
	fmt.Printf(ErrorColor, errorMsg+"\n")
}
