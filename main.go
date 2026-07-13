package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	const (
		Reset   = "\033[0m"
		Cyan    = "\033[36m"
		ClearLN = "\033[2K"
		Hide    = "\033[?25l"
		Show    = "\033[?25h"
	)

	clearCmd := exec.Command("clear")
	clearCmd.Stdout = os.Stdout
	_ = clearCmd.Run()

	fmt.Printf("\n%s[ go-builder v3.7 ] // cyber hack tech%s\n\n", Cyan, Reset)
	fmt.Print("➔ Enter app concept: ")
	
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() { return }
	userPrompt := strings.TrimSpace(scanner.Text())

	if userPrompt == "" { return }

	fmt.Print(Hide)
	done := make(chan bool)
	go func() {
		ticker := time.NewTicker(150 * time.Millisecond)
		defer ticker.Stop()
		dots := []string{".  ", ".. ", "...", "   "}
		i := 0
		for {
			select {
			case <-done: return
			case <-ticker.C:
				fmt.Print("\r", ClearLN)
				fmt.Printf("Building application%s", dots[i%len(dots)])
				i++
			}
		}
	}()

	cmd := exec.Command("go", "run", "builder.go", userPrompt)
	outputBytes, _ := cmd.CombinedOutput()
	
	done <- true
	fmt.Print("\r", ClearLN)
	fmt.Print(Show)

	fmt.Print(string(outputBytes))
}
