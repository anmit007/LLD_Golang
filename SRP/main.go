package main

import (
	"fmt"

	"github.com/anmit007/LLD/SRP/SRP/journal"
	"github.com/anmit007/LLD/SRP/SRP/persistence"
)

func main() {
	j := journal.Journal{}
	p := persistence.Persistence{LineSeparator: "\n"}
	j.AddEntry("I am sad")
	j.AddEntry("I am dead")
	j.RemoveEntry(1)
	fmt.Println(j.String())
	p.SaveToFile(&j, "journal11.txt")

}
