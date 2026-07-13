package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Intent Parser mapping text to key-value tokens (e.g., login=1;db=0)
func parseToTokens(prompt string) string {
	hasLogin, hasDB, hasAPI := "0", "0", "0"
	words := strings.Fields(strings.ToLower(prompt))
	for _, word := range words {
		if word == "login" || word == "auth" || word == "signin" { hasLogin = "1" }
		if word == "database" || word == "storage" || word == "db" { hasDB = "1" }
		if word == "api" || word == "internet" || word == "network" { hasAPI = "1" }
	}
	return fmt.Sprintf("login=%s;db=%s;api=%s", hasLogin, hasDB, hasAPI)
}

func main() {
	const (
		Reset   = "\033[0m"
		Green   = "\033[32m"
		Cyan    = "\033[36m"
		Yellow  = "\033[33m"
		ClearLN = "\033[2K"
		Hide    = "\033[?25l"
		Show    = "\033[?25h"
	)

	fmt.Printf("%s====================================================%s\n", Cyan, Reset)
	fmt.Printf("%s 🤖 Welcome to GoBuilder Smart Pipeline Engine    %s\n", Green, Reset)
	fmt.Printf("%s====================================================%s\n", Cyan, Reset)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Describe the mobile app features to compile into APK: ")
	if !scanner.Scan() { return }
	userPrompt := strings.TrimSpace(scanner.Text())

	if userPrompt == "" { return }
	tokenPayload := parseToTokens(userPrompt)

	// Single-line dynamic background process spinner animation
	fmt.Print(Hide)
	done := make(chan bool)
	go func() {
		frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		stages := []string{
			"GoBuilder: Mapping user intent tokens",
			"GoBuilder: Constructing Android XML layout assets",
			"GoBuilder: Packaging UI systems via AAPT2 Core",
			"GoBuilder: Synthesizing clean Java Activity classes",
			"GoBuilder: Compiling bytecodes into Dalvik Dex formats",
			"GoBuilder: Merging all deployment packages into APK",
		}
		i, stageIdx := 0, 0
		ticker := time.NewTicker(80 * time.Millisecond)
		stageTicker := time.NewTicker(1200 * time.Millisecond)
		defer ticker.Stop()
		defer stageTicker.Stop()

		for {
			select {
			case <-done: return
			case <-ticker.C:
				fmt.Print("\r", ClearLN)
				fmt.Printf("%s%s... %s[%s]%s", Yellow, stages[stageIdx], Cyan, frames[i%len(frames)], Reset)
				i++
			case <-stageTicker.C:
				if stageIdx < len(stages)-1 { stageIdx++ }
			}
		}
	}()

	// Send request and token payload to builder.go
	cmd := exec.Command("go", "run", "builder.go", tokenPayload)
	outputBytes, _ := cmd.CombinedOutput()
	
	done <- true
	fmt.Print("\r", ClearLN)
	fmt.Print(Show)

	// Display final logs from builder
	fmt.Print(string(outputBytes))
}
