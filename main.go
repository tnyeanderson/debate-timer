package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/mattn/go-tty"
	"github.com/tnyeanderson/debate-timer/internal/debatetimer"
	"gopkg.in/yaml.v3"
)

const doc = `Press a number to begin timing that speaker.
Press p to pause all timers.
Press q to quit and print the report.`

func printSpeakerNames(w io.Writer) {
	for i := 1; i < 10; i++ {
		if name := debatetimer.GetSpeakerName(i); name != "" {
			fmt.Fprintf(w, "Speaker %v is %v\n", i, name)
		}
	}
}

func listenForInput(d *debatetimer.DebateTimer) error {
	fmt.Fprintln(os.Stderr, doc)
	printSpeakerNames(os.Stderr)
	fmt.Fprintln(os.Stderr, "---")
	tty, err := tty.Open()
	if err != nil {
		return err
	}
	defer tty.Close()
	exitCleanlyOnSignal(tty)
	for {
		char, err := tty.ReadRune()
		if err != nil {
			return err
		}
		stop, err := handleInput(char, d)
		if err != nil {
			return err
		}
		if stop {
			return nil
		}
	}
}

func handleInput(char rune, d *debatetimer.DebateTimer) (stop bool, err error) {
	if string(char) == "q" {
		if err := printReport(d); err != nil {
			return true, err
		}
		return true, nil
	}
	if string(char) == "p" {
		if err := d.Pause(); err != nil {
			return true, err
		}
		fmt.Fprintln(os.Stderr, "Timer is paused")
		return false, nil
	}
	speakerNumber, err := strconv.Atoi(string(char))
	if err != nil {
		unsupported := debatetimer.NewErrorUnsupportedSpeaker(string(char))
		fmt.Fprintln(os.Stderr, unsupported.Error())
		return false, nil
	}
	if err := d.StartTimer(speakerNumber); err != nil {
		if unsupported, ok := err.(debatetimer.ErrorUnsupportedSpeaker); ok {
			fmt.Fprintln(os.Stderr, unsupported.Error())
			return false, nil
		}
		if alreadySpeaking, ok := err.(debatetimer.ErrorAlreadySpeaking); ok {
			fmt.Fprintln(os.Stderr, alreadySpeaking.Error())
			return false, nil
		}
		return true, err
	}
	fmt.Fprintf(os.Stderr, "%v is now speaking\n", debatetimer.GetSpeakerNameDefault(speakerNumber))
	return false, nil
}

func printReport(d *debatetimer.DebateTimer) error {
	r, err := d.Report()
	if err != nil {
		return err
	}
	if true {
		fmt.Println(r.String())
		return nil
	}
	out, err := yaml.Marshal(r)
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}

func exitCleanlyOnSignal(tty *tty.TTY) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		<-sigChan
		fmt.Fprintln(os.Stderr, "\nProgram interrupted. Exiting cleanly...")
		if err := tty.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}()
}

func main() {
	d := &debatetimer.DebateTimer{}
	if err := listenForInput(d); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
