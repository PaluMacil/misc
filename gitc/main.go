package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/PaluMacil/misc/gitc/ginit"
	"github.com/PaluMacil/misc/gitc/touch"
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

	// Subcommand for touching a repo to add commits to a date or range of dates
	touchCommand := flag.NewFlagSet("ginit", flag.ExitOnError)
	touchCommand.StringVar(&touch.RepoURL, "repo", "", "the Git repo")
	touchCommand.StringVar(&touch.SinceString, "since", "", "date to start touching in format YYMMDD; default is today")
	touchCommand.StringVar(&touch.User, "user", "", "Git username")
	touchCommand.StringVar(&touch.Password, "password", "", "password")
	touchCommand.StringVar(&touch.Name, "name", "", "password")
	touchCommand.StringVar(&touch.Email, "email", "", "password")

	switch os.Args[1] {
	case "nuke":
		nukeCommand.Parse(os.Args[2:])
	case "slice":
		sliceCommand.Parse(os.Args[2:])
	case "ginit":
		ginitCommand.Parse(os.Args[2:])
		ginit.Do()
	case "touch":
		touchCommand.Parse(os.Args[2:])
		touch.Do()
	case "deltmp":
		touchCommand.Parse(os.Args[2:])
		touch.DeleteTemp()
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
	fmt.Println("  touch")
	fmt.Println("  deltmp")
}
