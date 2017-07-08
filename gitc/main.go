package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/PaluMacil/misc/gitc/ginit"
)

func main() {
	if len(os.Args) == 1 {
		info()
		return
	}

	// Subcommand for total destruction of junk config files
	nukeCommand := flag.NewFlagSet("nuke", flag.ExitOnError)

	// Subcommand for bulk slicing and dicing of config file contents
	sliceCommand := flag.NewFlagSet("slice", flag.ExitOnError)

	// Subcommand for bulk git init
	ginitCommand := flag.NewFlagSet("ginit", flag.ExitOnError)
	ginitCommand.BoolVar(&ginit.AddIgnore, "ignore", false, "add a .gitignore to each Git repo")

	switch os.Args[1] {
	case "nuke":
		nukeCommand.Parse(os.Args[2:])
	case "slice":
		sliceCommand.Parse(os.Args[2:])
	case "ginit":
		ginitCommand.Parse(os.Args[2:])
		ginit.Do()
	default:
		fmt.Printf("%q is not valid command.\n\n", os.Args[1])
		info()
		os.Exit(2)
	}
}

func info() {
	fmt.Println("usage: gitc <command> [<args>]")
	fmt.Println("The most commonly used git commands are: ")
	fmt.Println("  nuke    destroy junk files")
	fmt.Println("  slice   slice junk settings from larger files")
	fmt.Println("  ginit")
}
