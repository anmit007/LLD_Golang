package persistence

import (
	"fmt"
	"os"
	"strings"

	"github.com/anmit007/LLD/SRP/SRP/journal"
)

type Persistence struct {
	LineSeparator string
}

func (p *Persistence) SaveToFile(j *journal.Journal, filename string) error {
	fmt.Println("SAVING TO FILE")
	err := os.WriteFile(filename,
		[]byte(strings.Join(j.Entries, p.LineSeparator)), 0644)
	if err != nil {
		fmt.Println("SAVING TO FILE Failed")
	}
	println("FILED SAVED")
	return err
}
