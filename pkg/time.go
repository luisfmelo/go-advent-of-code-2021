package pkg

import (
	"log"
	"time"
)

func RunWithTime(part1, part2 func() string) {
	start := time.Now()
	log.Printf("1st part result: %v\n", part1())
	log.Printf("1st part took: %s", time.Since(start))

	if part2 == nil {
		log.Printf("2nd part is not implemented yet")
		return
	}
	start = time.Now()
	log.Printf("2nd part result: %v\n", part2())
	log.Printf("2nd part took: %s", time.Since(start))
}
