package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/schollz/progressbar/v3"
)

const (
	workDuration     = 1 * time.Minute
	shortBreak       = 5 * time.Minute
	longBreak        = 15 * time.Minute
	pomodoroSessions = 4
)

func main() {
	sessionCount := 0

	fmt.Println("Pomodoro Timer Started!")
	for {
		// Start work session
		fmt.Printf("Session %d: Work for %v minutes\n", sessionCount+1, workDuration.Minutes())
		runTimer(workDuration)

		sessionCount++
		if sessionCount%pomodoroSessions == 0 {
			// After 4 sessions, take a long break
			fmt.Printf("Time for a long break of %v minutes\n", longBreak.Minutes())
			runTimer(longBreak)
		} else {
			// Short break
			fmt.Printf("Time for a short break of %v minutes\n", shortBreak.Minutes())
			runTimer(shortBreak)
		}

		if sessionCount == pomodoroSessions {
			fmt.Println("Congratulations! You completed a full Pomodoro cycle!")
			break
		}
	}

	fmt.Println("Pomodoro Timer Finished!")
}

// runTimer runs a timer for the specified duration and displays a progress bar.
func runTimer(duration time.Duration) {
	bar := progressbar.NewOptions(int(duration.Seconds()),
		progressbar.OptionSetDescription("Time remaining"),
		progressbar.OptionSetWidth(15),
		progressbar.OptionClearOnFinish(),
	)

	// Handle Interrupt Signal (Ctrl+C)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nTimer interrupted! Exiting...")
		os.Exit(1)
	}()

	for i := 0; i < int(duration.Seconds()); i++ {
		time.Sleep(1 * time.Second)
		bar.Add(1)
	}
	bar.Finish()
}
