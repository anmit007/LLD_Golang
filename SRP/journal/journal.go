package journal

import (
	"fmt"
	"strings"
)

type Journal struct {
	Entries []string
}

var entryCount = 0

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d:%s", entryCount, text)
	j.Entries = append(j.Entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) int {
	if entryCount == 0 {
		fmt.Println("NO ENTRIES TO REMOVE")
		return 0
	}
	j.Entries = append(j.Entries[:index], j.Entries[index+1:]...)
	entryCount--
	return entryCount
}

func (j *Journal) String() string {
	return strings.Join(j.Entries, "\n")
}
