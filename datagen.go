package main

import "fmt"
import "time"
import "strconv"
import "math"
import "strings"
import "os"
import "flag"

func high_resolution_sleep(duration float64) {
	// duration is expressed in seconds so needs conversion
	nanoduration := int64(duration * 1e9)
	var time_sleep0 int64 = 4900
	end := time.Now().UnixNano() + nanoduration - time_sleep0
	if duration > 0.02 {
		time.Sleep(time.Duration(nanoduration))
	}
	for time.Now().UnixNano() < end {
		time.Sleep(0)
	}
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

	lenPtr   := flag.Int("l", 1024, "record length")
	numPtr   := flag.Int("n", 100,  "number of records")
	jlenPtr  := flag.Int("j", 0,    "jitter length (default 0)")
	ratePtr  := flag.Float64("r", 100.0, "message rate")
	jratePtr := flag.Float64("f", 0.0, "message rate jitter (default 0.00)")
	flag.Parse()

	fmt.Fprintf(os.Stderr,"%s will generate %d records of %d [+/- %d] bytes at %.2f [+/- %.2f] rps\n",
			myName, *numPtr, *lenPtr, *jlenPtr, *ratePtr, *jratePtr)

	formatlen  := *lenPtr - 1
	formatlenj := formatlen 
	formatstr  := "%0" + strconv.Itoa(formatlenj) + "d"
	waittime   := 1.00 / *ratePtr
	bytecount  := 0
	progress_freq := 5
	if *numPtr / 50 > progress_freq {progress_freq = *numPtr / 50}
	time_start := time.Now().UnixNano()
	for i := 1; i <= *numPtr; i++ {
		time_now := time.Now().UnixNano()
		bytecount += formatlenj + 1
		fmt.Printf(formatstr + "\n",time_now)
		if i % progress_freq == 0 {
			status := fmt.Sprintf("%d @%.2f rps. Bytes: %d <%.2f bytes> ",i,float64(i)*1e9/float64(time_now-time_start),bytecount,float64(bytecount)/float64(i))
			progress(uint32(i), uint32(*numPtr), status)
		}
		high_resolution_sleep(waittime)
	}
	time_now := time.Now().UnixNano()
	status := fmt.Sprintf("%d @%.2f rps. Bytes: %d <%.2f bytes>\n",*numPtr,float64(*numPtr)*1e9/float64(time_now-time_start),bytecount,float64(bytecount)/float64(*numPtr))
	progress(uint32(*numPtr), uint32(*numPtr), status)
}
