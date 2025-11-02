package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/tnyeanderson/debate-timer/internal/debatetimer"
)

func listenForInput(r *bufio.Reader, d *debatetimer.DebateTimer) error {
	fmt.Println("Press a number to begin timing that speaker")
	for {
		char, _, err := r.ReadRune()
		if err != nil {
			return err
		}
		if string(char) == "q" {
			fmt.Println(d.Report())
			return nil
		}
		speakerNumber, err := strconv.Atoi(string(char))
		if err != nil {
			unsupported := debatetimer.NewErrorUnsupportedSpeaker(string(char))
			fmt.Println(unsupported.Error())
			continue
		}
		if err := d.StartTimer(speakerNumber); err != nil {
			if unsupported, ok := err.(debatetimer.ErrorUnsupportedSpeaker); ok {
				fmt.Println(unsupported.Error())
				continue
			}
			return err
		}
		fmt.Printf("speaker %v is now speaking\n", speakerNumber)
	}
}

func restoreTerminal() error {
	if err := exec.Command("stty", "-F", "/dev/tty", "sane").Run(); err != nil {
		return err
	}
	return nil
}

func setupTerminal() error {
	if err := exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run(); err != nil {
		return err
	}
	if err := exec.Command("stty", "-F", "/dev/tty", "-echo").Run(); err != nil {
		return err
	}
	return nil
}

func main() {
	defer func() {
		if err := restoreTerminal(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		<-sigChan
		fmt.Println("\nProgram interrupted. Exiting cleanly...")
		if err := restoreTerminal(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	if err := setupTerminal(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	d := &debatetimer.DebateTimer{}
	r := bufio.NewReader(os.Stdin)
	if err := listenForInput(r, d); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
