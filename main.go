package main

import (
	"github.com/colinricardo/opencode/cmd"
	"github.com/colinricardo/opencode/internal/logging"
)

func main() {
	defer logging.RecoverPanic("main", func() {
		logging.ErrorPersist("Application terminated due to unhandled panic")
	})

	cmd.Execute()
}
