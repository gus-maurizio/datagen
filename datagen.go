package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
*/
import "C"

import "fmt"
import "time"
import "strconv"
import "math"
import "math/rand"
import "strings"
import "os"
import "flag"
import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)


func hr_sleep_microsecond(micros int) {
        // duration is expressed in microseconds
        // minimum value accepted is actually 100 microseconds, 0.1 milliseconds
        // Precision deteriorates below 500 microseconds, so we will split that value
        // into 40 microseconds that yield about 100 microseconds effective waits
        single_call   := 500
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

func random(min, max int) int {
	return rand.Intn(max - min) + min
}

func progress(count uint32, total uint32, status string) {
	var bar_len uint32 = 50
	var filled_len int
	filled_len = int(math.Round(float64(bar_len*count) / float64(total)))
	percents := math.Round(100.0 * float64(count) / float64(total))
	bar := strings.Repeat("=", filled_len) + strings.Repeat("-", (int(bar_len)-filled_len))
	legend := fmt.Sprintf(">>> [%s] %5.2f%s ...%s\r", bar, percents, "%", status)
	os.Stderr.Write([]byte(legend))
}

func main() {
	myName   := os.Args[0]
	rand.Seed(time.Now().UTC().UnixNano())

	lenPtr   := flag.Int("l", 1024, "record length")
	numPtr   := flag.Int("n", 100,  "number of records")
	jlenPtr  := flag.Int("j", 0,    "jitter length (default 0)")
	burstPtr := flag.Int("b", 1,    "burst records sent together")
	ratePtr  := flag.Float64("r", 100.0, "message rate")
	jratePtr := flag.Float64("f", 0.0, "message rate jitter (default 0.00)")
	flag.Parse()
	
	waitperrec    := int(1000000 / *ratePtr)
	waitperrecmin := int(1000000 / (*ratePtr + *jratePtr))
	waitperrecmax := int(1000000 / (*ratePtr - *jratePtr))
	waitmicro     := waitperrec * *burstPtr
	waitmilli     := float64(waitmicro) / 1000.0

	p := message.NewPrinter(language.English)
	p.Fprintf(os.Stderr,"%s will generate %d records of %d [+/- %d] bytes at %.2f [+/- %.2f] rps (sending %d record together and waiting %d usec [%.2f msec] between bursts)\n",
		myName, *numPtr, *lenPtr, *jlenPtr, *ratePtr, *jratePtr, *burstPtr, waitmicro, waitmilli)
	if *jratePtr > 0.0 {
		p.Fprintf(os.Stderr,"%s jitter microseconds [%d - %d]\n", myName, waitperrecmin, waitperrecmax)
	}
	formatlen  := *lenPtr - 1
	formatlenj := formatlen
	formatstr  := "%0" + strconv.Itoa(formatlenj) + "f"
	bytecount  := 0
	progress_freq := 5
	if *numPtr / 50 > progress_freq {progress_freq = *numPtr / 50}
	time_start := time.Now().UnixNano()
	i := 1
	for i <= *numPtr {
		l := 1
		for l <= *burstPtr && i <= *numPtr {
			if *jlenPtr != 0 {
				formatlenj = random(formatlen - *jlenPtr, formatlen + *jlenPtr) 
				formatstr  = "%0" + strconv.Itoa(formatlenj) + "f"
			}
			time_now := time.Now().UnixNano()
			fmt.Printf(formatstr + "\n",float64(time_now)/1e9)
			bytecount += formatlenj + 1
                        if i % progress_freq == 0 {
                                status := p.Sprintf("%d @%.2f rps. Bytes: %d <%.2f bytes> ",i,float64(i)*1e9/float64(time_now-time_start),bytecount,float64(bytecount)/float64(i))
                                progress(uint32(i), uint32(*numPtr), status)
                        }
			i++
			l++
		}
		if *jratePtr > 0.00 {
			hr_sleep_microsecond(random( int(float64(1e6 * *burstPtr) / (*ratePtr + *jratePtr)) , int(float64(1e6 * *burstPtr) / (*ratePtr - *jratePtr)) ))
			
		} else {
			hr_sleep_microsecond(waitmicro)
		}
	}
	time_now := time.Now().UnixNano()
	status := p.Sprintf("%d @%.2f rps. Bytes: %d <%.2f bytes>\n",*numPtr,float64(*numPtr)*1e9/float64(time_now-time_start),bytecount,float64(bytecount)/float64(*numPtr))
	progress(uint32(*numPtr), uint32(*numPtr), status)
}
