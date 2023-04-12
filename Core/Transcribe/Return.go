package Transcribe

import (
	"os"
)

func ExitCheck() {
	if len(os.Args) != 3 {
		os.Exit(0)
	}
}
