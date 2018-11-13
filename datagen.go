package main

import "fmt"
import "time"
import "strconv"

func high_resolution_sleep(duration float64) {
    // duration is expressed in seconds so needs conversion
    end := time.Now().UnixNano() + int64(duration * 1e9)
    if duration > 0.02 {
        time.Sleep(time.Duration(duration * 1e9))
    }
    for time.Now().UnixNano() < end {
        time.Sleep(0)
    }
}

func main() {
    fmt.Println("Now,     " + strconv.FormatInt(time.Now().Unix(), 10))
    fmt.Println("Hello,   " + strconv.FormatInt(time.Now().UnixNano(), 10))
    high_resolution_sleep(.000066)
    fmt.Println("goodbye, " + strconv.FormatInt(time.Now().UnixNano(), 10))
}
