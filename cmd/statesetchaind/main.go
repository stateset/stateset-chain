package main

import (
	"os"

	"github.com/stateset/statesetchain/cmd/statesetchaind/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
