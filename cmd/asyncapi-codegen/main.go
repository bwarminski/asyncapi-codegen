package main

import (
	"fmt"
	"os"
)

func run() int {
	//flags := ProcessFlags()

	fmt.Fprintf(os.Stderr, "I work!")
	return 255

	//cg, err := codegen.FromFile(flags.InputPath)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	//	return 255
	//}
	//
	//opt, err := flags.ToCodegenOptions()
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	//	return 255
	//}
	//
	//if err := cg.Generate(opt); err != nil {
	//	fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	//	return 255
	//}
	//
	//return 0
}

func main() {
	os.Exit(run())
}
