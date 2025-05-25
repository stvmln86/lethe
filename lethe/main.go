package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/stvmln86/lethe/lethe/items/book"
	"github.com/stvmln86/lethe/lethe/tools/clui"
)

func try(err error) {
	if err != nil {
		fmt.Printf("Error: %s.\n", err.Error())
		os.Exit(2)
	}
}

func main() {
	// prep objects
	dire, err := clui.Env("LETHE_DIR")
	try(err)

	extn, err := clui.Env("LETHE_EXT")
	try(err)

	book := book.New(dire, extn, 0666)

	// prep flags
	fset := flag.NewFlagSet("lethe", flag.ExitOnError)
	read := fset.String("r", "", "print a note")
	try(fset.Parse(os.Args[1:]))

	// run main
	switch {
	case *read != "":
		note, err := book.Get(*read)
		try(err)

		body, err := note.Read()
		try(err)

		fmt.Print(body)

	default:
		for _, note := range book.Match(fset.Arg(0)) {
			fmt.Printf("%s\n", note.Name())
		}
	}
}
