package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
*/
import "C"

import "fmt"
import "time"

func hr_sleep_microsecond(micros int) {
	// duration is expressed in microseconds
	// minimum value accepted is actually 100 microseconds, 0.1 milliseconds
	// Precision deteriorates below 500 microseconds, so we will split that value
	// into 40 microseconds that yield about 100 microseconds effective waits
	single_call := 500
	minimum_value := 100
	low_precision := 40
	if micros >= single_call {
		C.usleep(_Ctype_uint(micros))
		return
	}
	// if we are here is because micros value is less than single call
	// let's check if it is below minimum
	if micros <= minimum_value {
		C.usleep(_Ctype_uint(low_precision))
		return
	}
	// if we are here is because micros value is low but not that low
	loops := micros / minimum_value
	for i := 1; i <= loops; i++ {
		C.usleep(_Ctype_uint(low_precision))
	}
}

func test(micros int) {
	fmt.Printf("Comparing time.Sleep with hr_sleep_microsecond for %d usec\n", micros)
        for i := 0; i < 5; i++ {
                start := time.Now()
                time.Sleep(time.Duration(micros * 1000))
                stop := time.Now()
                startc := time.Now()
                hr_sleep_microsecond(micros)
                stopc := time.Now()

                fmt.Printf("micros=%dus time.Sleep() hr_sleep_microsecond monotonic resolution:  %5.0fus %5.0fus\n", micros, stop.Sub(start).Seconds()*1e6, stopc.Sub(startc).Seconds()*1e6)
        }

}

func main() {
	fmt.Println()

	test(2000)
	test(1000)

	test(500)

	test(200)
	test(100)
	test(50)

}
