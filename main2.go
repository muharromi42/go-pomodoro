package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
    defaultWorkMinutes  = 25
    defaultBreakMinutes = 5
    defaultLongBreakMinutes = 15
    longBreakInterval  = 4
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go-pomodoro <work_minutes> <break_minutes> <long_break_minutes> <long_break_interval>")
        os.Exit(1)
    }

    workMinutes, _ := strconv.Atoi(os.Args[1])
    breakMinutes, _ := strconv.Atoi(os.Args[2])
    longBreakMinutes, _ := strconv.Atoi(os.Args[3])
    longBreakInterval, _ := strconv.Atoi(os.Args[4])

    if workMinutes <= 0 {
        workMinutes = defaultWorkMinutes
    }
    if breakMinutes <= 0 {
        breakMinutes = defaultBreakMinutes
    }
    if longBreakMinutes <= 0 {
        longBreakMinutes = defaultLongBreakMinutes
    }
    if longBreakInterval <= 0 {
        longBreakInterval = longBreakInterval
    }

    pomodoro(workMinutes, breakMinutes, longBreakMinutes, longBreakInterval)
}

func pomodoro(workMinutes, breakMinutes, longBreakMinutes, longBreakInterval int) {
    cycles := 0

    for {
        fmt.Printf("Starting work for %d minutes...\n", workMinutes)
        timer(workMinutes)

        cycles++
        if cycles%longBreakInterval == 0 {
            fmt.Printf("Taking a long break for %d minutes...\n", longBreakMinutes)
            timer(longBreakMinutes)
        } else {
            fmt.Printf("Taking a short break for %d minutes...\n", breakMinutes)
            timer(breakMinutes)
        }
    }
}

func timer(minutes int) {
    duration := time.Duration(minutes) * time.Minute
    start := time.Now()
    end := start.Add(duration)

    for time.Now().Before(end) {
        remaining := end.Sub(time.Now())
        fmt.Printf("\rTime remaining: %v", remaining.Round(time.Second))
        time.Sleep(1 * time.Second)
    }

    fmt.Println("\nTime's up!")
}
