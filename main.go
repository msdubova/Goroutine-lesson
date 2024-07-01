package main

import (
	"Goroutine-lesson/inserter"
	"Goroutine-lesson/log"
	"fmt"

	"math/rand"
)

// создадим обертку для записи в базу данніх - LogInserter - тип записує логи. Цей тип повинен записувати логи не одразу а раз в три сек, не блокуючи інщі частини програми.  - тобто зробити  LogInserter конкурентним.
func main() {
	var inserter inserter.Inserter

	for i := 0; i < 10; i++ {

		logs := make([]log.Log, 0, rand.Intn(10))
		for j := 0; j < cap(logs); j++ {
			l := log.NewRandom(fmt.Sprintf("%d-%d", i, j))
			logs = append(logs, l)
		}

		inserter.Insert(logs)
	}

}
