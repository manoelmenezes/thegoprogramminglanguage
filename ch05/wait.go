package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
)

func main() {
    log.SetPrefix("wait: ")
    // suppress the date and time log package prints by default
    log.SetFlags(0)

    if err := WaitForServer(os.Args[1]); err != nil {
        log.Fatalf("Site is down: %v\n", err)        
    }
}

func WaitForServer(url string) error {
    const timeout = 1 * time.Minute
    deadline := time.Now().Add(timeout)
    for tries := 0; time.Now().Before(deadline); tries++ {
        _, err := http.Head(url)
        if err == nil {
            return nil // success
        }
        log.Printf("server not responding (%s); retrying...", err)
        time.Sleep(time.Second << uint(tries)) // exponential back-off
    }
    return fmt.Errorf("server %s failed to respond after %s", url, timeout)   
}

