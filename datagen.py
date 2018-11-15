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
	end := time.Now().UnixNano() + nanoduration
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

	lenPtr   := flag.Int("l", 1024, "record length")
	numPtr   := flag.Int("n", 100,  "number of records")
	jlenPtr  := flag.Int("j", 0,    "jitter length (default 0)")
	ratePtr  := flag.Float64("r", 100.0, "message rate")
	jratePtr := flag.Float64("f", 0.0, "message rate jitter (default 0.00)")
	flag.Parse()
	fmt.Printf("args len %d num %d jlen %d rate %f jrate %f \n", *lenPtr, *numPtr, *jlenPtr, *ratePtr, *jratePtr) 
	fmt.Println("Now,     " + strconv.FormatInt(time.Now().Unix(), 10))
	fmt.Println("Hello,   " + strconv.FormatInt(time.Now().UnixNano(), 10))
	high_resolution_sleep(.00005)
	fmt.Println("goodbye, " + strconv.FormatInt(time.Now().UnixNano(), 10))
	for i := 0; i < 100; i++ {
		progress(uint32(i), 100, "hello")
		high_resolution_sleep(0.05)
	}
	progress(100, 100, "Finished \n")
}
