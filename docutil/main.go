package main

import (
	"fmt"
	"log"

	"github.com/0xrahulm/dxp-cli/cmd"
)

var (
	// Directory in which the documentation will be generated.
	docDir = "doc"

	// Tag used to delimitate the section of the README which is generated.
	delimiter = "generated"
)

func main() {
	dxpcli := cmd.NewDxpcliCommand()

	// Generate documentation for the `dxpcli` command.
	if err := genMarkdownDoc(dxpcli, docDir); err != nil {
		fmt.Println("Unable to generate documentation.")
		log.Fatal(err)
	}
	fmt.Println("Documentation generated!")

	// Update the summary of commands in the `README.md` (located inside <tag></tag>)
	if err := updateReadmeCommands(dxpcli, delimiter, docDir); err != nil {
		fmt.Println("Unable to update `README.md`.")
		log.Fatal(err)
	}
	fmt.Println("`README.md` updated!")
}
