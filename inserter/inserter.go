package inserter

import (
	"Goroutine-lesson/log"
	"fmt"
)

type Inserter struct{}

// Ємуляция отправки в базу данніх
func (li Inserter) Insert(logs []log.Log) {
	fmt.Println("<INSERTING LOGS> ")

	for _, l := range logs {
		fmt.Printf("ID = %v; V = %v\n", l.Id, l.Value)
	}

	fmt.Printf("<INSERTED %d LOGS > \n", len(logs))
}
