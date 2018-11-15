package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println()

	for j := 0; j < 3; j++ {
		// Determine time.Now() monotonic resolution.
		if j > 0 {
			//fmt.Println("time.Sleep(3ms)")
			//time.Sleep(3 * time.Millisecond)
		}
		for i := 0; i < 5; i++ {
			start := time.Now()
			time.Sleep(0 * time.Microsecond)
			stop  := time.Now()
			fmt.Printf("    time.Now() monotonic resolution:  %5.0fus\n", stop.Sub(start).Seconds()*1e6)
		}
	}

}
